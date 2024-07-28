package proto

const S1ProtoSetSvrInfoReqSize = 32

func (s *S1ProtoSetSvrInfoReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetSvrInfoReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetSvrInfoReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
