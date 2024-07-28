package proto

const S1ProtoTestAckSize = 10244

func (s *S1ProtoTestAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoTestAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTestAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
