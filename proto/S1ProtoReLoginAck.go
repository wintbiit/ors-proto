package proto

const S1ProtoReLoginAckSize = 108

func (s *S1ProtoReLoginAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoReLoginAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
