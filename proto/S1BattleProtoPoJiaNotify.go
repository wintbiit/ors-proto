package proto

type S1BattleProtoPoJiaNotify struct {
	Uid      uint64
	State    byte
	Timeleft float32
	Timemax  float32
}

const S1BattleProtoPoJiaNotifySize = 17

func (s *S1BattleProtoPoJiaNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPoJiaNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPoJiaNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
