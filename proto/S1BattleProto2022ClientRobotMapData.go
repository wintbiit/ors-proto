package proto

type S1BattleProto2022ClientRobotMapData struct {
	Yaw       float32
	Joinstate byte
	IsAlive   byte
	X         float32
	Y         float32
}

const S1BattleProto2022ClientRobotMapDataSize = 14

func (s *S1BattleProto2022ClientRobotMapData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotMapDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotMapData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
