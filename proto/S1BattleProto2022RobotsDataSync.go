package proto

type S1BattleProto2022RobotsDataSync struct {
	Utc        uint32
	RobotData  S1BattleProto2022RobotDataSync
	ModuleId   byte
	ModuleType byte
	Progress   byte
	TimeRemain uint16
}

const S1BattleProto2022RobotsDataSyncSize = 82

func (s *S1BattleProto2022RobotsDataSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotsDataSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotsDataSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
