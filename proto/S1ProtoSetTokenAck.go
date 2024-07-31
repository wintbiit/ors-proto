package proto

import "encoding/binary"

const S1ProtoSetTokenAckSize = 4

func (s *S1ProtoSetTokenAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTokenAckSize)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoSetTokenAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:4]))
	return nil
}
