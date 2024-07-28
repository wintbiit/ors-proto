package proto

type S1BattleProto2022RobotMeasureStartRsp struct {
	Type   byte
	Id     byte
	Status byte
}

const S1BattleProto2022RobotMeasureStartRspSize = 3

func (s *S1BattleProto2022RobotMeasureStartRsp) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotMeasureStartRspSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotMeasureStartRsp) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
