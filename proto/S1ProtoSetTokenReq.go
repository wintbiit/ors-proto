package proto

const S1ProtoSetTokenReqSize = 32

func (s *S1ProtoSetTokenReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTokenReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTokenReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
