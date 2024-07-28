package proto

type S1BattleProto2022EnergySetId struct {
	Res byte
}

const S1BattleProto2022EnergySetIdSize = 1

func (s *S1BattleProto2022EnergySetId) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnergySetIdSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnergySetId) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
