package proto

const S1ProtoGmcommandAckSize = 4

func (s *S1ProtoGmcommandAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoGmcommandAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoGmcommandAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
