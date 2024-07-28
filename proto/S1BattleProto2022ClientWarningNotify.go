package proto

type S1BattleProto2022ClientWarningNotify struct {
	HpChangePercent byte
	MaskTimeSelf    byte
	MaskTimeOthers  byte
	Type            byte
	Team            byte
	RobotId         byte
	Leftcredit      byte
}

const S1BattleProto2022ClientWarningNotifySize = 7

func (s *S1BattleProto2022ClientWarningNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientWarningNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientWarningNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
