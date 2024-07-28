package proto

type S1BattleProtoPlayerInitOkNotify struct {
	Uid uint64
}

const S1BattleProtoPlayerInitOkNotifySize = 8

func (s *S1BattleProtoPlayerInitOkNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerInitOkNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerInitOkNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
