package proto

const S1BattleProto2022ClientBuffItemSize = 5

func (s *S1BattleProto2022ClientBuffItem) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBuffItemSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientBuffItem) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
