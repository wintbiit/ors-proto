package proto

type S1ProtoSetTidReq struct {
	Tid   uint32
	Token string
}

const S1ProtoSetTidReqSize = 36

func (s *S1ProtoSetTidReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTidReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTidReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
