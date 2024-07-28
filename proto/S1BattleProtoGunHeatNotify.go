package proto

type S1BattleProtoGunHeatNotify struct {
	PlayerUid  uint64
	GunHeatMax float32
	GunHeat    float32
}

const S1BattleProtoGunHeatNotifySize = 16

func (s *S1BattleProtoGunHeatNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGunHeatNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGunHeatNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
