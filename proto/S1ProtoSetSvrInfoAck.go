package proto

const S1ProtoSetSvrInfoAckSize = 4

func (s *S1ProtoSetSvrInfoAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetSvrInfoAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetSvrInfoAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
