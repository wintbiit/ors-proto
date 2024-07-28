package proto

type S1ProtoSetReadyNotify struct {
	Uid   uint64
	State byte
}

const S1ProtoSetReadyNotifySize = 9

func (s *S1ProtoSetReadyNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetReadyNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetReadyNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
