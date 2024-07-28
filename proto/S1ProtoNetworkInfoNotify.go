package proto

const S1ProtoNetworkInfoNotifySize = 34

func (s *S1ProtoNetworkInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoNetworkInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoNetworkInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
