package proto

type S1BattleProtoBaseArmorStateNotify struct {
	ArmorCount int16
	ArmorValue int16
	PlayerUid  uint64
	ArmorId    int16
}

const S1BattleProtoBaseArmorStateNotifySize = 14

func (s *S1BattleProtoBaseArmorStateNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoBaseArmorStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoBaseArmorStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
