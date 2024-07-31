package proto

import "encoding/binary"

const S1ProtoPingAckSize = 4

func (s *S1ProtoPingAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoPingAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoPingAck) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoPingAckSize {
		return InvalidDataError
	}

	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
