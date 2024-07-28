package proto

const S1ProtoPlayerCommandUserInputNotifySize = 1

func (s *S1ProtoPlayerCommandUserInputNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoPlayerCommandUserInputNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoPlayerCommandUserInputNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
