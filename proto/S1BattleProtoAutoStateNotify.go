package proto

type S1BattleProtoAutoStateNotify struct {
	State int32
}

const S1BattleProtoAutoStateNotifySize = 4

func (s *S1BattleProtoAutoStateNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoAutoStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoAutoStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
