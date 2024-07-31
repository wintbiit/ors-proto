package proto

const S1ProtoSetTokenReqSize = 32

func (s *S1ProtoSetTokenReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTokenReqSize)
	copy(bytes[0:32], s.Token)
	return bytes
}

func (s *S1ProtoSetTokenReq) Deserialize(bytes []byte) error {
	s.Token = string(bytes[0:32])
	return nil
}
