package proto

type S1BattleProtoPlaceholderNotify struct {
	EventName string
	State     byte
}

const S1BattleProtoPlaceholderNotifySize = 33

func (s *S1BattleProtoPlaceholderNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlaceholderNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlaceholderNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
