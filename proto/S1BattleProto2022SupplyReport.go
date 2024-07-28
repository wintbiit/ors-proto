package proto

type S1BattleProto2022SupplyReport struct {
	BulletGear0   byte
	Loadselector0 byte
	Box2          byte
	State         byte
	BoxReadyNum   byte
	BulletGear1   byte
	Box0          byte
	Box1          byte
	Relase        byte
	Scales        byte
	BoxFreedNum   byte
}

const S1BattleProto2022SupplyReportSize = 11

func (s *S1BattleProto2022SupplyReport) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SupplyReportSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SupplyReport) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
