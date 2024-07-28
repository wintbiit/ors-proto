package proto

type S1BattleProtoGunBulletNotify struct {
	Uid          uint64
	GunBulletMax int16
	GunBullet    int16
}

const S1BattleProtoGunBulletNotifySize = 12

func (s *S1BattleProtoGunBulletNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGunBulletNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGunBulletNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
