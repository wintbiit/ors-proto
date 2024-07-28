package proto

type S1BattleProto2022IoctrCfgSet struct {
	IsPwmDp2Io2     byte
	IsTriggerDp1Io5 byte
	IsTriggerDp3Io7 byte
	IsDp2SpiMode    byte
	IsPwmDp2Io6     byte
	Dp2PwmFrq       uint16
	IsTriggerDp1Io4 byte
	IsTriggerDp1Io7 byte
	IsTriggerDp3Io2 byte
	IsTriggerDp3Io6 byte
	IsPwmDp2Io5     byte
	IsPwmDp2Io7     byte
	IsTriggerDp1Io0 byte
	IsTriggerDp1Io1 byte
	IsTriggerDp1Io2 byte
	IsTriggerDp3Io0 byte
	IsTriggerDp3Io3 byte
	IsTriggerDp3Io4 byte
	ModuleId        byte
	IsOutputDp1     byte
	IsTriggerDp1Io3 byte
	IsTriggerDp1Io6 byte
	IsTriggerDp3Io1 byte
	IsTriggerDp3Io5 byte
	ModuleType      byte
}

const S1BattleProto2022IoctrCfgSetSize = 26

func (s *S1BattleProto2022IoctrCfgSet) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrCfgSetSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrCfgSet) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
