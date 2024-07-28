package proto

type S1BattleProto2022ClientArmorLifeInfo struct {
	LifeState byte
	ModuleId  int16
}

const S1BattleProto2022ClientArmorLifeInfoSize = 3

func (s *S1BattleProto2022ClientArmorLifeInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientArmorLifeInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientArmorLifeInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
