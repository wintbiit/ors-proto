package proto

type S1BattleProto2022Energy2019StateReport struct {
	State       byte
	Armors      []byte
	Reserve     byte
	RotateDeg   uint16
	RotateSpeed float32
}

const S1BattleProto2022Energy2019StateReportSize = 13

func (s *S1BattleProto2022Energy2019StateReport) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022Energy2019StateReportSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022Energy2019StateReport) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
