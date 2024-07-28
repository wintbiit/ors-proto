package proto

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
