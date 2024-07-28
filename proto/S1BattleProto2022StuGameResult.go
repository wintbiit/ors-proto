package proto

type S1BattleProto2022StuGameResult struct {
	Cmd    uint16
	Winner byte
}

const S1BattleProto2022StuGameResultSize = 3

func (s *S1BattleProto2022StuGameResult) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuGameResultSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuGameResult) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
