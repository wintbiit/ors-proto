package proto

const S1ProtoLoginAckSize = 44

func (s *S1ProtoLoginAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLoginAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLoginAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
