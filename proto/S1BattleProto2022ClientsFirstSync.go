package proto

type S1BattleProto2022ClientsFirstSync struct {
	RobotStatus     []S1BattleProto2022ClientRobotStatusNotify
	CourtYawOffset  float32
	BattleFirstData S1BattleProto2022ClientBattleFirstData
	RobotsFirstData []S1BattleProto2022ClientRobotFirstData
	ClientNum       int32
	RobotsNum1      int32
	GameMode        byte
	RobotsNum       int32
	ClientsStatus   []byte
}

const S1BattleProto2022ClientsFirstSyncSize = 429

func (s *S1BattleProto2022ClientsFirstSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientsFirstSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientsFirstSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
