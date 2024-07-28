package proto

type S1BattleProtoGunFireNotify struct {
	AttackerUid uint64
	GunType     byte
	GunSpeed    float32
}

const S1BattleProtoGunFireNotifySize = 13

func (s *S1BattleProtoGunFireNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGunFireNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGunFireNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
