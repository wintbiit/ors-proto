package proto

const S1ProtoTeamInfoNotifySize = 4

func (s *S1ProtoTeamInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTeamInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTeamInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
