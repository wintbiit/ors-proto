package proto

const S1ProtoPingReqSize = 1

func (s *S1ProtoPingReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoPingReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoPingReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
