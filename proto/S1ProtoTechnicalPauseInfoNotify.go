package proto

type S1ProtoTechnicalPauseInfoNotify struct {
	BlueLongPauseLeftCount  byte
	PausedDuration          uint32
	PauseTimeType           byte
	PauseSide               byte
	RedShortPauseLeftCount  byte
	RedLongPauseLeftCount   byte
	BlueShortPauseLeftCount byte
}

const S1ProtoTechnicalPauseInfoNotifySize = 10

func (s *S1ProtoTechnicalPauseInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTechnicalPauseInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTechnicalPauseInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
