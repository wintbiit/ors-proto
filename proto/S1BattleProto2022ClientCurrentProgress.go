package proto

type S1BattleProto2022ClientCurrentProgress struct {
	CurProgress int32
	IsPaused    byte
}

const S1BattleProto2022ClientCurrentProgressSize = 5

func (s *S1BattleProto2022ClientCurrentProgress) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientCurrentProgressSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientCurrentProgress) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
