package proto

const S1BattleProtoPlayerDeadNotifySize = 36

func (s *S1BattleProtoPlayerDeadNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerDeadNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerDeadNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
