package proto

type S1BattleProtoPlayerResetNotify struct {
	PlayerUid uint64
}

const S1BattleProtoPlayerResetNotifySize = 8

func (s *S1BattleProtoPlayerResetNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerResetNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerResetNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
