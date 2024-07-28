package proto

type S1BattleProtoReduceReliveTimeNotify struct {
	PlayerUid uint64
	TimeMax   float32
	TimeLeft  float32
}

const S1BattleProtoReduceReliveTimeNotifySize = 16

func (s *S1BattleProtoReduceReliveTimeNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoReduceReliveTimeNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoReduceReliveTimeNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
