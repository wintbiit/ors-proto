package proto

type S1BattleProto2022ClientRobotDataInfo struct {
	DeadReason  int32
	BehitedVal  int32
	TotalAttack int32
	RobotId     int32
	RobotType   int32
	CurHp       int32
	MaxHp       int32
	IsConnected byte
}

const S1BattleProto2022ClientRobotDataInfoSize = 29

func (s *S1BattleProto2022ClientRobotDataInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotDataInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotDataInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
