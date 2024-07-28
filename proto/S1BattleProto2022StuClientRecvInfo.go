package proto

type S1BattleProto2022StuClientRecvInfo struct {
	TargetRobotId uint16
	TargetPosX    float32
	TargetPosY    float32
}

const S1BattleProto2022StuClientRecvInfoSize = 10

func (s *S1BattleProto2022StuClientRecvInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuClientRecvInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuClientRecvInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
