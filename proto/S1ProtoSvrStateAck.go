package proto

const S1ProtoSvrStateAckSize = 76

func (s *S1ProtoSvrStateAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSvrStateAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSvrStateAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
