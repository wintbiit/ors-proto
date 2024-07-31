package proto

import "encoding/binary"

const S1ProtoGmcommandAckSize = 4

func (s *S1ProtoGmcommandAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoGmcommandAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoGmcommandAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
