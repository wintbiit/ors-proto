package proto

type S1BattleProtoBuffDelNotify struct {
	PlayerUid uint64
	BuffUid   uint64
}

const S1BattleProtoBuffDelNotifySize = 16

func (s *S1BattleProtoBuffDelNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBuffDelNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBuffDelNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
