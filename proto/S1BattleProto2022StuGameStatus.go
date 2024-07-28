package proto

type S1BattleProto2022StuGameStatus struct {
	GameMode      byte
	GameStatus    byte
	RemainTime    uint16
	SyncTimeStamp uint64
	Cmd           uint16
}

const S1BattleProto2022StuGameStatusSize = 14

func (s *S1BattleProto2022StuGameStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuGameStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuGameStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
