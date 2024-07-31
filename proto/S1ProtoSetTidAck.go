package proto

import "encoding/binary"

const S1ProtoSetTidAckSize = 8

func (s *S1ProtoSetTidAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTidAckSize)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(s.ResultId))
	binary.LittleEndian.PutUint32(bytes[4:8], s.Tid)
	return bytes
}

func (s *S1ProtoSetTidAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:4]))
	s.Tid = binary.LittleEndian.Uint32(bytes[4:8])
	return nil
}
