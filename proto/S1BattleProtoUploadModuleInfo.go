package proto

const S1BattleProtoUploadModuleInfoSize = 4

func (s *S1BattleProtoUploadModuleInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoUploadModuleInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoUploadModuleInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
