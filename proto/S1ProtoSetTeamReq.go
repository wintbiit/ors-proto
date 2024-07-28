package proto

const S1ProtoSetTeamReqSize = 2

func (s *S1ProtoSetTeamReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTeamReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
