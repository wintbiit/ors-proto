package proto

type S1BattleProto2022IoctrSetLightsRgbType2 struct {
	ListLen       int32
	LedsColors    []S1BattleProto2022_LedColor
	Seq           uint32
	ModuleId      byte
	ModuleType    byte
	StartLedIndex uint16
	LedNum        uint16
}

const S1BattleProto2022IoctrSetLightsRgbType2Size = 398

func (s *S1BattleProto2022IoctrSetLightsRgbType2) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetLightsRgbType2Size)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetLightsRgbType2) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
