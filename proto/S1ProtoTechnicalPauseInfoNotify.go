package proto

import "encoding/binary"

const S1ProtoTechnicalPauseInfoNotifySize = 10

func (s *S1ProtoTechnicalPauseInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTechnicalPauseInfoNotifySize)
	bytes[0] = s.PauseTimeType
	bytes[1] = s.PauseSide
	bytes[2] = s.RedShortPauseLeftCount
	bytes[3] = s.RedLongPauseLeftCount
	bytes[4] = s.BlueShortPauseLeftCount
	bytes[5] = s.BlueLongPauseLeftCount
	binary.LittleEndian.PutUint32(bytes[6:], s.PausedDuration)
	return bytes
}

func (s *S1ProtoTechnicalPauseInfoNotify) Deserialize(bytes []byte) error {
	s.PauseTimeType = bytes[0]
	s.PauseSide = bytes[1]
	s.RedShortPauseLeftCount = bytes[2]
	s.RedLongPauseLeftCount = bytes[3]
	s.BlueShortPauseLeftCount = bytes[4]
	s.BlueLongPauseLeftCount = bytes[5]
	s.PausedDuration = binary.LittleEndian.Uint32(bytes[6:])
	return nil
}
