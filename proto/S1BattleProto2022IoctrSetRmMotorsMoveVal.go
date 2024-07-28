package proto

type S1BattleProto2022IoctrSetRmMotorsMoveVal struct {
	LoopMode        []byte
	Len             int32
	ExpectVal       []int64
	ModuleId        byte
	ModuleType      byte
	LoopModeListLen int32
}

const S1BattleProto2022IoctrSetRmMotorsMoveValSize = 82

func (s *S1BattleProto2022IoctrSetRmMotorsMoveVal) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetRmMotorsMoveValSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetRmMotorsMoveVal) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
