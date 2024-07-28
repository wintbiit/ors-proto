package proto

const S1ProtoGmcommandReqSize = 512

func (s *S1ProtoGmcommandReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoGmcommandReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoGmcommandReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
