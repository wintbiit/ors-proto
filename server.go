package ors

import (
	"crypto/sha1"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/wintbiit/ors-proto/proto"
)

const ServerPort = 64998

type ProtoHandler func(ctx *proto.S1ProtoContext)

type Server struct {
	IpAddress    string
	ClientId     uint16
	ClientTId    uint32
	ClientTeamId uint32
	Hash         int64
	Debug        bool

	logger    Logger
	control   chan struct{}
	heartbeat chan struct{}
	socket    net.Conn
	handlers  map[uint16]ProtoHandler
	seq       byte
	bufPool   sync.Pool
}

func (s *Server) Connect() error {
	s.Close()

	// tcp connection
	addr := fmt.Sprintf("%s:%d", s.IpAddress, ServerPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	s.socket = conn
	s.control = make(chan struct{})

	go s.recv()

	return nil
}

func (s *Server) Close() error {
	if s.control != nil {
		close(s.control)
	}

	s.control = nil

	if s.heartbeat != nil {
		close(s.heartbeat)
	}

	if s.socket != nil {
		if err := s.Logout(); err != nil {
			s.socket.Close()
			return err
		}

		if err := s.socket.Close(); err != nil {
			return err
		}
	}

	s.socket = nil

	return nil
}

func (s *Server) IsConnected() bool {
	return s.socket != nil
}

func (s *Server) send(data []byte) error {
	if !s.IsConnected() {
		return proto.NotConnectedError
	}

	_, err := s.socket.Write(data)

	return err
}

func (s *Server) autoSeq() byte {
	s.seq++
	return s.seq
}

func (s *Server) Send(protoId uint16, data proto.Proto) error {
	if !s.IsConnected() {
		return proto.NotConnectedError
	}

	dataBytes := s.packProtoData(protoId, data)

	// send data
	if err := s.send(dataBytes); err != nil {
		return err
	}

	if s.Debug {
		s.logger.Infof("send protoID: %d, data len: %v", protoId, len(dataBytes))
	}

	return nil
}

func (s *Server) packProtoData(protoId uint16, p proto.Proto) []byte {
	data := p.Serialize()

	header := proto.S1ProtoHeader{
		ProtoId:    protoId,
		DataLen:    uint32(len(data)),
		SequenceId: s.autoSeq(),
	}

	headerBytes := header.Serialize()

	return append(headerBytes, data...)
}

func (s *Server) Login(pass string) error {
	login := &proto.S1ProtoLoginReq{
		Account:  fmt.Sprintf("s0unit_client_%d_%d", s.ClientId, s.ClientId),
		Password: pass,
		Version:  proto.S1StuVersion,
		Tid:      s.ClientTId,
		Teamid:   s.ClientTeamId,
		Hash:     s.Hash,
	}

	if err := s.Send(proto.ProtoIDS1ProtoLoginReq, login); err != nil {
		return err
	}

	if s.heartbeat != nil {
		close(s.heartbeat)
	}

	s.heartbeat = make(chan struct{})
	go s.heartBeat()

	return nil
}

func (s *Server) Logout() error {
	logout := &proto.S1ProtoLogoutReq{
		Account: fmt.Sprintf("s0unit_client_%d_%d", s.ClientId, s.ClientId),
	}

	if err := s.Send(proto.ProtoIDS1ProtoLogoutReq, logout); err != nil {
		return err
	}

	if s.heartbeat != nil {
		close(s.heartbeat)
	}

	return nil
}

func (s *Server) sendHeartBeat() error {
	heartBeat := &proto.S1ProtoHeartBeatReq{
		Nouse:      0,
		S0Clientid: byte(s.ClientId),
	}

	if err := s.Send(proto.ProtoIDS1ProtoHeartBeatReq, heartBeat); err != nil {
		return err
	}

	return nil
}

func (s *Server) recv() {
	for {
		select {
		case <-s.control:
			return
		default:
			s.readTcpPipe()
		}
	}
}

func (s *Server) readTcpPipe() {
	// read data from socket
	//headerBytes := make([]byte, proto.S1ProtoHeaderSize)
	headerBytes := s.bufPool.Get().([]byte)
	defer s.bufPool.Put(headerBytes)

	_, err := s.socket.Read(headerBytes)
	if err != nil {
		s.logger.Errorf("read header error: %v", err)
		s.Close()
		return
	}

	// deserialize header
	var header proto.S1ProtoHeader
	if err := header.Deserialize(headerBytes); err != nil {
		s.logger.Errorf("deserialize header error: %v", err)
		return
	}

	// read body
	bodyBytes := make([]byte, header.DataLen)
	_, err = s.socket.Read(bodyBytes)

	if err != nil {
		s.logger.Errorf("read body error: %v", err)
		return
	}

	// handle protocol
	handler, ok := s.handlers[header.ProtoId]
	if !ok {
		if s.Debug {
			protoName, ok := proto.ProtoIdMap[int(header.ProtoId)]
			if !ok {
				protoName = "unknown"
			}
			s.logger.Warnf("ignore proto: [%d] %s", header.ProtoId, protoName)
		}
		return
	}

	// create context
	ctx := proto.NewS1ProtoContext(&header, bodyBytes)

	if s.Debug {
		s.logger.Infof("recv protoID: %d, data len: %v", header.ProtoId, len(headerBytes)+len(bodyBytes))
	}

	go handler(ctx)
}

func (s *Server) heartBeat() {
	for {
		time.Sleep(HeartBeatInterval)
		select {
		case <-s.control:
			return
		case <-s.heartbeat:
			return
		default:
			if err := s.sendHeartBeat(); err != nil {
				s.logger.Errorf("send heart beat error: %v", err)
				break
			}
		}
	}
}

func NewServer(ip string, clientId uint16, clientTId, clientTeamId uint32) *Server {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%d%d%d", clientId, clientTId, clientTeamId)))

	return &Server{
		IpAddress:    ip,
		ClientId:     clientId,
		ClientTId:    clientTId,
		ClientTeamId: clientTeamId,
		Hash:         0,
		handlers:     make(map[uint16]ProtoHandler),
		logger:       &slogLogger{},
	}
}

func (s *Server) WithLogger(logger Logger) *Server {
	s.logger = logger
	return s
}

func (s *Server) WithHandler(protoID uint16, handler ProtoHandler) *Server {
	if s.handlers == nil {
		s.handlers = make(map[uint16]ProtoHandler)
	}

	s.handlers[protoID] = handler

	return s
}

var HeartBeatInterval = 200 * time.Millisecond
