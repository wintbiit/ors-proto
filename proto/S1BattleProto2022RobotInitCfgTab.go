package proto

type S1BattleProto2022RobotInitCfgTab struct {
	Ext17MmSpeed byte
	ResCount     int32
	Res          []byte
}

const S1BattleProto2022RobotInitCfgTabSize = 5

func (s *S1BattleProto2022RobotInitCfgTab) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotInitCfgTabSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotInitCfgTab) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
