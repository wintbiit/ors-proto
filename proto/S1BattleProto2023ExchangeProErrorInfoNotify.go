package proto

type S1BattleProto2023ExchangeProErrorInfoNotify struct {
	MotorPitchState          byte
	IsHardwareEmergencyState byte
	IsOnline                 byte
	ArmErrorCode             uint16
	IsSoftwareEmergencyState byte
	MotorPushState           byte
	Ir1State                 byte
	Ir3State                 byte
	IrErrorCode              byte
	Ir2State                 byte
	TempAlertVal             int16
	TempRmMotorRoll          int16
	TempRmMotorPitch         int16
	TempRmMotorPush          int16
	TeamId                   byte
	RfidErrorCode            byte
	RunningState             byte
	MotorRollState           byte
}

const S1BattleProto2023ExchangeProErrorInfoNotifySize = 23

func (s *S1BattleProto2023ExchangeProErrorInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ExchangeProErrorInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ExchangeProErrorInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
