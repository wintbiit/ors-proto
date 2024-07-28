package proto

type S1BattleProto2022RobotGetCheckCapReq struct {
	Res byte
}

const S1BattleProto2022RobotGetCheckCapReqSize = 1

func (s *S1BattleProto2022RobotGetCheckCapReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotGetCheckCapReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotGetCheckCapReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
