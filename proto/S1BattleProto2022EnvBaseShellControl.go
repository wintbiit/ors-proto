package proto

type S1BattleProto2022EnvBaseShellControl struct {
	BaseShellControl byte
	Rgb              uint16
}

const S1BattleProto2022EnvBaseShellControlSize = 3

func (s *S1BattleProto2022EnvBaseShellControl) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnvBaseShellControlSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnvBaseShellControl) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
