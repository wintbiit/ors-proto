package proto

type S1BattleProtoPlayerCommandMoveNotify struct {
	Value byte
}

const S1BattleProtoPlayerCommandMoveNotifySize = 1

func (s *S1BattleProtoPlayerCommandMoveNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerCommandMoveNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerCommandMoveNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
