package proto

const S1BattleProtoBuffAddNotifySize = 65

func (s *S1BattleProtoBuffAddNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBuffAddNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBuffAddNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
