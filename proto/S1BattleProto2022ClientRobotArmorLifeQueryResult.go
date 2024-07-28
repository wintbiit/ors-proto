package proto

type S1BattleProto2022ClientRobotArmorLifeQueryResult struct {
	RobotId            byte
	RobotTotalArmorNum int32
	InfosLen           byte
	LifeInfos          []S1BattleProto2022ClientArmorLifeInfo
	IsWindMill         byte
	WindMillTeamId     byte
}

const S1BattleProto2022ClientRobotArmorLifeQueryResultSize = 56

func (s *S1BattleProto2022ClientRobotArmorLifeQueryResult) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotArmorLifeQueryResultSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotArmorLifeQueryResult) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
