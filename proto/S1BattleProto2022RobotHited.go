package proto

type S1BattleProto2022RobotHited struct {
	OnHitType   byte
	AttackSpeed uint16
	ArmorNumber byte
	Press       float32
	TimeStamp   uint64
}

const S1BattleProto2022RobotHitedSize = 16

func (s *S1BattleProto2022RobotHited) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotHitedSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotHited) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
