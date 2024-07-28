package proto

type S1BattleProto2022EnergyReset struct {
	Res byte
}

const S1BattleProto2022EnergyResetSize = 1

func (s *S1BattleProto2022EnergyReset) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnergyResetSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnergyReset) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
