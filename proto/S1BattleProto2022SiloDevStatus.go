package proto

type S1BattleProto2022SiloDevStatus struct {
	FloorStatus             byte
	Angle                   uint16
	CountdownTime           float32
	IsCanOpen               byte
	SiloErrorCode           byte
	ErrorGateSensorUp       byte
	Online                  byte
	OpenedTimeCount         int32
	ErrorBaseBoardSensor    byte
	ErrorGateSensorBoth     byte
	Target                  byte
	ShootCheckCountdownTime float32
	Speed                   uint16
	ErrorGateSensorDown     byte
	ErrorMotorBroken        byte
	DoorStatus              byte
	MissileCount            int32
	ErrorEmergencyStop      byte
	IsDetectionWindows      byte
	ErrorCloseOverTime      byte
	ErrorOpenOverTime       byte
	SiloCooldown            float32
	MissileHitCount         byte
	ErrorBrakeConstNo       byte
	ErrorBrakeConstYes      byte
	ErrorFacingObstacle     byte
	ErrorMotorOverHeat      byte
}

const S1BattleProto2022SiloDevStatusSize = 44

func (s *S1BattleProto2022SiloDevStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SiloDevStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SiloDevStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
