package proto

const S1ProtoTeamCampNotifySize = 4

func (s *S1ProtoTeamCampNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTeamCampNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTeamCampNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
