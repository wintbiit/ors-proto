package proto

const S1BattleProto2022ClientServerUimessageSize = 256

func (s *S1BattleProto2022ClientServerUimessage) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientServerUimessageSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientServerUimessage) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
