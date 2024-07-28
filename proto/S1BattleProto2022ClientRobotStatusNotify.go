package proto

type S1BattleProto2022ClientRobotStatusNotify struct {
	RobotUserId         int32
	JoinState           byte
	WifiWeak            byte
	ModuleOnline        byte
	BatteryLow          byte
	IsCanReliveByClient byte
	ConnState           byte
	SurvivalState       byte
	DeathSubState       int32
	IdConflict          byte
}

const S1BattleProto2022ClientRobotStatusNotifySize = 16

func (s *S1BattleProto2022ClientRobotStatusNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotStatusNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotStatusNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
