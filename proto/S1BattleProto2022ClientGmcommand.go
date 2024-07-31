package proto

const S1BattleProto2022ClientGmcommandSize = 129

func (s *S1BattleProto2022ClientGmcommand) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientGmcommandSize)
	bytes[0] = s.Len
	copy(bytes[1:], s.Cmd)
	return bytes
}

func (s *S1BattleProto2022ClientGmcommand) Deserialize(bytes []byte) error {
	s.Len = bytes[0]
	s.Cmd = string(bytes[1:129])
	return nil
}
