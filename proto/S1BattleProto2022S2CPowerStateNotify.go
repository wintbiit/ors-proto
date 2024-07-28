package proto

type S1BattleProto2022S2CPowerStateNotify struct {
	ClientId          byte
	ChassisPowerState int32
	GimbalPowerState  int32
	ShooterPowerState int32
}

const S1BattleProto2022S2CPowerStateNotifySize = 13

func (s *S1BattleProto2022S2CPowerStateNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022S2CPowerStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022S2CPowerStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
