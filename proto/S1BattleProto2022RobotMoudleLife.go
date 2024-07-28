package proto

type S1BattleProto2022RobotMoudleLife struct {
	Status        string
	ModuleId      byte
	ModuleType    byte
	AckModuleId   byte
	AckModuleType byte
}

const S1BattleProto2022RobotMoudleLifeSize = 260

func (s *S1BattleProto2022RobotMoudleLife) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotMoudleLifeSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotMoudleLife) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
