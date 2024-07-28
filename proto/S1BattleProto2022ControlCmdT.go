package proto

type S1BattleProto2022ControlCmdT struct {
	OptCode uint16
	Data    uint16
}

const S1BattleProto2022ControlCmdTSize = 4

func (s *S1BattleProto2022ControlCmdT) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ControlCmdTSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ControlCmdT) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
