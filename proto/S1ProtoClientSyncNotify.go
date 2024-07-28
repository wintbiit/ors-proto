package proto

const S1ProtoClientSyncNotifySize = 29

func (s *S1ProtoClientSyncNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoClientSyncNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoClientSyncNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
