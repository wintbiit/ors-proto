package proto

type S1BattleProto2022IoctrSetLightsRgbType1 struct {
	StartLedIndex uint16
	LedNum        uint16
	Color         S1BattleProto2022_LedColor
	Seq           uint32
	ModuleId      byte
	ModuleType    byte
}

const S1BattleProto2022IoctrSetLightsRgbType1Size = 13

func (s *S1BattleProto2022IoctrSetLightsRgbType1) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetLightsRgbType1Size)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetLightsRgbType1) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
