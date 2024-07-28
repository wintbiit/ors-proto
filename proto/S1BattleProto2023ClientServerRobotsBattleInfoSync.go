package proto

type S1BattleProto2023ClientServerRobotsBattleInfoSync struct {
	RobotsNum                 byte
	RobotsBattleInfoSyncDatas []S1BattleProto2023ClientRobotBattleInfo
}

const S1BattleProto2023ClientServerRobotsBattleInfoSyncSize = 397

func (s *S1BattleProto2023ClientServerRobotsBattleInfoSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientServerRobotsBattleInfoSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientServerRobotsBattleInfoSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
