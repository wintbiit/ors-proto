package proto

const S1ProtoMarkDetectResultAckSize = 25

func (s *S1ProtoMarkDetectResultAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoMarkDetectResultAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoMarkDetectResultAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
