package proto

type S1BattleProto2023ClientGameSystemInfoNotify struct {
	TokenLen      uint16
	CurMatchToken string
}

const S1BattleProto2023ClientGameSystemInfoNotifySize = 258

func (s *S1BattleProto2023ClientGameSystemInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientGameSystemInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientGameSystemInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
