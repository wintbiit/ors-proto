package proto

const S1ProtoSetReadyReqSize = 1

func (s *S1ProtoSetReadyReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetReadyReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetReadyReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
