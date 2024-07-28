package proto

type S1BattleProtoIsCanUsePjnotify struct {
	Teamcount int32
	Teamid    []int32
	State     []int32
}

const S1BattleProtoIsCanUsePjnotifySize = 4

func (s *S1BattleProtoIsCanUsePjnotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoIsCanUsePjnotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoIsCanUsePjnotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
