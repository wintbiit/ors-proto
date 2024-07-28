package proto

type S1ProtoSetTidNotify struct {
	Uid uint64
	Tid uint32
}

const S1ProtoSetTidNotifySize = 12

func (s *S1ProtoSetTidNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTidNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTidNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
