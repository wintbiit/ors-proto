package proto

type S1BattleProto2022ExchangeProCtrCmdack struct {
	ModuleId   byte
	ErrorCode  byte
	Seq        uint32
	ModuleType byte
}

const S1BattleProto2022ExchangeProCtrCmdackSize = 7

func (s *S1BattleProto2022ExchangeProCtrCmdack) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ExchangeProCtrCmdackSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ExchangeProCtrCmdack) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
