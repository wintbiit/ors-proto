package proto

type S1BattleProtoIsCanUseAtknotify struct {
	Teamid    []int32
	State     []int32
	Teamcount int32
}

const S1BattleProtoIsCanUseAtknotifySize = 4

func (s *S1BattleProtoIsCanUseAtknotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoIsCanUseAtknotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoIsCanUseAtknotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
