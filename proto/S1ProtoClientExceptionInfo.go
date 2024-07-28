package proto

const S1ProtoClientExceptionInfoSize = 4

func (s *S1ProtoClientExceptionInfo) Serialize() []byte {
	bytes := make([]byte, S1ProtoClientExceptionInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoClientExceptionInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
