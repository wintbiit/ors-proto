package proto

type S1BattleProto2022ClientRobotSyncData struct {
	SmallShootNum1              int32
	CurSmallFreq2               float32
	ShooterDetect0              byte
	PowerBuffer                 float32
	Blood1                      byte
	X                           float32
	HasshooterDetect0           byte
	Compass                     float32
	CoustomData3                float32
	UserId                      uint16
	BigLeftBulletCnt            int32
	CurExp                      float32
	Blood2                      byte
	Strike5                     byte
	CapNum                      byte
	CurBigSpeed                 float32
	HasshooterDetect1           byte
	Uwb                         byte
	Strike4                     byte
	CurBigHeatEnergy            float32
	Strike3                     byte
	Cap                         byte
	ChassisPowerLevel           int32
	BloodNum                    byte
	RebornRfid                  byte
	CurPower                    float32
	SmallLeftBulletsNum1        int32
	CurSmallHeatEnergyCoolRate2 float32
	ByAttackTotal               int32
	StrikeNum                   byte
	CameraNum                   byte
	Voltage                     float32
	Rfid0                       byte
	ShooterDetect1              byte
	ShooterDetect2              byte
	Current                     float32
	CurSmallHeatEnergyCoolRate1 float32
	CurSmallHeatEnergy2         float32
	Z                           float32
	Rfid1                       byte
	HasshooterDetect2           byte
	UwbNum                      byte
	HpLevel                     int32
	YellowWarningCount          byte
	CurSmallSpeed1              float32
	BigBulletsShootCnt          int32
	RfidNum                     byte
	CurHp                       uint16
	MurderId                    byte
	CoustomData2                float32
	CurrentPerformancePoint     int32
	Rssi                        byte
	CurSmallSpeed2              float32
	Strike6                     byte
	RtsDmgData                  S1BattleProto2022ClientHpChangeDetailVal
	SmallShootNum2              int32
	Blood0                      byte
	ShooterDetectNum            byte
	Strike9                     byte
	ShootSpeedLevel             int32
	SmallLeftBulletsNum2        int32
	CurBigFreq                  float32
	Power                       byte
	Strike7                     byte
	CurSmallFreq1               float32
	Mask                        byte
	CurSmallHeatEnergy1         float32
	CurBigHeatEnergyCoolRate    float32
	Strike1                     byte
	Strike2                     byte
	CoustomData1                float32
	Strike0                     byte
	Strike8                     byte
	Camera                      byte
	Y                           float32
}

const S1BattleProto2022ClientRobotSyncDataSize = 233

func (s *S1BattleProto2022ClientRobotSyncData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotSyncDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotSyncData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
