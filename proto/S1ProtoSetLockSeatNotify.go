package proto

const S1ProtoSetLockSeatNotifySize = 1

func (s *S1ProtoSetLockSeatNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetLockSeatNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetLockSeatNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
