package proto

type S1BattleProto2022EnergyStateSync struct {
	Armor3            byte
	Armor5            byte
	Motor             byte
	RuneId            byte
	Round             byte
	Time              byte
	AvailbleCountdown byte
	State             byte
	Armor0            byte
	Armor2            byte
	Armor1            byte
	Armor4            byte
	Connect           byte
	Error             byte
}

const S1BattleProto2022EnergyStateSyncSize = 14

func (s *S1BattleProto2022EnergyStateSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnergyStateSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnergyStateSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
