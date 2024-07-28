package proto

const S1ProtoTestReqSize = 10240

func (s *S1ProtoTestReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoTestReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTestReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
