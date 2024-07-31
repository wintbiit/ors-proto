package proto

const S1ProtoSvrStateReqSize = 1

func (s *S1ProtoSvrStateReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSvrStateReqSize)
	bytes[0] = s.Nouse
	return bytes
}

func (s *S1ProtoSvrStateReq) Deserialize(bytes []byte) error {
	s.Nouse = bytes[0]
	return nil
}
