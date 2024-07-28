package proto

const S1ProtoSetWifiInfoNotifySize = 64

func (s *S1ProtoSetWifiInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetWifiInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetWifiInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
