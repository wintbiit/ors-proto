package proto

type S1BattleProto2022PositionModePid struct {
	NoResponse     byte
	IntergralLimit uint16
	MaxSpeed       uint16
	Kp             float32
	Ki             float32
	Kd             float32
}

const S1BattleProto2022PositionModePidSize = 17

func (s *S1BattleProto2022PositionModePid) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022PositionModePidSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022PositionModePid) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
