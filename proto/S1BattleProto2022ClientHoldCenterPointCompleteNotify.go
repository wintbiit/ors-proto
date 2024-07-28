package proto

type S1BattleProto2022ClientHoldCenterPointCompleteNotify struct {
	Teamid    int32
	RewardExp float32
}

const S1BattleProto2022ClientHoldCenterPointCompleteNotifySize = 8

func (s *S1BattleProto2022ClientHoldCenterPointCompleteNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientHoldCenterPointCompleteNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientHoldCenterPointCompleteNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
