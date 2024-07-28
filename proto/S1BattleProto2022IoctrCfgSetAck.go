package proto

type S1BattleProto2022IoctrCfgSetAck struct {
	ErrCode    byte
	ModuleId   byte
	ModuleType byte
}

const S1BattleProto2022IoctrCfgSetAckSize = 3

func (s *S1BattleProto2022IoctrCfgSetAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrCfgSetAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrCfgSetAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
