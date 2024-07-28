package proto

type S1BattleProto2022ClientExceptionRecoverCardNotify struct {
	Teamid    byte
	Robotid   byte
	Exception byte
}

const S1BattleProto2022ClientExceptionRecoverCardNotifySize = 3

func (s *S1BattleProto2022ClientExceptionRecoverCardNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientExceptionRecoverCardNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientExceptionRecoverCardNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
