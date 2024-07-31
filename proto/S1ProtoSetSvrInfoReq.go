package proto

const S1ProtoSetSvrInfoReqSize = 32

func (s *S1ProtoSetSvrInfoReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetSvrInfoReqSize)
	copy(bytes[0:], s.SvrName)
	return bytes
}

func (s *S1ProtoSetSvrInfoReq) Deserialize(bytes []byte) error {
	s.SvrName = string(bytes[0:32])
	return nil
}
