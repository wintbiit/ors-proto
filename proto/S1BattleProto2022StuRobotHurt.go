package proto

type S1BattleProto2022StuRobotHurt struct {
	Cmd      uint16
	Armor    byte
	Hurttype byte
}

const S1BattleProto2022StuRobotHurtSize = 4

func (s *S1BattleProto2022StuRobotHurt) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuRobotHurtSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuRobotHurt) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
