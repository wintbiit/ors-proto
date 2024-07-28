package proto

type S1BattleProto2022StuEnvStatus struct {
	BaseShieldStatus        uint32
	OutpostStatus           uint32
	Cmd                     uint16
	Num1AddBloodPointStatus uint32
	Num3AddBloodPointStatus uint32
	RuneRfidStatus          uint32
	TrapezoidR3             uint32
	TrapezoidR4             uint32
	Num2AddBloodPointStatus uint32
	SmallRuneStatus         uint32
	BigRuneStatus           uint32
	RingUpland              uint32
}

const S1BattleProto2022StuEnvStatusSize = 46

func (s *S1BattleProto2022StuEnvStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuEnvStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuEnvStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
