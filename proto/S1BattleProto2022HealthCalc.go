package proto

type S1BattleProto2022HealthCalc struct {
	ImpactHurt       uint16
	HealthPointStart uint16
	HealthPointFull  uint16
	BulletHurt       uint16
	GolfHurt         uint16
}

const S1BattleProto2022HealthCalcSize = 10

func (s *S1BattleProto2022HealthCalc) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022HealthCalcSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022HealthCalc) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
