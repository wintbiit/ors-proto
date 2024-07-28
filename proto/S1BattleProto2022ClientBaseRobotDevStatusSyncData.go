package proto

type S1BattleProto2022ClientBaseRobotDevStatusSyncData struct {
	BaseRobotDevStatusList    []S1BattleProto2022ClientBaseRobotDevStatus
	BaseRobotDevStatusListLen byte
}

const S1BattleProto2022ClientBaseRobotDevStatusSyncDataSize = 89

func (s *S1BattleProto2022ClientBaseRobotDevStatusSyncData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBaseRobotDevStatusSyncDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientBaseRobotDevStatusSyncData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
