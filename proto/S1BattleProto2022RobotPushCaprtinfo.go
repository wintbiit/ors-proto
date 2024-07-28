package proto

type S1BattleProto2022RobotPushCaprtinfo struct {
	OutputCurrent     float32
	RechargingCurrent float32
	VoltageFirst      float32
	MeasureTime       float32
	MeasureCap        float32
	RecharingQEnergy  float32
	OutputQEnergy     float32
	Voltage           float32
}

const S1BattleProto2022RobotPushCaprtinfoSize = 32

func (s *S1BattleProto2022RobotPushCaprtinfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotPushCaprtinfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotPushCaprtinfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
