package proto

const S1ProtoSetTidAckSize = 8

func (s *S1ProtoSetTidAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTidAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTidAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
