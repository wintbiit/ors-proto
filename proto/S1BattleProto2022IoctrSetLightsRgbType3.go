package proto

type S1BattleProto2022IoctrSetLightsRgbType3 struct {
	LightEffect byte
	TimeSpan    uint16
	Seq         uint32
	ModuleId    byte
	ModuleType  byte
}

const S1BattleProto2022IoctrSetLightsRgbType3Size = 9

func (s *S1BattleProto2022IoctrSetLightsRgbType3) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetLightsRgbType3Size)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetLightsRgbType3) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
