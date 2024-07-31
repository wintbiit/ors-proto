package proto

import "encoding/binary"

const S1ProtoTeamCampNotifySize = 4

func (s *S1ProtoTeamCampNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTeamCampNotifySize+s.DataLen)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.DataLen))
	copy(bytes[4:], s.Data)
	return bytes
}

func (s *S1ProtoTeamCampNotify) Deserialize(bytes []byte) error {
	s.DataLen = int32(binary.LittleEndian.Uint32(bytes[0:]))
	s.Data = string(bytes[4:])
	return nil
}
