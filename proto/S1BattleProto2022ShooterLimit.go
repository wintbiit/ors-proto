package proto

type S1BattleProto2022ShooterLimit struct {
	ShooterSpdThreshold    float32
	ShooterSpdMaxHurt      uint16
	ShooterSqdHurtTabCount int32
	ShooterSqdHurtTab      []byte
}

const S1BattleProto2022ShooterLimitSize = 34

func (s *S1BattleProto2022ShooterLimit) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ShooterLimitSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ShooterLimit) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
