package proto

type S1BattleProto2022ClientRaderInfoToClient struct {
	TargetPosY    float32
	TorwardAngle  float32
	TargetRobotId uint16
	TargetPosX    float32
}

const S1BattleProto2022ClientRaderInfoToClientSize = 14

func (s *S1BattleProto2022ClientRaderInfoToClient) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRaderInfoToClientSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRaderInfoToClient) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
