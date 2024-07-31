package proto

const S1ProtoSetTeamReqSize = 2

func (s *S1ProtoSetTeamReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamReqSize)
	bytes[0] = s.TeamId
	bytes[1] = s.SeatIndex
	return bytes
}

func (s *S1ProtoSetTeamReq) Deserialize(bytes []byte) error {
	s.TeamId = bytes[0]
	s.SeatIndex = bytes[1]
	return nil
}
