package proto

type S1BattleProto2022ClientsServerSync struct {
	PassTime            float32
	LeftTime            float32
	CenterPointCoolTime float32
	TeamInfosLen        int32
	TeamInfos           []S1BattleProto2022ClientTeamInfo
	RobotsSyncDatasLen  int32
	RobotsSyncDatas     []S1BattleProto2022ClientRobotSyncData
}

const S1BattleProto2022ClientsServerSyncSize = 5238

func (s *S1BattleProto2022ClientsServerSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientsServerSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientsServerSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
