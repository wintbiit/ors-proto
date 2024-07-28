package proto

const S1ProtoHeartBeatReqSize = 2

func (s *S1ProtoHeartBeatReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoHeartBeatReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoHeartBeatReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
