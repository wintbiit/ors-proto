package proto

type S1BattleProto2022RobotWeaponFire struct {
	RollAngle  float32
	TimeStamp  uint64
	ShooterId  byte
	BulletType byte
	Speed      float32
	Frequency  float32
	Yaw        float32
	Pitch      float32
}

const S1BattleProto2022RobotWeaponFireSize = 30

func (s *S1BattleProto2022RobotWeaponFire) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotWeaponFireSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotWeaponFire) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
