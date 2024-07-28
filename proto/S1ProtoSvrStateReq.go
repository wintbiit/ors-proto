package proto

const S1ProtoSvrStateReqSize = 1

func (s *S1ProtoSvrStateReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSvrStateReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSvrStateReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
