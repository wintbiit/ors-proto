package proto

type S1BattleProtoBuffAddNotify struct {
	BuffUid      uint64
	BuffTid      uint32
	BuffLevel    uint32
	BuffVisible  byte
	BuffMaxTime  float32
	BuffLeftTime float32
	MsgParams    string
	PlayerUid    uint64
}

const S1BattleProtoBuffAddNotifySize = 65

func (s *S1BattleProtoBuffAddNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBuffAddNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBuffAddNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
