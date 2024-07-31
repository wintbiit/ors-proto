package proto

const S1ProtoGmcommandReqSize = 512

func (s *S1ProtoGmcommandReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoGmcommandReqSize)
	copy(bytes[0:], s.Command)
	copy(bytes[256:], s.Pars)
	return bytes
}

func (s *S1ProtoGmcommandReq) Deserialize(bytes []byte) error {
	s.Command = string(bytes[0:256])
	s.Pars = string(bytes[256:512])
	return nil
}
