package proto

type S1BattleProto2022ClientRecoverCardStatus struct {
	Redcard  byte
	Bluecard byte
}

const S1BattleProto2022ClientRecoverCardStatusSize = 2

func (s *S1BattleProto2022ClientRecoverCardStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRecoverCardStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRecoverCardStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
