package proto

type S1BattleProto2022ClientRobotFirstData struct {
	Userid uint16
}

const S1BattleProto2022ClientRobotFirstDataSize = 2

func (s *S1BattleProto2022ClientRobotFirstData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotFirstDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotFirstData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
