package proto

const S1ProtoTestReqSize = 10240

func (s *S1ProtoTestReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoTestReqSize)
	copy(bytes[0:], s.Test)
	return bytes
}

func (s *S1ProtoTestReq) Deserialize(bytes []byte) error {
	s.Test = string(bytes[0:10240])
	return nil
}
