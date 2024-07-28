package proto

const S1ProtoLogoutAckSize = 4

func (s *S1ProtoLogoutAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLogoutAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLogoutAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
