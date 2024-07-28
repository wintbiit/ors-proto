package proto

type S1BattleProto2022IoctrSetRmMotorsCfg struct {
	IsStatusReturnListLen int32
	Seq                   uint32
	ModuleId              byte
	SpeedPidListLen       int32
	PositionPidListLen    int32
	PositionPids          []S1BattleProto2022_PositionModePid
	TransRatioListLen     int32
	TransRatios           []float32
	SpeedPids             []S1BattleProto2022_SpeedModePid
	IsStatusReturn        []byte
	ModuleType            byte
}

const S1BattleProto2022IoctrSetRmMotorsCfgSize = 310

func (s *S1BattleProto2022IoctrSetRmMotorsCfg) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetRmMotorsCfgSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetRmMotorsCfg) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
