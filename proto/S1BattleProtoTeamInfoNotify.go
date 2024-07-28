package proto

type S1BattleProtoTeamInfoNotify struct {
	DataLen int32
	Data    string
}

const S1BattleProtoTeamInfoNotifySize = 4

func (s *S1BattleProtoTeamInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoTeamInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoTeamInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
