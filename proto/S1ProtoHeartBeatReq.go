package proto

const S1ProtoHeartBeatReqSize = 2

func (s *S1ProtoHeartBeatReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoHeartBeatReqSize)
	bytes[0] = s.Nouse
	bytes[1] = s.S0Clientid
	return bytes
}

func (s *S1ProtoHeartBeatReq) Deserialize(bytes []byte) error {
	s.Nouse = bytes[0]
	s.S0Clientid = bytes[1]
	return nil
}
