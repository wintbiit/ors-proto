package proto

import "encoding/binary"

const S1ProtoHeaderSize = 12

func (s *S1ProtoHeader) Serialize() []byte {
	bytes := make([]byte, S1ProtoHeaderSize)
	binary.LittleEndian.PutUint16(bytes[0:2], s.ProtoId)
	binary.LittleEndian.PutUint16(bytes[2:4], s.ProtoType)
	binary.LittleEndian.PutUint32(bytes[4:8], s.DataLen)
	bytes[8] = s.SequenceId
	bytes[9] = s.AckType
	bytes[10] = s.Nouse1
	bytes[11] = s.Nouse2
	return bytes
}

func (s *S1ProtoHeader) Deserialize(bytes []byte) error {
	s.ProtoId = binary.LittleEndian.Uint16(bytes[0:2])
	s.ProtoType = binary.LittleEndian.Uint16(bytes[2:4])
	s.DataLen = binary.LittleEndian.Uint32(bytes[4:8])
	s.SequenceId = bytes[8]
	s.AckType = bytes[9]
	s.Nouse1 = bytes[10]
	s.Nouse2 = bytes[11]
	return nil
}
