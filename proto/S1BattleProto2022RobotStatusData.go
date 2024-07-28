package proto

type S1BattleProto2022RobotStatusData struct {
	Hp             uint16
	GimbalVoltage  float32
	ChassisPower   float32
	PowerBuffer    float32
	ShooterVoltage float32
	PowerState     S1BattleProto2022PowerSwitchState
	GimbalPower    float32
	ShooterPower   float32
	ShooterPitch   int16
	ShooterRoll    int16
	Uwb            S1BattleProto2022UwbData
	Voltage        float32
	Current        float32
	Rssi           byte
	HwId           uint32
	ShooterYaw     int16
	Status         S1BattleProto2022RobotSystemStatus
	TimeStamp      uint64
}

const S1BattleProto2022RobotStatusDataSize = 0

func (s *S1BattleProto2022RobotStatusData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotStatusDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotStatusData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
