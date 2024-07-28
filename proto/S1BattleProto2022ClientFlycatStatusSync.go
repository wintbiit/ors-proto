package proto

type S1BattleProto2022ClientFlycatStatusSync struct {
	SysState                   byte
	WorkState                  byte
	Battery                    byte
	MotorOneOnlineError        byte
	MotorTwoOnlineError        byte
	MotorOneOverheatError      byte
	TeamId                     byte
	Online                     byte
	CtrlState                  byte
	MotorTwoOverheatError      byte
	LowBatteryWarning          byte
	LowBatteryWarningThreshold byte
	BatteryRenewThreshold      byte
}

const S1BattleProto2022ClientFlycatStatusSyncSize = 13

func (s *S1BattleProto2022ClientFlycatStatusSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientFlycatStatusSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientFlycatStatusSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
