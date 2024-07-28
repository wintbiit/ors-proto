package proto

type S1ProtoSetTeamPlayerNumNotify struct {
	PlayerNum byte
}

const S1ProtoSetTeamPlayerNumNotifySize = 1

func (s *S1ProtoSetTeamPlayerNumNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamPlayerNumNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTeamPlayerNumNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
