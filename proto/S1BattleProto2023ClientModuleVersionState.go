package proto

type S1BattleProto2023ClientModuleVersionState struct {
	RobotId       byte
	ModuleNum     byte
	ModuleNumMax  byte
	ModuleVersion []S1BattleProto2023ClientModuleVersionData
}

const S1BattleProto2023ClientModuleVersionStateSize = 1659

func (s *S1BattleProto2023ClientModuleVersionState) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientModuleVersionStateSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientModuleVersionState) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
