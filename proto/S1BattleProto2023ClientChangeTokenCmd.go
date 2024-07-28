package proto

type S1BattleProto2023ClientChangeTokenCmd struct {
	NewToken string
	Len      uint16
}

const S1BattleProto2023ClientChangeTokenCmdSize = 258

func (s *S1BattleProto2023ClientChangeTokenCmd) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientChangeTokenCmdSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientChangeTokenCmd) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
