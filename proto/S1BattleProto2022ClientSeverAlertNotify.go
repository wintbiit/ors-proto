package proto

type S1BattleProto2022ClientSeverAlertNotify struct {
	YellowRateWarningCount          int32
	RedRateWarningCount             int32
	LastRedTimestamp                uint32
	LastRedGameLeftTime             int32
	LastYellowTimestamp             uint32
	LastYellowGameLeftTime          int32
	CurrAlertLevel                  byte
	GetQingflowData                 byte
	BalancedInfantryUltraLimitError byte
}

const S1BattleProto2022ClientSeverAlertNotifySize = 27

func (s *S1BattleProto2022ClientSeverAlertNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientSeverAlertNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientSeverAlertNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
