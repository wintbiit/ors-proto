package proto

type S1BattleProto2022Energy2019AmorSelfCheck struct {
	Nouse byte
}

const S1BattleProto2022Energy2019AmorSelfCheckSize = 1

func (s *S1BattleProto2022Energy2019AmorSelfCheck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022Energy2019AmorSelfCheckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022Energy2019AmorSelfCheck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
