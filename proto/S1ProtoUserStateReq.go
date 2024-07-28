package proto

const S1ProtoUserStateReqSize = 64

func (s *S1ProtoUserStateReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoUserStateReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoUserStateReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
