package proto

type S1BattleProto2022ArmorParaConfigItem struct {
	TiggerPress    float32
	BulletMaxPress float32
	GolfMinPress   float32
}

const S1BattleProto2022ArmorParaConfigItemSize = 12

func (s *S1BattleProto2022ArmorParaConfigItem) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ArmorParaConfigItemSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ArmorParaConfigItem) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
