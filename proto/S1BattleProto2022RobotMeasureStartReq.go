package proto

type S1BattleProto2022RobotMeasureStartReq struct {
	MaxPrriod  uint32
	PushPrriod uint32
}

const S1BattleProto2022RobotMeasureStartReqSize = 8

func (s *S1BattleProto2022RobotMeasureStartReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotMeasureStartReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotMeasureStartReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
