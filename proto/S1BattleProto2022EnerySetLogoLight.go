package proto

type S1BattleProto2022EnerySetLogoLight struct {
	Row        byte
	IconRColor uint32
}

const S1BattleProto2022EnerySetLogoLightSize = 5

func (s *S1BattleProto2022EnerySetLogoLight) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnerySetLogoLightSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnerySetLogoLight) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
