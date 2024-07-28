package proto

type S1BattleProto2022ExchangeProRealPosition struct {
	X          int32
	Y          int32
	Pitch      int32
	Yaw        int32
	ModuleType byte
	Z          int32
	Roll       int32
	Seq        uint32
	ModuleId   byte
}

const S1BattleProto2022ExchangeProRealPositionSize = 30

func (s *S1BattleProto2022ExchangeProRealPosition) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ExchangeProRealPositionSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ExchangeProRealPosition) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
