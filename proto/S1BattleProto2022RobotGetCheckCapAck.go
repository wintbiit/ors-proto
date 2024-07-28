package proto

type S1BattleProto2022RobotGetCheckCapAck struct {
	VoltageFirst     float32
	VoltageFinal     float32
	MeasureTime      float32
	OutputQEnergy    float32
	RecharingQEnergy float32
	Date             uint32
	Time             uint32
	CheckedCap       float32
}

const S1BattleProto2022RobotGetCheckCapAckSize = 32

func (s *S1BattleProto2022RobotGetCheckCapAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotGetCheckCapAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotGetCheckCapAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
