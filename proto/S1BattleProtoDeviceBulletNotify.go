package proto

type S1BattleProtoDeviceBulletNotify struct {
	BulletBottleCount int32
	BulletBottleState []byte
	PlayersUid        []uint64
}

const S1BattleProtoDeviceBulletNotifySize = 4

func (s *S1BattleProtoDeviceBulletNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoDeviceBulletNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoDeviceBulletNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
