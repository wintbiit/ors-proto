package proto

type S1BattleProtoBeAttackNotify struct {
	Index         int32
	AttackType    int32
	AttackerUid   uint64
	BeAttackerUid uint64
	Damage        int32
}

const S1BattleProtoBeAttackNotifySize = 28

func (s *S1BattleProtoBeAttackNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBeAttackNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBeAttackNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
