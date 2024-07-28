package proto

type S1BattleProtoPlayerCommandSetDisconnectNotify struct {
	Value byte
	Uid   uint64
}

const S1BattleProtoPlayerCommandSetDisconnectNotifySize = 9

func (s *S1BattleProtoPlayerCommandSetDisconnectNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerCommandSetDisconnectNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerCommandSetDisconnectNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
