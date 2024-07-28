package proto

type S1BattleProto2022ClientsRobotModuleErrorInfo struct {
	Strike1               byte
	Strike2               byte
	Strike6               byte
	Blood1                byte
	BloodNum              byte
	SamllShooterDetectNum byte
	Strike0               byte
	ShooterDetect2        byte
	Armor1                byte
	Blood0                byte
	BigshooterDetect      byte
	Strike8               byte
	Cap                   byte
	Blood1                byte
	Strike3               byte
	Camera                byte
	BigShooterDetectNum   byte
	Strike4               byte
	Strike7               byte
	ShooterDetect1        byte
	Rfid0                 byte
	Rfid1                 byte
	Rfid1                 byte
	SamllshooterDetect1   byte
	Camera                byte
	PowerNum              byte
	RfidNum               byte
	Armor3                byte
	Power                 byte
	Blood2                byte
	Armor0                byte
	Armor7                byte
	IshasExShooter        byte
	UwbNum                byte
	Strike5               byte
	ShooterDetect0        byte
	Armor8                byte
	Armor9                byte
	MainController        byte
	Armor4                byte
	RobotId               int32
	Uwb                   byte
	StrikeNum             byte
	CapNum                byte
	Uwb                   byte
	Armor2                byte
	Armor5                byte
	Armor6                byte
	Strike9               byte
	Power                 byte
	Blood0                byte
	Rfid0                 byte
	SamllshooterDetect0   byte
	CameraNum             byte
	Cap                   byte
}

const S1BattleProto2022ClientsRobotModuleErrorInfoSize = 58

func (s *S1BattleProto2022ClientsRobotModuleErrorInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientsRobotModuleErrorInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientsRobotModuleErrorInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
