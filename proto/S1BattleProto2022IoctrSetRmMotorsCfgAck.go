package proto

type S1BattleProto2022IoctrSetRmMotorsCfgAck struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

const S1BattleProto2022IoctrSetRmMotorsCfgAckSize = 7

func (s *S1BattleProto2022IoctrSetRmMotorsCfgAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetRmMotorsCfgAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetRmMotorsCfgAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
