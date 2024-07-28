package proto

type S1BattleProto2022SiloStatus struct {
	DoorStatus  byte
	FloorStatus byte
	Errorcode   byte
}

const S1BattleProto2022SiloStatusSize = 3

func (s *S1BattleProto2022SiloStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SiloStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SiloStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
