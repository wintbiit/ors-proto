package proto

type S1BattleProto2022StuLeftBullet struct {
	Cmd                      uint16
	SmallAvaliableBulletsNum uint16
	BigAvaliableBulletsNum   uint16
	LeftCoinsNum             uint16
}

const S1BattleProto2022StuLeftBulletSize = 8

func (s *S1BattleProto2022StuLeftBullet) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuLeftBulletSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuLeftBullet) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
