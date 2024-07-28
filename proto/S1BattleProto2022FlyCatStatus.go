package proto

type S1BattleProto2022FlyCatStatus struct {
	SysState    byte
	CtrlState   byte
	WorkState   byte
	Battery     byte
	SensorState byte
}

const S1BattleProto2022FlyCatStatusSize = 5

func (s *S1BattleProto2022FlyCatStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022FlyCatStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022FlyCatStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
