package proto

type S1BattleProto2022RobotNtpServerIpAck struct {
	Ip uint32
}

const S1BattleProto2022RobotNtpServerIpAckSize = 4

func (s *S1BattleProto2022RobotNtpServerIpAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotNtpServerIpAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotNtpServerIpAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
