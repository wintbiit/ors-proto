package proto

type S1BattleProto2022RobotCheckinTimestamp struct {
	Magic     uint32
	DataVer   byte
	Timestamp uint32
}

const S1BattleProto2022RobotCheckinTimestampSize = 9

func (s *S1BattleProto2022RobotCheckinTimestamp) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotCheckinTimestampSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotCheckinTimestamp) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
