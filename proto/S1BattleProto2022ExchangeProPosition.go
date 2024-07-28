package proto

type S1BattleProto2022ExchangeProPosition struct {
	Y          int32
	Yaw        int32
	Seq        uint32
	ModuleId   byte
	X          int32
	Z          int32
	Pitch      int32
	Roll       int32
	ModuleType byte
}

const S1BattleProto2022ExchangeProPositionSize = 30

func (s *S1BattleProto2022ExchangeProPosition) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ExchangeProPositionSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ExchangeProPosition) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
