package proto

type S1BattleProto2022StuMissileRemainingTime struct {
	Cmdid uint16
	Time  byte
}

const S1BattleProto2022StuMissileRemainingTimeSize = 3

func (s *S1BattleProto2022StuMissileRemainingTime) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuMissileRemainingTimeSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuMissileRemainingTime) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
