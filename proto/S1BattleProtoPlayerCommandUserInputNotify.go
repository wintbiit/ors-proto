package proto

type S1BattleProtoPlayerCommandUserInputNotify struct {
	Value byte
}

const S1BattleProtoPlayerCommandUserInputNotifySize = 1

func (s *S1BattleProtoPlayerCommandUserInputNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerCommandUserInputNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerCommandUserInputNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
