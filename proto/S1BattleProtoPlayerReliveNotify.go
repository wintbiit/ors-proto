package proto

type S1BattleProtoPlayerReliveNotify struct {
	Uid uint64
}

const S1BattleProtoPlayerReliveNotifySize = 8

func (s *S1BattleProtoPlayerReliveNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerReliveNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerReliveNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
