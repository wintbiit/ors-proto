package proto

type S1BattleProto2022EnergyStateChangeNotify struct {
	RuneId      byte
	State       byte
	RingSum     byte
	AtkBuffVal  uint16
	DefBuffVal  byte
	BeforeState byte
}

const S1BattleProto2022EnergyStateChangeNotifySize = 7

func (s *S1BattleProto2022EnergyStateChangeNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnergyStateChangeNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnergyStateChangeNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
