package proto

const S1ProtoTeamLineupInfoNotifySize = 4

func (s *S1ProtoTeamLineupInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTeamLineupInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTeamLineupInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
