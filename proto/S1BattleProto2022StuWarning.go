package proto

type S1BattleProto2022StuWarning struct {
	Cmd          uint16
	WarningLevel byte
	WarningRobot byte
}

const S1BattleProto2022StuWarningSize = 4

func (s *S1BattleProto2022StuWarning) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuWarningSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuWarning) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
