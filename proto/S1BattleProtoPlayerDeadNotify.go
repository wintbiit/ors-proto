package proto

type S1BattleProtoPlayerDeadNotify struct {
	Flag2     int32
	Uid       uint64
	KillerUid uint64
	KillHonor int32
	DeadType  int32
	Flag0     int32
	Flag1     int32
}

const S1BattleProtoPlayerDeadNotifySize = 36

func (s *S1BattleProtoPlayerDeadNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerDeadNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerDeadNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
