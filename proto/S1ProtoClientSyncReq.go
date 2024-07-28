package proto

const S1ProtoClientSyncReqSize = 29

func (s *S1ProtoClientSyncReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoClientSyncReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoClientSyncReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
