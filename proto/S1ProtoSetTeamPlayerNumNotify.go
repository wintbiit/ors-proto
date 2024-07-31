package proto

const S1ProtoSetTeamPlayerNumNotifySize = 1

func (s *S1ProtoSetTeamPlayerNumNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamPlayerNumNotifySize)
	bytes[0] = s.PlayerNum
	return bytes
}

func (s *S1ProtoSetTeamPlayerNumNotify) Deserialize(bytes []byte) error {
	s.PlayerNum = bytes[0]
	return nil
}
