package proto

const S1BattleProto2022GameLimitSize = 20

func (s *S1BattleProto2022GameLimit) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022GameLimitSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022GameLimit) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
