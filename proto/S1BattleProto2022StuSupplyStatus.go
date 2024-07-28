package proto

type S1BattleProto2022StuSupplyStatus struct {
	BulletsNum      byte
	Cmd             uint16
	SupplyEnvId     byte
	SupplyRobotId   byte
	SupplyEnvStatus byte
}

const S1BattleProto2022StuSupplyStatusSize = 6

func (s *S1BattleProto2022StuSupplyStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuSupplyStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuSupplyStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
