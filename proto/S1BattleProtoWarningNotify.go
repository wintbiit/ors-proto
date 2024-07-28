package proto

const S1BattleProtoWarningNotifySize = 9

func (s *S1BattleProtoWarningNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoWarningNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoWarningNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
