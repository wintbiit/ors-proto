package proto

type S1BattleProto2022IoctrSetLightsRgbType3Ack struct {
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
	ErrCode    byte
}

const S1BattleProto2022IoctrSetLightsRgbType3AckSize = 7

func (s *S1BattleProto2022IoctrSetLightsRgbType3Ack) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetLightsRgbType3AckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetLightsRgbType3Ack) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
