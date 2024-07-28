package proto

type S1BattleProto2022SpeedModePid struct {
	Kp        float32
	Ki        float32
	Kd        float32
	MaxOutput uint16
}

const S1BattleProto2022SpeedModePidSize = 14

func (s *S1BattleProto2022SpeedModePid) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SpeedModePidSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SpeedModePid) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
