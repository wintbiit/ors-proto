package proto

type S1BattleProtoPlayerCommandLockScreenNotify struct {
	Value byte
}

const S1BattleProtoPlayerCommandLockScreenNotifySize = 1

func (s *S1BattleProtoPlayerCommandLockScreenNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoPlayerCommandLockScreenNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoPlayerCommandLockScreenNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
