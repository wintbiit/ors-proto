package proto

const S1ProtoClientNetworkInfoNotifySize = 9

func (s *S1ProtoClientNetworkInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoClientNetworkInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoClientNetworkInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
