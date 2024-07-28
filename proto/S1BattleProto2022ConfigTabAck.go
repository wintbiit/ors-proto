package proto

type S1BattleProto2022ConfigTabAck struct {
	ModuleNum          S1BattleProto2022ModuleNum
	HealthCalc         S1BattleProto2022HealthCalc
	GameLimit          S1BattleProto2022GameLimit
	ArmorData          S1BattleProtot2022RobotArmorCfgData
	Magic              uint32
	RobotConfigVersion byte
	Color              byte
}

const S1BattleProto2022ConfigTabAckSize = 0

func (s *S1BattleProto2022ConfigTabAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ConfigTabAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ConfigTabAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
