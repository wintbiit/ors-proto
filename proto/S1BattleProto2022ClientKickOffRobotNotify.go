package proto

type S1BattleProto2022ClientKickOffRobotNotify struct {
	RobotId int32
	Reason  byte
}

const S1BattleProto2022ClientKickOffRobotNotifySize = 5

func (s *S1BattleProto2022ClientKickOffRobotNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientKickOffRobotNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientKickOffRobotNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
