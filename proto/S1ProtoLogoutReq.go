package proto

const S1ProtoLogoutReqSize = 32

func (s *S1ProtoLogoutReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLogoutReqSize)
	copy(bytes[0:], s.Account)
	return bytes
}

func (s *S1ProtoLogoutReq) Deserialize(bytes []byte) error {
	s.Account = string(bytes[0:32])
	return nil
}
