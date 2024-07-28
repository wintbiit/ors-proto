package proto

const S1ProtoSetTokenAckSize = 4

func (s *S1ProtoSetTokenAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTokenAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTokenAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
