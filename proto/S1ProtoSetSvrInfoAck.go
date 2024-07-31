package proto

import "encoding/binary"

const S1ProtoSetSvrInfoAckSize = 4

func (s *S1ProtoSetSvrInfoAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetSvrInfoAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoSetSvrInfoAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
