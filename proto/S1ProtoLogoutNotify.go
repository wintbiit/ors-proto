package proto

type S1ProtoLogoutNotify struct {
	Uid    uint64
	Reason byte
}

const S1ProtoLogoutNotifySize = 9

func (s *S1ProtoLogoutNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoLogoutNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLogoutNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
