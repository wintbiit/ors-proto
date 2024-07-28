package proto

type S1BattleProto2022ClientBattleFirstData struct {
	Progress byte
	IsPaused byte
}

const S1BattleProto2022ClientBattleFirstDataSize = 2

func (s *S1BattleProto2022ClientBattleFirstData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBattleFirstDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientBattleFirstData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
