package proto

type S1BattleProto2023ClientModuleVersionData struct {
	MoudleId          byte
	MoudleType        byte
	MoudleSubType     byte
	NewMoudleVersion  string
	CurMoudleVersion  string
	VersionState      byte
	MoudleIsImportant byte
}

const S1BattleProto2023ClientModuleVersionDataSize = 69

func (s *S1BattleProto2023ClientModuleVersionData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientModuleVersionDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientModuleVersionData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
