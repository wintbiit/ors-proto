package proto

import "encoding/binary"

const S1ProtoTestAckSize = 10244

func (s *S1ProtoTestAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoTestAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	copy(bytes[4:], s.Test)
	return bytes
}

func (s *S1ProtoTestAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	s.Test = string(bytes[4:10244])
	return nil
}
