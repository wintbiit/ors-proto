package proto

type S1BattleProto2022ClientShowMessage struct {
	Teamid    byte
	MsgType   byte
	MsgEnum   uint32
	MsgParams string
}

const S1BattleProto2022ClientShowMessageSize = 262

func (s *S1BattleProto2022ClientShowMessage) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientShowMessageSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientShowMessage) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
