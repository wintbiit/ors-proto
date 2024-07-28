package proto

type S1BattleProto2022RobotMeasureStopReq struct {
	Res byte
}

const S1BattleProto2022RobotMeasureStopReqSize = 1

func (s *S1BattleProto2022RobotMeasureStopReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotMeasureStopReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotMeasureStopReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
