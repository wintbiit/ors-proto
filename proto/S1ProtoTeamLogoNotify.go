package proto

type S1ProtoTeamLogoNotify struct {
	Data    string
	Teamid  int32
	DataLen int32
}

const S1ProtoTeamLogoNotifySize = 8

func (s *S1ProtoTeamLogoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTeamLogoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoTeamLogoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
