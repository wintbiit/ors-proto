package proto

type S1BattleProto2022StuRfidstatus struct {
	Cmd         uint16
	BaseArea    uint32
	UplandArea  uint32
	FlyArea     uint32
	IslandArea  uint32
	SapperRfid  uint32
	RuneArea    uint32
	OutpostArea uint32
	HomeArea    uint32
}

const S1BattleProto2022StuRfidstatusSize = 34

func (s *S1BattleProto2022StuRfidstatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuRfidstatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuRfidstatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
