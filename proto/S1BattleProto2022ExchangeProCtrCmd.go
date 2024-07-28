package proto

type S1BattleProto2022ExchangeProCtrCmd struct {
	Seq        uint32
	ModuleType byte
	ModuleId   byte
	Command    byte
	Reserved1  byte
	Reserved2  byte
}

const S1BattleProto2022ExchangeProCtrCmdSize = 9

func (s *S1BattleProto2022ExchangeProCtrCmd) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ExchangeProCtrCmdSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ExchangeProCtrCmd) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
