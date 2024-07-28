package proto

const S1ProtoSetCurrMatchNotifySize = 160

func (s *S1ProtoSetCurrMatchNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetCurrMatchNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetCurrMatchNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
