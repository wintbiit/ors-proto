package proto

const S1BattleProtoBeAttackNotifySize = 28

func (s *S1BattleProtoBeAttackNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBeAttackNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBeAttackNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
