package proto

const S1BattleProto2022StuRobotBuffSize = 6

func (s *S1BattleProto2022StuRobotBuff) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuRobotBuffSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuRobotBuff) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
