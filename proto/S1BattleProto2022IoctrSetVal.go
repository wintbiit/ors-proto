package proto

type S1BattleProto2022IoctrSetVal struct {
	Dp2OutputVal     []uint16
	OutDp1Io0        byte
	OutDp1Io1        byte
	OutDp1Io2        byte
	OutDp1Io3        byte
	OutDp1Io5        byte
	ModuleType       byte
	Dp2OutputValsLen int32
	OutDp1Io4        byte
	OutDp1Io6        byte
	OutDp1Io7        byte
	ModuleId         byte
}

const S1BattleProto2022IoctrSetValSize = 30

func (s *S1BattleProto2022IoctrSetVal) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetValSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetVal) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
