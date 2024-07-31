package proto

const S1ProtoSetReadyReqSize = 1

func (s *S1ProtoSetReadyReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetReadyReqSize)
	bytes[0] = s.State
	return bytes
}

func (s *S1ProtoSetReadyReq) Deserialize(bytes []byte) error {
	s.State = bytes[0]
	return nil
}
