package proto

const S1ProtoSetTeamNotifySize = 10

func (s *S1ProtoSetTeamNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTeamNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
