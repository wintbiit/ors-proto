package proto

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
