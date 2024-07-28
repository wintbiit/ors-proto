package proto

type S1BattleProto2022ClientOutpostDeviceStatusSync struct {
	MotorCurrentWarning   byte
	MotorSpeedoverWarning byte
	ActionTimeout         byte
	DartBoardIrOnline     byte
	TemperatureWarning    byte
	SpinSpeedRatio        float32
	DartBoardBrokenErr    byte
	Online                byte
	CurTemperature        int32
	MotorOnline           byte
	IsMagnetOn            byte
	IrledStatusRight      uint32
	RobotId               byte
	StatusCode            byte
	IrledStatusLeft       uint32
}

const S1BattleProto2022ClientOutpostDeviceStatusSyncSize = 27

func (s *S1BattleProto2022ClientOutpostDeviceStatusSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientOutpostDeviceStatusSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientOutpostDeviceStatusSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
