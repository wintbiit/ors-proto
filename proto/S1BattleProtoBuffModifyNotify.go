package proto

type S1BattleProtoBuffModifyNotify struct {
	BuffLevel      uint32
	BuffMaxTime    float32
	UpdateProgress byte
	MsgParams      string
	BuffTid        uint32
	BuffUid        uint64
	BuffVisible    byte
	BuffLeftTime   float32
	PlayerUid      uint64
}

const S1BattleProtoBuffModifyNotifySize = 66

func (s *S1BattleProtoBuffModifyNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBuffModifyNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBuffModifyNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
