package proto

type S1BattleProto2022IoctrSetLightsRgbType2Ack struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

const S1BattleProto2022IoctrSetLightsRgbType2AckSize = 7

func (s *S1BattleProto2022IoctrSetLightsRgbType2Ack) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetLightsRgbType2AckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetLightsRgbType2Ack) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
