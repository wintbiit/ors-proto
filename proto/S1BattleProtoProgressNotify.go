package proto

type S1BattleProtoProgressNotify struct {
	TimeMax   float32
	TimeLeft  float32
	State     byte
	PlayerUid uint64
	EventName string
}

const S1BattleProtoProgressNotifySize = 49

func (s *S1BattleProtoProgressNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoProgressNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoProgressNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
