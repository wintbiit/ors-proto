package proto

type S1BattleProto2022ClientWeaponFireNotify struct {
	Robotid    byte
	BulletType byte
	Speed      float32
	Angle      float32
}

const S1BattleProto2022ClientWeaponFireNotifySize = 10

func (s *S1BattleProto2022ClientWeaponFireNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientWeaponFireNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientWeaponFireNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
