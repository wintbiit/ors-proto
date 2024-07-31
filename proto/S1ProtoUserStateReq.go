package proto

const S1ProtoUserStateReqSize = 64

func (s *S1ProtoUserStateReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoUserStateReqSize)
	copy(bytes[0:], s.Account)
	copy(bytes[32:], s.Password)
	return bytes
}

func (s *S1ProtoUserStateReq) Deserialize(bytes []byte) error {
	s.Account = string(bytes[0:32])
	s.Password = string(bytes[32:64])
	return nil
}
