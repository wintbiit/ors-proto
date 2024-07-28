package proto

type S1BattleProto2022ShooterFreqLimit struct {
	ShooterFreqThreshold    float32
	ShooterFreqMaxHurt      uint16
	ShooterFreqHurtTabCount int32
	ShooterFreqHurtTab      []byte
}

const S1BattleProto2022ShooterFreqLimitSize = 22

func (s *S1BattleProto2022ShooterFreqLimit) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ShooterFreqLimitSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ShooterFreqLimit) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
