package proto

type S1BattleProtoAddBulletNotify struct {
	Uid         uint64
	BulletIndex byte
}

const S1BattleProtoAddBulletNotifySize = 9

func (s *S1BattleProtoAddBulletNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoAddBulletNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoAddBulletNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
