package proto

type S1BattleProto2022SetFlyCatStatusAck struct {
	RecvedWorkState byte
}

const S1BattleProto2022SetFlyCatStatusAckSize = 1

func (s *S1BattleProto2022SetFlyCatStatusAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SetFlyCatStatusAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SetFlyCatStatusAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
