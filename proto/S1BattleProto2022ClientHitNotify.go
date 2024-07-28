package proto

type S1BattleProto2022ClientHitNotify struct {
	ArmorNumber byte
	HpReduce    int32
	HpCurr      int32
	HpMax       int32
	RobotId     byte
	OnHitType   byte
}

const S1BattleProto2022ClientHitNotifySize = 15

func (s *S1BattleProto2022ClientHitNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientHitNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientHitNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
