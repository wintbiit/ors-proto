package ors

import (
	"crypto/sha1"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/wintbiit/ors-proto/proto"
)

const (
	ServerPort        = 64998
	HeartBeatInterval = 200 * time.Millisecond
)

type ProtoHandler func(ctx *proto.S1ProtoContext)

type Client struct {
	IpAddress    string
	ClientId     uint16
	ClientTId    uint32
	ClientTeamId uint32
	Hash         int64
	Debug        bool

	logger     Logger
	control    chan struct{}
	heartbeat  chan struct{}
	socket     net.Conn
	handlers   map[uint16]ProtoHandler
	anyHandler ProtoHandler
	seq        byte
	bufPool    sync.Pool
}

func (s *Client) Connect() error {
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

func (s *Client) Close() error {
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

func (s *Client) IsConnected() bool {
	return s.socket != nil
}

func (s *Client) send(data []byte) error {
	if !s.IsConnected() {
		return proto.NotConnectedError
	}

	_, err := s.socket.Write(data)

	return err
}

func (s *Client) autoSeq() byte {
	s.seq++
	return s.seq
}

func (s *Client) Send(protoId uint16, data proto.Proto) error {
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

func (s *Client) packProtoData(protoId uint16, p proto.Proto) []byte {
	data := p.Serialize()

	header := proto.S1ProtoHeader{
		ProtoId:    protoId,
		DataLen:    uint32(len(data)),
		SequenceId: s.autoSeq(),
	}

	headerBytes := header.Serialize()

	return append(headerBytes, data...)
}

func (s *Client) Login(pass string) error {
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

func (s *Client) Logout() error {
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

func (s *Client) sendHeartBeat() error {
	heartBeat := &proto.S1ProtoHeartBeatReq{
		Nouse:      0,
		S0Clientid: byte(s.ClientId),
	}

	if err := s.Send(proto.ProtoIDS1ProtoHeartBeatReq, heartBeat); err != nil {
		return err
	}

	return nil
}

func (s *Client) recv() {
	for {
		select {
		case <-s.control:
			return
		default:
			s.readTcpPipe()
		}
	}
}

func (s *Client) readTcpPipe() {
	// read data from socket
	// headerBytes := make([]byte, proto.S1ProtoHeaderSize)
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

	protoName, ok := proto.ProtoIdMap[header.ProtoId]
	if !ok {
		protoName = "unknown"
	}

	if s.Debug {
		s.logger.Infof("recv proto: [%d] %s, data len: %v", header.ProtoId, protoName, len(headerBytes)+len(bodyBytes))
	}

	ctx, cancel := proto.NewS1ProtoContext(&header, bodyBytes)

	if s.anyHandler != nil {
		s.anyHandler(ctx)
	}

	// handle protocol
	handler, ok := s.handlers[header.ProtoId]
	if !ok {
		if s.Debug {
			s.logger.Warnf("unhandled proto: [%d] %s", header.ProtoId, protoName)
		}
		return
	}

	go func() {
		defer cancel()
		handler(ctx)
	}()
}

func (s *Client) heartBeat() {
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

func NewClient(ip string, clientId uint16, clientTId, clientTeamId uint32) *Client {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%d%d%d", clientId, clientTId, clientTeamId)))

	return &Client{
		IpAddress:    ip,
		ClientId:     clientId,
		ClientTId:    clientTId,
		ClientTeamId: clientTeamId,
		Hash:         0,
		handlers:     make(map[uint16]ProtoHandler),
		logger:       &slogLogger{},
		bufPool: sync.Pool{
			New: func() interface{} {
				return make([]byte, proto.S1ProtoHeaderSize)
			},
		},
	}
}

func (s *Client) WithLogger(logger Logger) *Client {
	s.logger = logger
	return s
}

func (s *Client) WithHandler(protoID uint16, handler ProtoHandler) *Client {
	if s.handlers == nil {
		s.handlers = make(map[uint16]ProtoHandler)
	}

	s.handlers[protoID] = handler

	return s
}

func (s *Client) WithAnyHandler(handler ProtoHandler) *Client {
	s.anyHandler = handler
	return s
}
