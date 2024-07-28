package proto

type S1BattleProto2022RobotDataSync struct {
	InventedShieldValue   uint16
	LightEffectMask       uint32
	FixPower              uint16
	FixSamllShootSpeed2   uint16
	ShooterHeats          []float32
	MaxHp                 int32
	Level                 byte
	ShooterHeatsCount     int32
	ShooterHeatThres      []float32
	CurHeatCoolCount      int32
	CurHeatCool           []float32
	FixSamllShootSpeed1   uint16
	CurHp                 int32
	FixBigShootSpeed      uint16
	IsGuardAlive          uint16
	ShooterHeatThresCount int32
}

const S1BattleProto2022RobotDataSyncSize = 73

func (s *S1BattleProto2022RobotDataSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotDataSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotDataSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
