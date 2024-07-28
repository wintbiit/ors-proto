package proto

type S1BattleProtoAttrsNotify struct {
	HpMax         int32
	ArmorMax      int32
	Durability    int32
	Bullet        int32
	Attack2       float32
	DurabilityMax int32
	SeatIndex     byte
	Hp            int32
	Uid           uint64
	Tid           uint32
	Teamid        byte
	Armor         int32
	BulletMax     int32
	Attack1       int32
}

const S1BattleProtoAttrsNotifySize = 54

func (s *S1BattleProtoAttrsNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoAttrsNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoAttrsNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
