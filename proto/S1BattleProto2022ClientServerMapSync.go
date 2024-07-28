package proto

type S1BattleProto2022ClientServerMapSync struct {
	AnchorMask     byte
	RobotidListLen byte
	RobotidList    []S1BattleProto2022ClientRobotMapData
	YawOffset      float32
}

const S1BattleProto2022ClientServerMapSyncSize = 314

func (s *S1BattleProto2022ClientServerMapSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientServerMapSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientServerMapSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
