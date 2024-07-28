package proto

type S1BattleProto2022IoctrSetLightsRgbType1Ack struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

const S1BattleProto2022IoctrSetLightsRgbType1AckSize = 7

func (s *S1BattleProto2022IoctrSetLightsRgbType1Ack) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetLightsRgbType1AckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetLightsRgbType1Ack) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
