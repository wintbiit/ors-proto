package proto

type S1BattleProto2022RobotSystemStatus struct {
	Blood1         byte
	Armor4         byte
	Armor7         byte
	Camera         byte
	ExtA1          byte
	ExtA4          byte
	ExtA6          byte
	ShooterDetect1 byte
	Armor3         byte
	ExtA3          byte
	ExtB1          byte
	ExtC0          byte
	ExtC2          byte
	Rfid1          byte
	Blood0         byte
	Uwb            byte
	ExtA0          byte
	ExtA5          byte
	Rfid0          byte
	Armor0         byte
	Armor1         byte
	Armor8         byte
	Armor9         byte
	Cap            byte
	ShooterDetect0 byte
	ShooterDetect2 byte
	Armor2         byte
	Armor6         byte
	ExtB0          byte
	ExtC1          byte
	ExtA2          byte
	ExtB2          byte
	ReservedList   []byte
	Power          byte
	Armor5         byte
	ReservedCount  int32
}

const S1BattleProto2022RobotSystemStatusSize = 41

func (s *S1BattleProto2022RobotSystemStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotSystemStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotSystemStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
