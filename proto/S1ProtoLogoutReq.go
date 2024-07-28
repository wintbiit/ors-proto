package proto

const S1ProtoLogoutReqSize = 32

func (s *S1ProtoLogoutReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLogoutReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLogoutReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
