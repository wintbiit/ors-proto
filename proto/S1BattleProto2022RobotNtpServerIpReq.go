package proto

type S1BattleProto2022RobotNtpServerIpReq struct {
	Res byte
}

const S1BattleProto2022RobotNtpServerIpReqSize = 1

func (s *S1BattleProto2022RobotNtpServerIpReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotNtpServerIpReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotNtpServerIpReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
