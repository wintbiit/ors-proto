package proto

type S1BattleProto2022RobotPushSensorInfo struct {
	OutputCurrent     float32
	RechargingCurrent float32
	Status            byte
	Voltage           float32
}

const S1BattleProto2022RobotPushSensorInfoSize = 13

func (s *S1BattleProto2022RobotPushSensorInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotPushSensorInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotPushSensorInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
