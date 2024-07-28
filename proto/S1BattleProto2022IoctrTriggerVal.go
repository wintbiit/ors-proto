package proto

type S1BattleProto2022IoctrTriggerVal struct {
	Dp3InputValAfter  byte
	SysTickMs         uint32
	ModuleId          byte
	ModuleType        byte
	Dp1InputValBefore byte
	Dp3InputValBefore byte
	Dp1InputValAfter  byte
}

const S1BattleProto2022IoctrTriggerValSize = 10

func (s *S1BattleProto2022IoctrTriggerVal) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrTriggerValSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrTriggerVal) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
