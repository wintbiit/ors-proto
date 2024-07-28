package proto

type S1BattleProto2022RobotPushCapinfo struct {
	MeasureTime      float32
	OutputQEnergy    float32
	RecharingQEnergy float32
	MeasureCap       float32
	CheckCap         float32
	VoltageFirst     float32
	VoltageFinal     float32
}

const S1BattleProto2022RobotPushCapinfoSize = 28

func (s *S1BattleProto2022RobotPushCapinfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotPushCapinfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotPushCapinfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
