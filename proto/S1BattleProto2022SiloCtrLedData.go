package proto

type S1BattleProto2022SiloCtrLedData struct {
	R byte
	G byte
	B byte
	A byte
}

const S1BattleProto2022SiloCtrLedDataSize = 4

func (s *S1BattleProto2022SiloCtrLedData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SiloCtrLedDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SiloCtrLedData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
