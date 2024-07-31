package proto

import "encoding/binary"

const S1ProtoLogoutAckSize = 4

func (s *S1ProtoLogoutAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLogoutAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoLogoutAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
