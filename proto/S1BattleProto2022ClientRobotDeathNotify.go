package proto

type S1BattleProto2022ClientRobotDeathNotify struct {
	RobotidDeath  byte
	RobotidKiller byte
	DeathReason   byte
	BFirstBlood   byte
	KillCount     int32
}

const S1BattleProto2022ClientRobotDeathNotifySize = 8

func (s *S1BattleProto2022ClientRobotDeathNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotDeathNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotDeathNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
