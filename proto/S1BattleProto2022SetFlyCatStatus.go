package proto

type S1BattleProto2022SetFlyCatStatus struct {
	WorkState byte
}

const S1BattleProto2022SetFlyCatStatusSize = 1

func (s *S1BattleProto2022SetFlyCatStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SetFlyCatStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SetFlyCatStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
