package proto

import "encoding/binary"

const S1ProtoUserStateAckSize = 4

func (s *S1ProtoUserStateAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoUserStateAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.Online))
	return bytes
}

func (s *S1ProtoUserStateAck) Deserialize(bytes []byte) error {
	s.Online = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
