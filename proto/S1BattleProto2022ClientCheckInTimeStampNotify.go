package proto

type S1BattleProto2022ClientCheckInTimeStampNotify struct {
	RobotidList    []uint32
	RobotidListLen byte
}

const S1BattleProto2022ClientCheckInTimeStampNotifySize = 89

func (s *S1BattleProto2022ClientCheckInTimeStampNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientCheckInTimeStampNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientCheckInTimeStampNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
