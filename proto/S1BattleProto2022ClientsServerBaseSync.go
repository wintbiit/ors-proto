package proto

type S1BattleProto2022ClientsServerBaseSync struct {
	RobotsSyncDatasLen int32
	RobotsBaseSyncData []S1BattleProto2022ClientRobotBaseDataSync
}

const S1BattleProto2022ClientsServerBaseSyncSize = 2138

func (s *S1BattleProto2022ClientsServerBaseSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientsServerBaseSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientsServerBaseSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
