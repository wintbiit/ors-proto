package proto

type S1BattleProto2022Energy2019SetArmLight struct {
	ExtVar     []uint16
	Row        byte
	ArmorColor []uint32
	ArmColor   []uint32
	Effect     []byte
}

const S1BattleProto2022Energy2019SetArmLightSize = 56

func (s *S1BattleProto2022Energy2019SetArmLight) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022Energy2019SetArmLightSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022Energy2019SetArmLight) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
