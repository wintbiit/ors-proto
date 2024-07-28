package proto

type S1BattleProto2023ClientPenaltyInfo struct {
	RobotId       byte
	PenaltyType   byte
	PenaltyReason string
}

const S1BattleProto2023ClientPenaltyInfoSize = 130

func (s *S1BattleProto2023ClientPenaltyInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientPenaltyInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientPenaltyInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
