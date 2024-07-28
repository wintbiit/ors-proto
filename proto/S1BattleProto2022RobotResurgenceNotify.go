package proto

type S1BattleProto2022RobotResurgenceNotify struct {
	RobotId byte
}

const S1BattleProto2022RobotResurgenceNotifySize = 1

func (s *S1BattleProto2022RobotResurgenceNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotResurgenceNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotResurgenceNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
