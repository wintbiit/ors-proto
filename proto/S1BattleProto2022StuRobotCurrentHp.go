package proto

type S1BattleProto2022StuRobotCurrentHp struct {
	BlueBase      uint16
	Cmd           uint16
	RedSapper     uint16
	RedInfantry1  uint16
	RedBase       uint16
	BlueHero      uint16
	BlueGuard     uint16
	RedInfantry2  uint16
	BlueSapper    uint16
	BlueInfantry3 uint16
	BlueOutpost   uint16
	RedHero       uint16
	RedInfantry3  uint16
	RedGuard      uint16
	RedOutpost    uint16
	BlueInfantry2 uint16
	BlueInfantry1 uint16
}

const S1BattleProto2022StuRobotCurrentHpSize = 34

func (s *S1BattleProto2022StuRobotCurrentHp) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuRobotCurrentHpSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuRobotCurrentHp) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
