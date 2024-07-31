package proto

import "encoding/binary"

const S1ProtoSetReadyAckSize = 4

func (s *S1ProtoSetReadyAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetReadyAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoSetReadyAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
