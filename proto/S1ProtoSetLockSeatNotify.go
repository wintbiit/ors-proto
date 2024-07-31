package proto

const S1ProtoSetLockSeatNotifySize = 1

func (s *S1ProtoSetLockSeatNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetLockSeatNotifySize)
	bytes[0] = s.LockSeat
	return bytes
}

func (s *S1ProtoSetLockSeatNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoSetLockSeatNotifySize {
		return InvalidDataError
	}

	s.LockSeat = bytes[0]
	return nil
}
