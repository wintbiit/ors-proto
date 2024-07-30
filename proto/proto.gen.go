package proto

type S1BattleProto2022AirplaneStatusNotify struct {
	Robotid       byte    `json:"robotid"`
	Energy        int32   `json:"energy"`
	IsFire        byte    `json:"isFire"`
	Curbullet     int16   `json:"curbullet"`
	Maxbullet     int16   `json:"maxbullet"`
	Curshoottime  float32 `json:"curshoottime"`
	Fixshoottime  float32 `json:"fixshoottime"`
	Cdleft        float32 `json:"cdleft"`
	Countleft     uint16  `json:"countleft"`
	Cdrefreshcost int32   `json:"cdrefreshcost"`
	Isincd        byte    `json:"isincd"`
}

type S1BattleProto2022ArmorParaConfigItem struct {
	TiggerPress    float32 `json:"tiggerPress"`
	BulletMaxPress float32 `json:"bulletMaxPress"`
	GolfMinPress   float32 `json:"golfMinPress"`
}

type S1BattleProto2022BaseLightingEffect struct {
	LightColor uint32 `json:"lightColor"`
}

type S1BattleProto2022BaseShellStatus struct {
	BaseStatus                byte `json:"baseStatus"`
	BaseMotoOneStatus         byte `json:"baseMotoOneStatus"`
	BaseMotoOneBlockStatus    byte `json:"baseMotoOneBlockStatus"`
	BaseMotoOneOnlineStatus   byte `json:"baseMotoOneOnlineStatus"`
	BaseMotoTwoStatus         byte `json:"baseMotoTwoStatus"`
	BaseMotoTwoBlockStatus    byte `json:"baseMotoTwoBlockStatus"`
	BaseMotoTwoOnlineStatus   byte `json:"baseMotoTwoOnlineStatus"`
	BaseMotoThreeStatus       byte `json:"baseMotoThreeStatus"`
	BaseMotoThreeBlockStatus  byte `json:"baseMotoThreeBlockStatus"`
	BaseMotoThreeOnlineStatus byte `json:"baseMotoThreeOnlineStatus"`
	Reserved0                 byte `json:"reserved0"`
	Reserved2                 byte `json:"reserved2"`
}

type S1BattleProto2022ControlCmdT struct {
	OptCode uint16 `json:"optCode"`
	Data    uint16 `json:"data"`
}

type S1BattleProto2022ClientAirplaneCtrlReq struct {
	Robotid     byte `json:"robotid"`
	ControlCode byte `json:"controlCode"`
}

type S1BattleProto2022ClientAllClientStatusSync struct {
	Listlength byte   `json:"listlength"`
	Status     []byte `json:"status"`
}

type S1BattleProto2022ClientArmorLifeInfo struct {
	ModuleId  int16 `json:"moduleId"`
	LifeState byte  `json:"lifeState"`
}

type S1BattleProto2022ClientBaseRobotDevStatus struct {
	ShellStatus                    byte   `json:"shellStatus"`
	Online                         byte   `json:"online"`
	DartTargetPosition             int16  `json:"dartTargetPosition"`
	DartTargetIsInplace            byte   `json:"dartTargetIsInplace"`
	BaseShellOneSensorStatus       byte   `json:"baseShellOneSensorStatus"`
	BaseShellTwoSensorStatus       byte   `json:"baseShellTwoSensorStatus"`
	BaseShellThreeSensorStatus     byte   `json:"baseShellThreeSensorStatus"`
	BaseDartboardSensorStatus      byte   `json:"baseDartboardSensorStatus"`
	BaseShellMotoOneBlockStatus    byte   `json:"baseShellMotoOneBlockStatus"`
	BaseShellMotoOneOnlineStatus   byte   `json:"baseShellMotoOneOnlineStatus"`
	BaseShellMotoTwoBlockStatus    byte   `json:"baseShellMotoTwoBlockStatus"`
	BaseShellMotoTwoOnlineStatus   byte   `json:"baseShellMotoTwoOnlineStatus"`
	BaseShellMotoThreeBlockStatus  byte   `json:"baseShellMotoThreeBlockStatus"`
	BaseShellMotoThreeOnlineStatus byte   `json:"baseShellMotoThreeOnlineStatus"`
	BaseDartboardMotoBlockStatus   byte   `json:"baseDartboardMotoBlockStatus"`
	BaseDartboardMotoOnlineStatus  byte   `json:"baseDartboardMotoOnlineStatus"`
	BaseShellSelfcheckStatus       byte   `json:"baseShellSelfcheckStatus"`
	BaseDartboardSelfcheckStatus   byte   `json:"baseDartboardSelfcheckStatus"`
	BaseDartboardMoveTimeout       byte   `json:"baseDartboardMoveTimeout"`
	BaseShellOpenTimeout           byte   `json:"baseShellOpenTimeout"`
	BaseShellCloseTimeout          byte   `json:"baseShellCloseTimeout"`
	DartBoardIronline              byte   `json:"dartBoardIronline"`
	DartBoardBrokenErr             byte   `json:"dartBoardBrokenErr"`
	IrledStatusLeft                uint32 `json:"irledStatusLeft"`
	IrledStatusRight               uint32 `json:"irledStatusRight"`
	BaseShellMotoOneOverHeat       byte   `json:"baseShellMotoOneOverHeat"`
	BaseShellMotoTwoOverHeat       byte   `json:"baseShellMotoTwoOverHeat"`
	BaseShellMotoThreeOverHeat     byte   `json:"baseShellMotoThreeOverHeat"`
	BaseDartBoardMotoOverHeat      byte   `json:"baseDartBoardMotoOverHeat"`
	BaseShellMotoOneOverCurrent    byte   `json:"baseShellMotoOneOverCurrent"`
	BaseShellMotoTwoOverCurrent    byte   `json:"baseShellMotoTwoOverCurrent"`
	BaseShellMotoThreeOverCurrent  byte   `json:"baseShellMotoThreeOverCurrent"`
	BaseDartBoardMotoOverCurrent   byte   `json:"baseDartBoardMotoOverCurrent"`
	BaseShellMotoOneOverSpeed      byte   `json:"baseShellMotoOneOverSpeed"`
	BaseShellMotoTwoOverSpeed      byte   `json:"baseShellMotoTwoOverSpeed"`
	BaseShellMotoThreeOverSpeed    byte   `json:"baseShellMotoThreeOverSpeed"`
	BaseDartBoardMotoOverSpeed     byte   `json:"baseDartBoardMotoOverSpeed"`
}

type S1BattleProto2022ClientBaseRobotDevStatusSyncData struct {
	BaseRobotDevStatusListLen byte                                        `json:"baseRobotDevStatusListLen"`
	BaseRobotDevStatusList    []S1BattleProto2022ClientBaseRobotDevStatus `json:"baseRobotDevStatusList"`
}

type S1BattleProto2022ClientBattleFirstData struct {
	Progress byte `json:"progress"`
	IsPaused byte `json:"isPaused"`
}

type S1BattleProto2022ClientBuffItem struct {
	BuffId byte    `json:"buffId"`
	CdTime float32 `json:"cdTime"`
}

type S1BattleProto2022ClientBuffSlot struct {
	BuffItemsLen int32                             `json:"buffItemsLen"`
	BuffItems    []S1BattleProto2022ClientBuffItem `json:"buffItems"`
}

type S1BattleProto2022ClientBuyBulletAck struct {
	Result int32 `json:"result"`
}

type S1BattleProto2022ClientBuyBulletReq struct {
	Type  byte `json:"type"`
	Count byte `json:"count"`
}

type S1BattleProto2022ClientCheckInTimeStampNotify struct {
	RobotidListLen byte     `json:"robotidListLen"`
	RobotidList    []uint32 `json:"robotidList"`
}

type S1BattleProto2022ClientCurrentProgress struct {
	CurProgress int32 `json:"curProgress"`
	IsPaused    byte  `json:"isPaused"`
}

type S1BattleProto2022ClientCustomControlData struct {
	RobotId     int32  `json:"robotId"`
	DataLen     int32  `json:"dataLen"`
	DataListLen byte   `json:"dataListLen"`
	DataList    []byte `json:"dataList"`
}

type S1BattleProto2022ClientEconomyNotify struct {
	Clientid byte  `json:"clientid"`
	Type     int32 `json:"type"`
	Money    int32 `json:"money"`
}

type S1BattleProto2022ClientExceptionCapDataNotify struct {
	RobotidListLen byte   `json:"robotidListLen"`
	RobotidList    []byte `json:"robotidList"`
}

type S1BattleProto2022ClientExceptionRecoverCardNotify struct {
	Teamid    byte `json:"teamid"`
	Robotid   byte `json:"robotid"`
	Exception byte `json:"exception"`
}

type S1BattleProto2022ClientExchangeProStateNotify struct {
	RobotId          byte    `json:"robotId"`
	IsOnline         byte    `json:"isOnline"`
	ExchangeLevel    byte    `json:"exchangeLevel"`
	IrState          byte    `json:"irState"`
	Ir1              byte    `json:"ir1"`
	Ir2              byte    `json:"ir2"`
	X                float32 `json:"x"`
	Y                float32 `json:"y"`
	Z                float32 `json:"z"`
	Pitch            float32 `json:"pitch"`
	Roll             float32 `json:"roll"`
	Yaw              float32 `json:"yaw"`
	ErrorCode        byte    `json:"errorCode"`
	GoldCount        byte    `json:"goldCount"`
	SilverCount      byte    `json:"silverCount"`
	State            byte    `json:"state"`
	MineralRfidState byte    `json:"mineralRfidState"`
	CoinsExchange    int32   `json:"coinsExchange"`
	CoinsAddRate     float32 `json:"coinsAddRate"`
	EngineerLevel    byte    `json:"engineerLevel"`
}

type S1BattleProto2022ClientFlycatStatusSync struct {
	TeamId                     byte `json:"teamId"`
	Online                     byte `json:"online"`
	SysState                   byte `json:"sysState"`
	CtrlState                  byte `json:"ctrlState"`
	WorkState                  byte `json:"workState"`
	Battery                    byte `json:"battery"`
	MotorOneOnlineError        byte `json:"motorOneOnlineError"`
	MotorTwoOnlineError        byte `json:"motorTwoOnlineError"`
	MotorOneOverheatError      byte `json:"motorOneOverheatError"`
	MotorTwoOverheatError      byte `json:"motorTwoOverheatError"`
	LowBatteryWarning          byte `json:"lowBatteryWarning"`
	LowBatteryWarningThreshold byte `json:"lowBatteryWarningThreshold"`
	BatteryRenewThreshold      byte `json:"batteryRenewThreshold"`
}

type S1BattleProto2022ClientGmcommand struct {
	Len byte   `json:"len"`
	Cmd string `json:"cmd"`
}

type S1BattleProto2022ClientHitNotify struct {
	RobotId     byte  `json:"robotId"`
	OnHitType   byte  `json:"onHitType"`
	ArmorNumber byte  `json:"armorNumber"`
	HpReduce    int32 `json:"hpReduce"`
	HpCurr      int32 `json:"hpCurr"`
	HpMax       int32 `json:"hpMax"`
}

type S1BattleProto2022ClientHoldCenterPointCompleteNotify struct {
	Teamid    int32   `json:"teamid"`
	RewardExp float32 `json:"rewardExp"`
}

type S1BattleProto2022ClientHpChangeDetailVal struct {
	ByBullet        int32 `json:"byBullet"`
	ByGolf          int32 `json:"byGolf"`
	ByMissile       int32 `json:"byMissile"`
	ByOverSpeed     int32 `json:"byOverSpeed"`
	ByoverFreq      int32 `json:"byoverFreq"`
	ByOverPower     int32 `json:"byOverPower"`
	ByOverHeat      int32 `json:"byOverHeat"`
	ByModuleOffline int32 `json:"byModuleOffline"`
	ByPunish        int32 `json:"byPunish"`
	ByKill          int32 `json:"byKill"`
	ByCrash         int32 `json:"byCrash"`
	ByWifiOffline   int32 `json:"byWifiOffline"`
	All             int32 `json:"all"`
}

type S1BattleProto2022ClientKickOffRobotNotify struct {
	RobotId int32 `json:"robotId"`
	Reason  byte  `json:"reason"`
}

type S1BattleProto2022ClientMineralExchangeNotify struct {
	TeamId      byte  `json:"teamId"`
	MineralType byte  `json:"mineralType"`
	Value       int32 `json:"value"`
}

type S1BattleProto2022ClientMineralInfoSync struct {
	IsRedConnect             byte   `json:"isRedConnect"`
	IsBlueConnect            byte   `json:"isBlueConnect"`
	RedGoldCount             byte   `json:"redGoldCount"`
	RedSilverCount           byte   `json:"redSilverCount"`
	BlueGoldCount            byte   `json:"blueGoldCount"`
	BlueSilverCount          byte   `json:"blueSilverCount"`
	RedInfraredStateListLen  byte   `json:"redInfraredStateListLen"`
	RedInfraredState         []byte `json:"redInfraredState"`
	BlueInfraredStateListLen byte   `json:"blueInfraredStateListLen"`
	BlueInfraredState        []byte `json:"blueInfraredState"`
}

type S1BattleProto2022ClientOutpostDeviceStatusSync struct {
	RobotId               byte    `json:"robotId"`
	Online                byte    `json:"online"`
	StatusCode            byte    `json:"statusCode"`
	CurTemperature        int32   `json:"curTemperature"`
	TemperatureWarning    byte    `json:"temperatureWarning"`
	SpinSpeedRatio        float32 `json:"spinSpeedRatio"`
	MotorOnline           byte    `json:"motorOnline"`
	MotorCurrentWarning   byte    `json:"motorCurrentWarning"`
	MotorSpeedoverWarning byte    `json:"motorSpeedoverWarning"`
	IsMagnetOn            byte    `json:"isMagnetOn"`
	ActionTimeout         byte    `json:"actionTimeout"`
	DartBoardIrOnline     byte    `json:"dartBoardIrOnline"`
	DartBoardBrokenErr    byte    `json:"dartBoardBrokenErr"`
	IrledStatusLeft       uint32  `json:"irledStatusLeft"`
	IrledStatusRight      uint32  `json:"irledStatusRight"`
}

type S1BattleProto2022ClientRaderInfoToClient struct {
	TargetRobotId uint16  `json:"targetRobotId"`
	TargetPosX    float32 `json:"targetPosX"`
	TargetPosY    float32 `json:"targetPosY"`
	TorwardAngle  float32 `json:"torwardAngle"`
}

type S1BattleProto2022ClientRecoverCardStatus struct {
	Redcard  byte `json:"redcard"`
	Bluecard byte `json:"bluecard"`
}

type S1BattleProto2022ClientResultDataInfo struct {
	WebGameId                    int32                                  `json:"webGameId"`
	GameId                       int64                                  `json:"gameId"`
	GameOverReasonId             byte                                   `json:"gameOverReasonId"`
	GameOverReasonListLen        int32                                  `json:"gameOverReasonListLen"`
	GameOverReason               []byte                                 `json:"gameOverReason"`
	StartTime                    int64                                  `json:"startTime"`
	DurationTime                 int32                                  `json:"durationTime"`
	GameOrder                    byte                                   `json:"gameOrder"`
	TotalRound                   byte                                   `json:"totalRound"`
	RecordsListLen               int32                                  `json:"recordsListLen"`
	Records                      []int32                                `json:"records"`
	Winner                       byte                                   `json:"winner"`
	ScoreRed                     int32                                  `json:"scoreRed"`
	ScoreBlue                    int32                                  `json:"scoreBlue"`
	RedCurrentHp                 int32                                  `json:"redCurrentHp"`
	BlueCurrentHp                int32                                  `json:"blueCurrentHp"`
	RedTotalHp                   int32                                  `json:"redTotalHp"`
	BlueTotalHp                  int32                                  `json:"blueTotalHp"`
	RedWarning1                  int32                                  `json:"redWarning1"`
	RedWarning2                  int32                                  `json:"redWarning2"`
	RedWarning3                  int32                                  `json:"redWarning3"`
	BlueWarning1                 int32                                  `json:"blueWarning1"`
	BlueWarning2                 int32                                  `json:"blueWarning2"`
	BlueWarning3                 int32                                  `json:"blueWarning3"`
	RedHits                      int32                                  `json:"redHits"`
	BlueHits                     int32                                  `json:"blueHits"`
	RedRuneCount                 int32                                  `json:"redRuneCount"`
	BlueRuneCount                int32                                  `json:"blueRuneCount"`
	RedKillCount                 int32                                  `json:"redKillCount"`
	BlueKillCount                int32                                  `json:"blueKillCount"`
	RedDartHitCount              int32                                  `json:"redDartHitCount"`
	BlueDartHitCount             int32                                  `json:"blueDartHitCount"`
	RedShootSmallCount           int32                                  `json:"redShootSmallCount"`
	RedShootBigCount             int32                                  `json:"redShootBigCount"`
	BlueShootSmallCount          int32                                  `json:"blueShootSmallCount"`
	BlueShootBigCount            int32                                  `json:"blueShootBigCount"`
	RedGuardLeftLives            byte                                   `json:"redGuardLeftLives"`
	BlueGuardLeftLives           byte                                   `json:"blueGuardLeftLives"`
	GuardFixedLives              byte                                   `json:"guardFixedLives"`
	RobotsNum                    int32                                  `json:"robotsNum"`
	RobotsInfos                  []S1BattleProto2022ClientRobotDataInfo `json:"robotsInfos"`
	RedSnipeSuccCount            byte                                   `json:"redSnipeSuccCount"`
	BlueSnipeSuccCount           byte                                   `json:"blueSnipeSuccCount"`
	RedUavCallCount              byte                                   `json:"redUavCallCount"`
	BlueUavCallCount             byte                                   `json:"blueUavCallCount"`
	RedRadarBuffDoubleUsedCount  byte                                   `json:"redRadarBuffDoubleUsedCount"`
	BlueRadarBuffDoubleUsedCount byte                                   `json:"blueRadarBuffDoubleUsedCount"`
}

type S1BattleProto2022ClientRobotArmorLifeQueryResult struct {
	IsWindMill         byte                                   `json:"isWindMill"`
	WindMillTeamId     byte                                   `json:"windMillTeamId"`
	RobotId            byte                                   `json:"robotId"`
	RobotTotalArmorNum int32                                  `json:"robotTotalArmorNum"`
	InfosLen           byte                                   `json:"infosLen"`
	LifeInfos          []S1BattleProto2022ClientArmorLifeInfo `json:"lifeInfos"`
}

type S1BattleProto2022ClientRobotBaseDataSync struct {
	UserId                      uint16  `json:"userId"`
	Id                          byte    `json:"id"`
	Type                        byte    `json:"type"`
	Level                       byte    `json:"level"`
	CpuId                       uint32  `json:"cpuId"`
	FixHp                       int32   `json:"fixHp"`
	FixPower                    float32 `json:"fixPower"`
	FixPowerBufferRecoverable   float32 `json:"fixPowerBufferRecoverable"`
	FixPowerBufferUnrecoverable float32 `json:"fixPowerBufferUnrecoverable"`
	FixSmallSpeed               float32 `json:"fixSmallSpeed"`
	FixSmallFreq                float32 `json:"fixSmallFreq"`
	FixSmallHeatEnergy          float32 `json:"fixSmallHeatEnergy"`
	FixSmallHeatEnergyCoolRate  float32 `json:"fixSmallHeatEnergyCoolRate"`
	FixSmallSpeed2              float32 `json:"fixSmallSpeed2"`
	FixSmallFreq2               float32 `json:"fixSmallFreq2"`
	FixSmallHeatEnergy2         float32 `json:"fixSmallHeatEnergy2"`
	FixSmallHeatEnergyCoolRate2 float32 `json:"fixSmallHeatEnergyCoolRate2"`
	FixBigSpeed                 float32 `json:"fixBigSpeed"`
	FixBigFreq                  float32 `json:"fixBigFreq"`
	FixBigHeatEnergy            float32 `json:"fixBigHeatEnergy"`
	FixBigHeatEnergyCoolRate    float32 `json:"fixBigHeatEnergyCoolRate"`
	ExpOffer                    float32 `json:"expOffer"`
	CurExpAddRate               float32 `json:"curExpAddRate"`
	CurRebornCd                 int32   `json:"curRebornCd"`
	FixExpUpgradeNeed           float32 `json:"fixExpUpgradeNeed"`
	FixExpAddRate               float32 `json:"fixExpAddRate"`
	FixRebornCd                 int32   `json:"fixRebornCd"`
}

type S1BattleProto2022ClientRobotDataInfo struct {
	RobotId     int32 `json:"robotId"`
	RobotType   int32 `json:"robotType"`
	CurHp       int32 `json:"curHp"`
	MaxHp       int32 `json:"maxHp"`
	IsConnected byte  `json:"isConnected"`
	DeadReason  int32 `json:"deadReason"`
	BehitedVal  int32 `json:"behitedVal"`
	TotalAttack int32 `json:"totalAttack"`
}

type S1BattleProto2022ClientRobotDeathNotify struct {
	RobotidDeath  byte  `json:"robotidDeath"`
	RobotidKiller byte  `json:"robotidKiller"`
	DeathReason   byte  `json:"deathReason"`
	BFirstBlood   byte  `json:"bfirstBlood"`
	KillCount     int32 `json:"killCount"`
}

type S1BattleProto2022ClientRobotFirstData struct {
	Userid uint16 `json:"userid"`
}

type S1BattleProto2022ClientRobotMapData struct {
	Joinstate byte    `json:"joinstate"`
	IsAlive   byte    `json:"isAlive"`
	X         float32 `json:"x"`
	Y         float32 `json:"y"`
	Yaw       float32 `json:"yaw"`
}

type S1BattleProto2022ClientRobotStatusNotify struct {
	RobotUserId         int32 `json:"robotUserId"`
	ConnState           byte  `json:"connState"`
	JoinState           byte  `json:"joinState"`
	SurvivalState       byte  `json:"survivalState"`
	DeathSubState       int32 `json:"deathSubState"`
	WifiWeak            byte  `json:"wifiWeak"`
	ModuleOnline        byte  `json:"moduleOnline"`
	BatteryLow          byte  `json:"batteryLow"`
	IdConflict          byte  `json:"idConflict"`
	IsCanReliveByClient byte  `json:"isCanReliveByClient"`
}

type S1BattleProto2022ClientRobotSyncData struct {
	UserId                      uint16                                   `json:"userId"`
	CurHp                       uint16                                   `json:"curHp"`
	Voltage                     float32                                  `json:"voltage"`
	Current                     float32                                  `json:"current"`
	Rssi                        byte                                     `json:"rssi"`
	CurPower                    float32                                  `json:"curPower"`
	YellowWarningCount          byte                                     `json:"yellowWarningCount"`
	PowerBuffer                 float32                                  `json:"powerBuffer"`
	CurSmallSpeed1              float32                                  `json:"curSmallSpeed1"`
	CurSmallFreq1               float32                                  `json:"curSmallFreq1"`
	CurSmallHeatEnergy1         float32                                  `json:"curSmallHeatEnergy1"`
	CurSmallHeatEnergyCoolRate1 float32                                  `json:"curSmallHeatEnergyCoolRate1"`
	SmallShootNum1              int32                                    `json:"smallShootNum1"`
	SmallLeftBulletsNum1        int32                                    `json:"smallLeftBulletsNum1"`
	CurSmallSpeed2              float32                                  `json:"curSmallSpeed2"`
	CurSmallFreq2               float32                                  `json:"curSmallFreq2"`
	CurSmallHeatEnergy2         float32                                  `json:"curSmallHeatEnergy2"`
	CurSmallHeatEnergyCoolRate2 float32                                  `json:"curSmallHeatEnergyCoolRate2"`
	SmallShootNum2              int32                                    `json:"smallShootNum2"`
	SmallLeftBulletsNum2        int32                                    `json:"smallLeftBulletsNum2"`
	CurBigSpeed                 float32                                  `json:"curBigSpeed"`
	CurBigFreq                  float32                                  `json:"curBigFreq"`
	CurBigHeatEnergy            float32                                  `json:"curBigHeatEnergy"`
	CurBigHeatEnergyCoolRate    float32                                  `json:"curBigHeatEnergyCoolRate"`
	BigBulletsShootCnt          int32                                    `json:"bigBulletsShootCnt"`
	BigLeftBulletCnt            int32                                    `json:"bigLeftBulletCnt"`
	CurExp                      float32                                  `json:"curExp"`
	ByAttackTotal               int32                                    `json:"byAttackTotal"`
	MurderId                    byte                                     `json:"murderId"`
	Power                       byte                                     `json:"power"`
	RfidNum                     byte                                     `json:"rfidNum"`
	Rfid0                       byte                                     `json:"rfid0"`
	Rfid1                       byte                                     `json:"rfid1"`
	BloodNum                    byte                                     `json:"bloodNum"`
	Blood0                      byte                                     `json:"blood0"`
	Blood1                      byte                                     `json:"blood1"`
	Blood2                      byte                                     `json:"blood2"`
	ShooterDetectNum            byte                                     `json:"shooterDetectNum"`
	ShooterDetect0              byte                                     `json:"shooterDetect0"`
	ShooterDetect1              byte                                     `json:"shooterDetect1"`
	ShooterDetect2              byte                                     `json:"shooterDetect2"`
	HasshooterDetect0           byte                                     `json:"hasshooterDetect0"`
	HasshooterDetect1           byte                                     `json:"hasshooterDetect1"`
	HasshooterDetect2           byte                                     `json:"hasshooterDetect2"`
	UwbNum                      byte                                     `json:"uwbNum"`
	Uwb                         byte                                     `json:"uwb"`
	StrikeNum                   byte                                     `json:"strikeNum"`
	Strike0                     byte                                     `json:"strike0"`
	Strike1                     byte                                     `json:"strike1"`
	Strike2                     byte                                     `json:"strike2"`
	Strike3                     byte                                     `json:"strike3"`
	Strike4                     byte                                     `json:"strike4"`
	Strike5                     byte                                     `json:"strike5"`
	Strike6                     byte                                     `json:"strike6"`
	Strike7                     byte                                     `json:"strike7"`
	Strike8                     byte                                     `json:"strike8"`
	Strike9                     byte                                     `json:"strike9"`
	CameraNum                   byte                                     `json:"cameraNum"`
	Camera                      byte                                     `json:"camera"`
	CapNum                      byte                                     `json:"capNum"`
	Cap                         byte                                     `json:"cap"`
	RebornRfid                  byte                                     `json:"rebornRfid"`
	X                           float32                                  `json:"x"`
	Y                           float32                                  `json:"y"`
	Z                           float32                                  `json:"z"`
	Compass                     float32                                  `json:"compass"`
	CoustomData1                float32                                  `json:"coustomData1"`
	CoustomData2                float32                                  `json:"coustomData2"`
	CoustomData3                float32                                  `json:"coustomData3"`
	Mask                        byte                                     `json:"mask"`
	RtsDmgData                  S1BattleProto2022ClientHpChangeDetailVal `json:"rtsDmgData"`
	CurrentPerformancePoint     int32                                    `json:"currentPerformancePoint"`
	HpLevel                     int32                                    `json:"hpLevel"`
	ChassisPowerLevel           int32                                    `json:"chassisPowerLevel"`
	ShootSpeedLevel             int32                                    `json:"shootSpeedLevel"`
}

type S1BattleProto2022ClientShielddata struct {
	RedHasShield  byte  `json:"redHasShield"`
	BlueHasShield byte  `json:"blueHasShield"`
	RedShield     int32 `json:"redShield"`
	BlueShield    int32 `json:"blueShield"`
	ShieldMax     int32 `json:"shieldMax"`
}

type S1BattleProto2022ClientSafetySync struct {
	Id            byte `json:"id"`
	Lift1Connect  byte `json:"lift1Connect"`
	Lift2Connect  byte `json:"lift2Connect"`
	FlycatConnect byte `json:"flycatConnect"`
	Lift1State    byte `json:"lift1State"`
	Lift2State    byte `json:"lift2State"`
	FlycatState   byte `json:"flycatState"`
	Lift1Error    byte `json:"lift1Error"`
	Lift2Error    byte `json:"lift2Error"`
	CatError      byte `json:"catError"`
	FlycatBattery byte `json:"flycatBattery"`
}

type S1BattleProto2022ClientServerMapSync struct {
	YawOffset      float32                               `json:"yawOffset"`
	AnchorMask     byte                                  `json:"anchorMask"`
	RobotidListLen byte                                  `json:"robotidListLen"`
	RobotidList    []S1BattleProto2022ClientRobotMapData `json:"robotidList"`
}

type S1BattleProto2022ClientServerUimessage struct {
	Msg string `json:"msg"`
}

type S1BattleProto2022ClientSeverAlertNotify struct {
	YellowRateWarningCount          int32  `json:"yellowRateWarningCount"`
	RedRateWarningCount             int32  `json:"redRateWarningCount"`
	LastYellowTimestamp             uint32 `json:"lastYellowTimestamp"`
	LastRedTimestamp                uint32 `json:"lastRedTimestamp"`
	LastYellowGameLeftTime          int32  `json:"lastYellowGameLeftTime"`
	LastRedGameLeftTime             int32  `json:"lastRedGameLeftTime"`
	CurrAlertLevel                  byte   `json:"currAlertLevel"`
	GetQingflowData                 byte   `json:"getQingflowData"`
	BalancedInfantryUltraLimitError byte   `json:"balancedInfantryUltraLimitError"`
}

type S1BattleProto2022ClientShowMessage struct {
	Teamid    byte   `json:"teamid"`
	MsgType   byte   `json:"msgType"`
	MsgEnum   uint32 `json:"msgEnum"`
	MsgParams string `json:"msgParams"`
}

type S1BattleProto2022ClientSiloEnvCtrReq struct {
	Teamid      byte `json:"teamid"`
	ControlCode byte `json:"controlCode"`
	Target      byte `json:"target"`
}

type S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify struct {
	TeamId      byte `json:"teamId"`
	DoorState   byte `json:"doorState"`
	DoorOpenCnt byte `json:"doorOpenCnt"`
}

type S1BattleProto2022ClientStatusSync struct {
	ClientId int32 `json:"clientId"`
	Status   int32 `json:"status"`
}

type S1BattleProto2022ClientSupplySync struct {
	SupplyId            int32   `json:"supplyId"`
	Connect             int32   `json:"connect"`
	State               int32   `json:"state"`
	UsableBullets       int32   `json:"usableBullets"`
	UsedBullets         int32   `json:"usedBullets"`
	ReadyBox            int32   `json:"readyBox"`
	FreedBox            int32   `json:"freedBox"`
	Timespan            float32 `json:"timespan"`
	NextaddBullets      int32   `json:"nextaddBullets"`
	MakerRobotid        int32   `json:"makerRobotid"`
	MakeBullets         int32   `json:"makeBullets"`
	ErrorCodeArrListLen int32   `json:"errorCodeArrListLen"`
	ErrorCodeArrList    []int32 `json:"errorCodeArrList"`
	RfidRobotIdListLen  int32   `json:"rfidRobotIdListLen"`
	RfidRobotIdList     []int32 `json:"rfidRobotIdList"`
}

type S1BattleProto2022ClientTalentDataAck struct {
	Result int32 `json:"result"`
}

type S1BattleProto2022ClientTalentDataNotify struct {
	Robotid             byte      `json:"robotid"`
	Data                []性能体系_数据 `json:"data"`
	IsBalance           byte      `json:"isBalance"`
	IsSemiAutomaticCtrl byte      `json:"isSemiAutomaticCtrl"`
	IsTempManualCtrl    byte      `json:"isTempManualCtrl"`
}

type S1BattleProto2022ClientTalentDataReq struct {
	Robotid byte   `json:"robotid"`
	目标类型    string `json:""`
	性能类型    string `json:""`
	是否设置副武器 byte   `json:""`
}

type S1BattleProto2022ClientTeamInfo struct {
	TotalHp                 uint32 `json:"totalHp"`
	CurrentHp               uint32 `json:"currentHp"`
	Warning1Count           uint16 `json:"warning1Count"`
	Warning2Count           uint16 `json:"warning2Count"`
	Warning3Count           uint16 `json:"warning3Count"`
	TotalHurt               uint16 `json:"totalHurt"`
	BigBulletAvailableNum   int16  `json:"bigBulletAvailableNum"`
	BigBulletQuotaNum       int16  `json:"bigBulletQuotaNum"`
	BigBulletCanBuyNum      int16  `json:"bigBulletCanBuyNum"`
	SmallBulletAvailableNum int16  `json:"smallBulletAvailableNum"`
	SmallBulletQuotaNum     int16  `json:"smallBulletQuotaNum"`
	SmallBulletCanBuyNum    int16  `json:"smallBulletCanBuyNum"`
	AirSupportNum           uint16 `json:"airSupportNum"`
	BloodPack               uint32 `json:"bloodPack"`
	CenterPointEnergy       uint32 `json:"centerPointEnergy"`
	TeamAvaliableCoinsNum   int32  `json:"teamAvaliableCoinsNum"`
	TotalCoins              int32  `json:"totalCoins"`
}

type S1BattleProto2022ClientVtmSpeedSync struct {
	Clientid             int32 `json:"clientid"`
	CurrentFreq          int32 `json:"currentFreq"`
	TxConnect            int32 `json:"txConnect"`
	CurrentSpeedRate     int32 `json:"currentSpeedRate"`
	CurrentTxTemperature int32 `json:"currentTxTemperature"`
}

type S1BattleProto2022ClientWarningNotify struct {
	Type            byte `json:"type"`
	Team            byte `json:"team"`
	RobotId         byte `json:"robotId"`
	Leftcredit      byte `json:"leftcredit"`
	HpChangePercent byte `json:"hpChangePercent"`
	MaskTimeSelf    byte `json:"maskTimeSelf"`
	MaskTimeOthers  byte `json:"maskTimeOthers"`
}

type S1BattleProto2022ClientWeaponFireNotify struct {
	Robotid    byte    `json:"robotid"`
	BulletType byte    `json:"bulletType"`
	Speed      float32 `json:"speed"`
	Angle      float32 `json:"angle"`
}

type S1BattleProto2022ClientsFirstSync struct {
	GameMode        byte                                       `json:"gameMode"`
	BattleFirstData S1BattleProto2022ClientBattleFirstData     `json:"battleFirstData"`
	RobotsNum       int32                                      `json:"robotsNum"`
	RobotsFirstData []S1BattleProto2022ClientRobotFirstData    `json:"robotsFirstData"`
	ClientNum       int32                                      `json:"clientNum"`
	ClientsStatus   []byte                                     `json:"clientsStatus"`
	RobotsNum1      int32                                      `json:"robotsNum1"`
	RobotStatus     []S1BattleProto2022ClientRobotStatusNotify `json:"robotStatus"`
	CourtYawOffset  float32                                    `json:"courtYawOffset"`
}

type S1BattleProto2022ClientsRobotModuleErrorInfo struct {
	RobotId               int32 `json:"robotId"`
	PowerNum              byte  `json:"powerNum"`
	Power                 byte  `json:"power"`
	RfidNum               byte  `json:"rfidNum"`
	Rfid0                 byte  `json:"rfid0"`
	Rfid1                 byte  `json:"rfid1"`
	BloodNum              byte  `json:"bloodNum"`
	Blood0                byte  `json:"blood0"`
	Blood1                byte  `json:"blood1"`
	Blood2                byte  `json:"blood2"`
	IshasExShooter        byte  `json:"ishasExShooter"`
	SamllShooterDetectNum byte  `json:"samllShooterDetectNum"`
	SamllshooterDetect0   byte  `json:"samllshooterDetect0"`
	SamllshooterDetect1   byte  `json:"samllshooterDetect1"`
	BigShooterDetectNum   byte  `json:"bigShooterDetectNum"`
	BigshooterDetect      byte  `json:"bigshooterDetect"`
	UwbNum                byte  `json:"uwbNum"`
	Uwb                   byte  `json:"uwb"`
	StrikeNum             byte  `json:"strikeNum"`
	Strike0               byte  `json:"strike0"`
	Strike1               byte  `json:"strike1"`
	Strike2               byte  `json:"strike2"`
	Strike3               byte  `json:"strike3"`
	Strike4               byte  `json:"strike4"`
	Strike5               byte  `json:"strike5"`
	Strike6               byte  `json:"strike6"`
	Strike7               byte  `json:"strike7"`
	Strike8               byte  `json:"strike8"`
	Strike9               byte  `json:"strike9"`
	CameraNum             byte  `json:"cameraNum"`
	Camera                byte  `json:"camera"`
	CapNum                byte  `json:"capNum"`
	Cap                   byte  `json:"cap"`
	Spower                byte  `json:"spower"`
	Srfid0                byte  `json:"srfid0"`
	Srfid1                byte  `json:"srfid1"`
	Scamera               byte  `json:"scamera"`
	Sblood0               byte  `json:"sblood0"`
	Sblood1               byte  `json:"sblood1"`
	SshooterDetect0       byte  `json:"sshooterDetect0"`
	SshooterDetect1       byte  `json:"sshooterDetect1"`
	SshooterDetect2       byte  `json:"sshooterDetect2"`
	Suwb                  byte  `json:"suwb"`
	Sarmor0               byte  `json:"sarmor0"`
	Sarmor1               byte  `json:"sarmor1"`
	Sarmor2               byte  `json:"sarmor2"`
	Sarmor3               byte  `json:"sarmor3"`
	Sarmor4               byte  `json:"sarmor4"`
	Sarmor5               byte  `json:"sarmor5"`
	Sarmor6               byte  `json:"sarmor6"`
	Sarmor7               byte  `json:"sarmor7"`
	Sarmor8               byte  `json:"sarmor8"`
	Sarmor9               byte  `json:"sarmor9"`
	Scap                  byte  `json:"scap"`
	SmainController       byte  `json:"smainController"`
}

type S1BattleProto2022ClientsServerBaseSync struct {
	RobotsSyncDatasLen int32                                      `json:"robotsSyncDatasLen"`
	RobotsBaseSyncData []S1BattleProto2022ClientRobotBaseDataSync `json:"robotsBaseSyncData"`
}

type S1BattleProto2022ClientsServerSync struct {
	PassTime            float32                                `json:"passTime"`
	LeftTime            float32                                `json:"leftTime"`
	CenterPointCoolTime float32                                `json:"centerPointCoolTime"`
	TeamInfosLen        int32                                  `json:"teamInfosLen"`
	TeamInfos           []S1BattleProto2022ClientTeamInfo      `json:"teamInfos"`
	RobotsSyncDatasLen  int32                                  `json:"robotsSyncDatasLen"`
	RobotsSyncDatas     []S1BattleProto2022ClientRobotSyncData `json:"robotsSyncDatas"`
}

type S1BattleProto2022ConfigTabAck struct {
	Magic              uint32                              `json:"magic"`
	RobotConfigVersion byte                                `json:"robotConfigVersion"`
	Color              byte                                `json:"color"`
	ModuleNum          S1BattleProto2022ModuleNum          `json:"moduleNum"`
	HealthCalc         S1BattleProto2022HealthCalc         `json:"healthCalc"`
	GameLimit          S1BattleProto2022GameLimit          `json:"gameLimit"`
	ArmorData          S1BattleProtot2022RobotArmorCfgData `json:"armorData"`
}

type S1BattleProto2022Energy2019AmorSelfCheck struct {
	Nouse byte `json:"nouse"`
}

type S1BattleProto2022Energy2019ArmorHitUpload struct {
	Row     byte   `json:"row"`
	ArmorId byte   `json:"armorId"`
	HitType uint16 `json:"hitType"`
}

type S1BattleProto2022Energy2019SetArmLight struct {
	ArmorColor []uint32 `json:"armorColor"`
	ArmColor   []uint32 `json:"armColor"`
	Effect     []byte   `json:"effect"`
	ExtVar     []uint16 `json:"extVar"`
	Row        byte     `json:"row"`
}

type S1BattleProto2022Energy2019StateReport struct {
	State       byte    `json:"state"`
	Armors      []byte  `json:"armors"`
	Reserve     byte    `json:"reserve"`
	RotateDeg   uint16  `json:"rotateDeg"`
	RotateSpeed float32 `json:"rotateSpeed"`
}

type S1BattleProto2022Energy2020SetArmRotate struct {
	CtrlMode byte `json:"ctrlMode"`
	Direct   byte `json:"direct"`
}

type S1BattleProto2022EnergyReset struct {
	Res byte `json:"res"`
}

type S1BattleProto2022EnergySetId struct {
	Res byte `json:"res"`
}

type S1BattleProto2022EnergyStateChangeNotify struct {
	RuneId      byte   `json:"runeId"`
	State       byte   `json:"state"`
	RingSum     byte   `json:"ringSum"`
	AtkBuffVal  uint16 `json:"atkBuffVal"`
	DefBuffVal  byte   `json:"defBuffVal"`
	BeforeState byte   `json:"beforeState"`
}

type S1BattleProto2022EnergyStateSync struct {
	RuneId            byte `json:"runeId"`
	Connect           byte `json:"connect"`
	State             byte `json:"state"`
	Round             byte `json:"round"`
	Time              byte `json:"time"`
	Error             byte `json:"error"`
	Armor0            byte `json:"armor0"`
	Armor1            byte `json:"armor1"`
	Armor2            byte `json:"armor2"`
	Armor3            byte `json:"armor3"`
	Armor4            byte `json:"armor4"`
	Armor5            byte `json:"armor5"`
	Motor             byte `json:"motor"`
	AvailbleCountdown byte `json:"availbleCountdown"`
}

type S1BattleProto2022EnerySetLogoLight struct {
	IconRColor uint32 `json:"iconRcolor"`
	Row        byte   `json:"row"`
}

type S1BattleProto2022EnvBaseShellControl struct {
	BaseShellControl byte   `json:"baseShellControl"`
	Rgb              uint16 `json:"rgb"`
}

type S1BattleProto2022EnvDevicesDescripeAck struct {
	Name S1BattleProto2022RobotDynamicUistr `json:"name"`
}

type S1BattleProto2022EnvDevicesDescripeReq struct {
	Res byte `json:"res"`
}

type S1BattleProto2022EnvHeartBeatAck struct {
	Time uint32 `json:"time"`
}

type S1BattleProto2022EnvHeartBeatReq struct {
	Status byte `json:"status"`
}

type S1BattleProto2022ExchangeProClearOreRes struct {
	Res byte `json:"res"`
}

type S1BattleProto2022ExchangeProCtrCmd struct {
	Command    byte   `json:"command"`
	Reserved1  byte   `json:"reserved1"`
	Reserved2  byte   `json:"reserved2"`
	Seq        uint32 `json:"seq"`
	ModuleType byte   `json:"moduleType"`
	ModuleId   byte   `json:"moduleId"`
}

type S1BattleProto2022ExchangeProCtrCmdack struct {
	ErrorCode  byte   `json:"errorCode"`
	Seq        uint32 `json:"seq"`
	ModuleType byte   `json:"moduleType"`
	ModuleId   byte   `json:"moduleId"`
}

type S1BattleProto2022ExchangeProLidarInfo struct {
	State      byte `json:"state"`
	ModuleType byte `json:"moduleType"`
	ModuleId   byte `json:"moduleId"`
}

type S1BattleProto2022ExchangeProPosition struct {
	X          int32  `json:"x"`
	Y          int32  `json:"y"`
	Z          int32  `json:"z"`
	Pitch      int32  `json:"pitch"`
	Roll       int32  `json:"roll"`
	Yaw        int32  `json:"yaw"`
	Seq        uint32 `json:"seq"`
	ModuleType byte   `json:"moduleType"`
	ModuleId   byte   `json:"moduleId"`
}

type S1BattleProto2022ExchangeProRealPosition struct {
	X          int32  `json:"x"`
	Y          int32  `json:"y"`
	Z          int32  `json:"z"`
	Pitch      int32  `json:"pitch"`
	Roll       int32  `json:"roll"`
	Yaw        int32  `json:"yaw"`
	Seq        uint32 `json:"seq"`
	ModuleType byte   `json:"moduleType"`
	ModuleId   byte   `json:"moduleId"`
}

type S1BattleProto2022FlyCatStatus struct {
	SysState    byte `json:"sysState"`
	CtrlState   byte `json:"ctrlState"`
	WorkState   byte `json:"workState"`
	Battery     byte `json:"battery"`
	SensorState byte `json:"sensorState"`
}

type S1BattleProto2022GameLimit struct {
	PowerThreshold        uint16                              `json:"powerThreshold"`
	PowerBuffer1          uint16                              `json:"powerBuffer1"`
	PowerBuffer2          uint16                              `json:"powerBuffer2"`
	PowerMaxhurt          uint16                              `json:"powerMaxhurt"`
	PowerHurtTabCount     int32                               `json:"powerHurtTabCount"`
	PowerHurtTabs         []byte                              `json:"powerHurtTabs"`
	ShooterLimitsCount    int32                               `json:"shooterLimitsCount"`
	ShooterLimits         []S1BattleProto2022ShooterLimit     `json:"shooterLimits"`
	HeatLimitsCount       int32                               `json:"heatLimitsCount"`
	HeatLimits            []S1BattleProto2022HeatLimit        `json:"heatLimits"`
	ShooterFreqLimitCount int32                               `json:"shooterFreqLimitCount"`
	FreqLimits            []S1BattleProto2022ShooterFreqLimit `json:"freqLimits"`
}

type S1BattleProto2022GripperStateNotify struct {
	IsConnect             byte   `json:"isConnect"`
	IsHasMineralsListLen  byte   `json:"isHasMineralsListLen"`
	IsHasMineralsList     []byte `json:"isHasMineralsList"`
	ErrorsListLen         byte   `json:"errorsListLen"`
	ErrorsList            []byte `json:"errorsList"`
	GrippersStatesListLen byte   `json:"grippersStatesListLen"`
	GripperStates         []byte `json:"gripperStates"`
}

type S1BattleProto2022HealthCalc struct {
	HealthPointStart uint16 `json:"healthPointStart"`
	HealthPointFull  uint16 `json:"healthPointFull"`
	BulletHurt       uint16 `json:"bulletHurt"`
	GolfHurt         uint16 `json:"golfHurt"`
	ImpactHurt       uint16 `json:"impactHurt"`
}

type S1BattleProto2022HeatLimit struct {
	HeatSpdThreshold uint16 `json:"heatSpdThreshold"`
	HeatSpdMaxHurt   uint16 `json:"heatSpdMaxHurt"`
	HeatHurtTabCount int32  `json:"heatHurtTabCount"`
	HeatHurtTab      []byte `json:"heatHurtTab"`
	HeatCdFreq       byte   `json:"heatCdFreq"`
	HeatCdStep       uint16 `json:"heatCdStep"`
}

type S1BattleProto2022IoctrCfgSet struct {
	IsDp2SpiMode    byte   `json:"isDp2SpiMode"`
	IsPwmDp2Io2     byte   `json:"isPwmDp2Io2"`
	IsPwmDp2Io5     byte   `json:"isPwmDp2Io5"`
	IsPwmDp2Io6     byte   `json:"isPwmDp2Io6"`
	IsPwmDp2Io7     byte   `json:"isPwmDp2Io7"`
	Dp2PwmFrq       uint16 `json:"dp2PwmFrq"`
	IsOutputDp1     byte   `json:"isOutputDp1"`
	IsTriggerDp1Io0 byte   `json:"isTriggerDp1Io0"`
	IsTriggerDp1Io1 byte   `json:"isTriggerDp1Io1"`
	IsTriggerDp1Io2 byte   `json:"isTriggerDp1Io2"`
	IsTriggerDp1Io3 byte   `json:"isTriggerDp1Io3"`
	IsTriggerDp1Io4 byte   `json:"isTriggerDp1Io4"`
	IsTriggerDp1Io5 byte   `json:"isTriggerDp1Io5"`
	IsTriggerDp1Io6 byte   `json:"isTriggerDp1Io6"`
	IsTriggerDp1Io7 byte   `json:"isTriggerDp1Io7"`
	IsTriggerDp3Io0 byte   `json:"isTriggerDp3Io0"`
	IsTriggerDp3Io1 byte   `json:"isTriggerDp3Io1"`
	IsTriggerDp3Io2 byte   `json:"isTriggerDp3Io2"`
	IsTriggerDp3Io3 byte   `json:"isTriggerDp3Io3"`
	IsTriggerDp3Io4 byte   `json:"isTriggerDp3Io4"`
	IsTriggerDp3Io5 byte   `json:"isTriggerDp3Io5"`
	IsTriggerDp3Io6 byte   `json:"isTriggerDp3Io6"`
	IsTriggerDp3Io7 byte   `json:"isTriggerDp3Io7"`
	ModuleId        byte   `json:"moduleId"`
	ModuleType      byte   `json:"moduleType"`
}

type S1BattleProto2022IoctrCfgSetAck struct {
	ErrCode    byte `json:"errCode"`
	ModuleId   byte `json:"moduleId"`
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetActuator struct {
	Actuator0  byte `json:"actuator0"`
	Actuator1  byte `json:"actuator1"`
	ModuleId   byte `json:"moduleId"`
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetActuatorAck struct {
	Actuator0  byte `json:"actuator0"`
	Actuator1  byte `json:"actuator1"`
	ModuleId   byte `json:"moduleId"`
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetVal struct {
	Dp2OutputValsLen int32    `json:"dp2OutputValsLen"`
	Dp2OutputVal     []uint16 `json:"dp2OutputVal"`
	OutDp1Io0        byte     `json:"outDp1Io0"`
	OutDp1Io1        byte     `json:"outDp1Io1"`
	OutDp1Io2        byte     `json:"outDp1Io2"`
	OutDp1Io3        byte     `json:"outDp1Io3"`
	OutDp1Io4        byte     `json:"outDp1Io4"`
	OutDp1Io5        byte     `json:"outDp1Io5"`
	OutDp1Io6        byte     `json:"outDp1Io6"`
	OutDp1Io7        byte     `json:"outDp1Io7"`
	ModuleId         byte     `json:"moduleId"`
	ModuleType       byte     `json:"moduleType"`
}

type S1BattleProto2022IoctrSetValAck struct {
	ErrCode    byte `json:"errCode"`
	ModuleId   byte `json:"moduleId"`
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrState struct {
	Dp1Gpio0          byte     `json:"dp1Gpio0"`
	Dp1Gpio1          byte     `json:"dp1Gpio1"`
	Dp1Gpio2          byte     `json:"dp1Gpio2"`
	Dp1Gpio3          byte     `json:"dp1Gpio3"`
	Dp1Gpio4          byte     `json:"dp1Gpio4"`
	Dp1Gpio5          byte     `json:"dp1Gpio5"`
	Dp1Gpio6          byte     `json:"dp1Gpio6"`
	Dp1Gpio7          byte     `json:"dp1Gpio7"`
	Dp3Adc1In0        byte     `json:"dp3Adc1In0"`
	Dp3Adc1In1        byte     `json:"dp3Adc1In1"`
	Dp3Adc1In2        byte     `json:"dp3Adc1In2"`
	Dp3Adc1In3        byte     `json:"dp3Adc1In3"`
	Dp3Adc1In4        byte     `json:"dp3Adc1In4"`
	Dp3Adc1In5        byte     `json:"dp3Adc1In5"`
	Dp3Adc1In6        byte     `json:"dp3Adc1In6"`
	Dp3Adc1In7        byte     `json:"dp3Adc1In7"`
	Dp3AdcValsListLen int32    `json:"dp3AdcValsListLen"`
	Dp3AdcVals        []uint16 `json:"dp3AdcVals"`
	IsPwmDp2Io0       byte     `json:"isPwmDp2Io0"`
	IsPwmDp2Io1       byte     `json:"isPwmDp2Io1"`
	IsPwmDp2Io2       byte     `json:"isPwmDp2Io2"`
	IsPwmDp2Io3       byte     `json:"isPwmDp2Io3"`
	IsPwmDp2Io4       byte     `json:"isPwmDp2Io4"`
	IsPwmDp2Io5       byte     `json:"isPwmDp2Io5"`
	IsPwmDp2Io6       byte     `json:"isPwmDp2Io6"`
	IsPwmDp2Io7       byte     `json:"isPwmDp2Io7"`
	Dp2OutValsListLen int32    `json:"dp2OutValsListLen"`
	Dp2OutVars        []uint16 `json:"dp2OutVars"`
	Dp2PwmFrq         uint16   `json:"dp2PwmFrq"`
	Error             byte     `json:"error"`
	ModuleId          byte     `json:"moduleId"`
	ModuleType        byte     `json:"moduleType"`
}

type S1BattleProto2022IoctrTriggerVal struct {
	Dp1InputValBefore byte   `json:"dp1InputValBefore"`
	Dp3InputValBefore byte   `json:"dp3InputValBefore"`
	Dp1InputValAfter  byte   `json:"dp1InputValAfter"`
	Dp3InputValAfter  byte   `json:"dp3InputValAfter"`
	SysTickMs         uint32 `json:"sysTickMs"`
	ModuleId          byte   `json:"moduleId"`
	ModuleType        byte   `json:"moduleType"`
}

type S1BattleProto2022MapClickInfoNotify struct {
	TeamId         byte    `json:"teamId"`
	IsSendAll      byte    `json:"isSendAll"`
	RobotidListLen byte    `json:"robotidListLen"`
	RobotidList    []byte  `json:"robotidList"`
	Mode           byte    `json:"mode"`
	EnemyId        byte    `json:"enemyId"`
	Ascii          byte    `json:"ascii"`
	Type           byte    `json:"type"`
	ScreenX        uint16  `json:"screenX"`
	ScreenY        uint16  `json:"screenY"`
	MapX           float32 `json:"mapX"`
	MapY           float32 `json:"mapY"`
}

type S1BattleProto2022ModuleNum struct {
	SmallStrikeNum byte `json:"smallStrikeNum"`
	BigStrikeNum   byte `json:"bigStrikeNum"`
	LightBarNum    byte `json:"lightBarNum"`
	PowerNum       byte `json:"powerNum"`
	SmallShootNum  byte `json:"smallShootNum"`
	BigShootNum    byte `json:"bigShootNum"`
	CameraNum      byte `json:"cameraNum"`
	RfidNum        byte `json:"rfidNum"`
	UwbNum         byte `json:"uwbNum"`
	SmallArmorImp  byte `json:"smallArmorImp"`
	BigArmorImp    byte `json:"bigArmorImp"`
	LightBrdImp    byte `json:"lightBrdImp"`
	PowerBrdImp    byte `json:"powerBrdImp"`
	SmallShotImp   byte `json:"smallShotImp"`
	BigShotImp     byte `json:"bigShotImp"`
	CameraImp      byte `json:"cameraImp"`
	RfidImp        byte `json:"rfidImp"`
	UwbImp         byte `json:"uwbImp"`
	CapNum         byte `json:"capNum"`
	CapImp         byte `json:"capImp"`
	ExtTypeAnum    byte `json:"extTypeAnum"`
	ExtTypeAimp    byte `json:"extTypeAimp"`
	ExtTypeBnum    byte `json:"extTypeBnum"`
	ExtTypeBimp    byte `json:"extTypeBimp"`
	ExtTypeCnum    byte `json:"extTypeCnum"`
	ExtTypeCimp    byte `json:"extTypeCimp"`
	Res            byte `json:"res"`
}

type S1BattleProto2022PowerSwitchState struct {
	Chassis byte `json:"chassis"`
	Gimbal  byte `json:"gimbal"`
	Shooter byte `json:"shooter"`
}

type S1BattleProto2022QueryRobotInfo struct {
	ModuleType byte `json:"moduleType"`
	ModuleId   byte `json:"moduleId"`
}

type S1BattleProto2022QueryRobotInfoResult struct {
	LoaderVersion   uint32  `json:"loaderVersion"`
	AppVersion      uint32  `json:"appVersion"`
	DeviceIdListLen int32   `json:"deviceIdListLen"`
	DeviceIdList    []int32 `json:"deviceIdList"`
	Reserved        uint32  `json:"reserved"`
	ModuleType      byte    `json:"moduleType"`
	ModuleId        byte    `json:"moduleId"`
}

type S1BattleProto2022ReLoginDataSync struct {
	RoleS0Id     uint64    `json:"roleS0Id"`
	BuffCount    byte      `json:"buffCount"`
	BuffUids     []uint64  `json:"buffUids"`
	BuffTids     []uint32  `json:"buffTids"`
	BuffLevels   []uint32  `json:"buffLevels"`
	BuffVisibles []byte    `json:"buffVisibles"`
	BuffMaxTime  []float32 `json:"buffMaxTime"`
	BuffLeftTime []float32 `json:"buffLeftTime"`
}

type S1BattleProto2022RobotBaseDevCtlCmdAck struct {
	Placeholder byte `json:"placeholder"`
}

type S1BattleProto2022RobotCheckinTimestamp struct {
	Magic     uint32 `json:"magic"`
	DataVer   byte   `json:"dataVer"`
	Timestamp uint32 `json:"timestamp"`
}

type S1BattleProto2022RobotDataSync struct {
	CurHp                 int32     `json:"curHp"`
	MaxHp                 int32     `json:"maxHp"`
	Level                 byte      `json:"level"`
	LightEffectMask       uint32    `json:"lightEffectMask"`
	ShooterHeatsCount     int32     `json:"shooterHeatsCount"`
	ShooterHeats          []float32 `json:"shooterHeats"`
	ShooterHeatThresCount int32     `json:"shooterHeatThresCount"`
	ShooterHeatThres      []float32 `json:"shooterHeatThres"`
	CurHeatCoolCount      int32     `json:"curHeatCoolCount"`
	CurHeatCool           []float32 `json:"curHeatCool"`
	IsGuardAlive          uint16    `json:"isGuardAlive"`
	InventedShieldValue   uint16    `json:"inventedShieldValue"`
	FixPower              uint16    `json:"fixPower"`
	FixSamllShootSpeed1   uint16    `json:"fixSamllShootSpeed1"`
	FixBigShootSpeed      uint16    `json:"fixBigShootSpeed"`
	FixSamllShootSpeed2   uint16    `json:"fixSamllShootSpeed2"`
}

type S1BattleProto2022RobotDynamicUistr struct {
	EnvDeviceXnameEn string `json:"envDeviceXnameEn"`
	EnvDeviceXnameZh string `json:"envDeviceXnameZh"`
	EnvDeviceYnameEn string `json:"envDeviceYnameEn"`
	EnvDeviceYnameZh string `json:"envDeviceYnameZh"`
	EnvDeviceZnameEn string `json:"envDeviceZnameEn"`
	EnvDeviceZnameZh string `json:"envDeviceZnameZh"`
	ExtModuleAnameEn string `json:"extModuleAnameEn"`
	ExtModuleAnameZh string `json:"extModuleAnameZh"`
	ExtModuleBnameEn string `json:"extModuleBnameEn"`
	ExtModuleBnameZh string `json:"extModuleBnameZh"`
	ExtModuleCnameEn string `json:"extModuleCnameEn"`
	ExtModuleCnameZh string `json:"extModuleCnameZh"`
}

type S1BattleProto2022RobotGetCheckCapAck struct {
	CheckedCap       float32 `json:"checkedCap"`
	VoltageFirst     float32 `json:"voltageFirst"`
	VoltageFinal     float32 `json:"voltageFinal"`
	MeasureTime      float32 `json:"measureTime"`
	OutputQEnergy    float32 `json:"outputQenergy"`
	RecharingQEnergy float32 `json:"recharingQenergy"`
	Date             uint32  `json:"date"`
	Time             uint32  `json:"time"`
}

type S1BattleProto2022RobotGetCheckCapReq struct {
	Res byte `json:"res"`
}

type S1BattleProto2022RobotHited struct {
	OnHitType   byte    `json:"onHitType"`
	AttackSpeed uint16  `json:"attackSpeed"`
	ArmorNumber byte    `json:"armorNumber"`
	Press       float32 `json:"press"`
	TimeStamp   uint64  `json:"timeStamp"`
}

type S1BattleProto2022RobotIrcheckReq struct {
	Check byte `json:"check"`
}

type S1BattleProto2022RobotInitCfgTab struct {
	Ext17MmSpeed byte   `json:"ext17MmSpeed"`
	ResCount     int32  `json:"resCount"`
	Res          []byte `json:"res"`
}

type S1BattleProto2022RobotMeasureStartReq struct {
	MaxPrriod  uint32 `json:"maxPrriod"`
	PushPrriod uint32 `json:"pushPrriod"`
}

type S1BattleProto2022RobotMeasureStartRsp struct {
	Type   byte `json:"type"`
	Id     byte `json:"id"`
	Status byte `json:"status"`
}

type S1BattleProto2022RobotMeasureStopReq struct {
	Res byte `json:"res"`
}

type S1BattleProto2022RobotMeasureStopRsp struct {
	Res byte `json:"res"`
}

type S1BattleProto2022RobotMoudleLife struct {
	ModuleId      byte   `json:"moduleId"`
	ModuleType    byte   `json:"moduleType"`
	AckModuleId   byte   `json:"ackModuleId"`
	AckModuleType byte   `json:"ackModuleType"`
	Status        string `json:"status"`
}

type S1BattleProto2022RobotNewHeartBeatAck struct {
	Nouse byte `json:"nouse"`
}

type S1BattleProto2022RobotNewHeartBeatReq struct {
	Nouse byte `json:"nouse"`
}

type S1BattleProto2022RobotNtpServerIpAck struct {
	Ip uint32 `json:"ip"`
}

type S1BattleProto2022RobotNtpServerIpReq struct {
	Res byte `json:"res"`
}

type S1BattleProto2022RobotPushCapinfo struct {
	VoltageFirst     float32 `json:"voltageFirst"`
	VoltageFinal     float32 `json:"voltageFinal"`
	MeasureTime      float32 `json:"measureTime"`
	OutputQEnergy    float32 `json:"outputQenergy"`
	RecharingQEnergy float32 `json:"recharingQenergy"`
	MeasureCap       float32 `json:"measureCap"`
	CheckCap         float32 `json:"checkCap"`
}

type S1BattleProto2022RobotPushCaprtinfo struct {
	Voltage           float32 `json:"voltage"`
	OutputCurrent     float32 `json:"outputCurrent"`
	RechargingCurrent float32 `json:"rechargingCurrent"`
	VoltageFirst      float32 `json:"voltageFirst"`
	MeasureTime       float32 `json:"measureTime"`
	MeasureCap        float32 `json:"measureCap"`
	RecharingQEnergy  float32 `json:"recharingQenergy"`
	OutputQEnergy     float32 `json:"outputQenergy"`
}

type S1BattleProto2022RobotPushSensorInfo struct {
	Voltage           float32 `json:"voltage"`
	OutputCurrent     float32 `json:"outputCurrent"`
	RechargingCurrent float32 `json:"rechargingCurrent"`
	Status            byte    `json:"status"`
}

type S1BattleProto2022RobotResurgenceNotify struct {
	RobotId byte `json:"robotId"`
}

type S1BattleProto2022RobotStatus struct {
	RobotStatusData S1BattleProto2022RobotStatusData `json:"robotStatusData"`
}

type S1BattleProto2022RobotStatusData struct {
	Hp             uint16                             `json:"hp"`
	Voltage        float32                            `json:"voltage"`
	Current        float32                            `json:"current"`
	ChassisPower   float32                            `json:"chassisPower"`
	PowerBuffer    float32                            `json:"powerBuffer"`
	GimbalVoltage  float32                            `json:"gimbalVoltage"`
	GimbalPower    float32                            `json:"gimbalPower"`
	ShooterVoltage float32                            `json:"shooterVoltage"`
	ShooterPower   float32                            `json:"shooterPower"`
	Rssi           byte                               `json:"rssi"`
	HwId           uint32                             `json:"hwId"`
	ShooterPitch   int16                              `json:"shooterPitch"`
	ShooterRoll    int16                              `json:"shooterRoll"`
	ShooterYaw     int16                              `json:"shooterYaw"`
	Status         S1BattleProto2022RobotSystemStatus `json:"status"`
	Uwb            S1BattleProto2022UwbData           `json:"uwb"`
	PowerState     S1BattleProto2022PowerSwitchState  `json:"powerState"`
	TimeStamp      uint64                             `json:"timeStamp"`
}

type S1BattleProto2022RobotSystemStatus struct {
	Power          byte   `json:"power"`
	Rfid0          byte   `json:"rfid0"`
	Rfid1          byte   `json:"rfid1"`
	Camera         byte   `json:"camera"`
	Blood0         byte   `json:"blood0"`
	Blood1         byte   `json:"blood1"`
	ShooterDetect0 byte   `json:"shooterDetect0"`
	ShooterDetect1 byte   `json:"shooterDetect1"`
	ShooterDetect2 byte   `json:"shooterDetect2"`
	Uwb            byte   `json:"uwb"`
	Armor0         byte   `json:"armor0"`
	Armor1         byte   `json:"armor1"`
	Armor2         byte   `json:"armor2"`
	Armor3         byte   `json:"armor3"`
	Armor4         byte   `json:"armor4"`
	Armor5         byte   `json:"armor5"`
	Armor6         byte   `json:"armor6"`
	Armor7         byte   `json:"armor7"`
	Armor8         byte   `json:"armor8"`
	Armor9         byte   `json:"armor9"`
	Cap            byte   `json:"cap"`
	ExtA0          byte   `json:"extA0"`
	ExtA1          byte   `json:"extA1"`
	ExtA2          byte   `json:"extA2"`
	ExtA3          byte   `json:"extA3"`
	ExtA4          byte   `json:"extA4"`
	ExtA5          byte   `json:"extA5"`
	ExtA6          byte   `json:"extA6"`
	ExtB0          byte   `json:"extB0"`
	ExtB1          byte   `json:"extB1"`
	ExtB2          byte   `json:"extB2"`
	ExtC0          byte   `json:"extC0"`
	ExtC1          byte   `json:"extC1"`
	ExtC2          byte   `json:"extC2"`
	ReservedCount  int32  `json:"reservedCount"`
	ReservedList   []byte `json:"reservedList"`
}

type S1BattleProto2022RobotVtmsettingFull struct {
	ListLen byte   `json:"listLen"`
	Tds1    []byte `json:"tds1"`
	Tds2    []byte `json:"tds2"`
	Tds3    []byte `json:"tds3"`
}

type S1BattleProto2022RobotWeaponFire struct {
	ShooterId  byte    `json:"shooterId"`
	BulletType byte    `json:"bulletType"`
	Speed      float32 `json:"speed"`
	Frequency  float32 `json:"frequency"`
	Yaw        float32 `json:"yaw"`
	Pitch      float32 `json:"pitch"`
	RollAngle  float32 `json:"rollAngle"`
	TimeStamp  uint64  `json:"timeStamp"`
}

type S1BattleProto2022RobotsDataSync struct {
	Progress   byte                           `json:"progress"`
	TimeRemain uint16                         `json:"timeRemain"`
	Utc        uint32                         `json:"utc"`
	RobotData  S1BattleProto2022RobotDataSync `json:"robotData"`
	ModuleId   byte                           `json:"moduleId"`
	ModuleType byte                           `json:"moduleType"`
}

type S1BattleProto2022S2CPowerStateNotify struct {
	ClientId          byte  `json:"clientId"`
	ChassisPowerState int32 `json:"chassisPowerState"`
	GimbalPowerState  int32 `json:"gimbalPowerState"`
	ShooterPowerState int32 `json:"shooterPowerState"`
}

type S1BattleProto2022SendDartRobotStatus struct {
	Connectstatus byte `json:"connectstatus"`
}

type S1BattleProto2022SetFlyCatLight struct {
	Lightcode uint32 `json:"lightcode"`
}

type S1BattleProto2022SetFlyCatLightAck struct {
	RecvedLightcode uint32 `json:"recvedLightcode"`
}

type S1BattleProto2022SetFlyCatStatus struct {
	WorkState byte `json:"workState"`
}

type S1BattleProto2022SetFlyCatStatusAck struct {
	RecvedWorkState byte `json:"recvedWorkState"`
}

type S1BattleProto2022SetSupplyStateReq struct {
	ReportTime uint16 `json:"reportTime"`
	ReportReq  byte   `json:"reportReq"`
}

type S1BattleProto2022ShooterFreqLimit struct {
	ShooterFreqThreshold    float32 `json:"shooterFreqThreshold"`
	ShooterFreqMaxHurt      uint16  `json:"shooterFreqMaxHurt"`
	ShooterFreqHurtTabCount int32   `json:"shooterFreqHurtTabCount"`
	ShooterFreqHurtTab      []byte  `json:"shooterFreqHurtTab"`
}

type S1BattleProto2022ShooterLimit struct {
	ShooterSpdThreshold    float32 `json:"shooterSpdThreshold"`
	ShooterSpdMaxHurt      uint16  `json:"shooterSpdMaxHurt"`
	ShooterSqdHurtTabCount int32   `json:"shooterSqdHurtTabCount"`
	ShooterSqdHurtTab      []byte  `json:"shooterSqdHurtTab"`
}

type S1BattleProto2022SiloCtrLedData struct {
	R byte `json:"r"`
	G byte `json:"g"`
	B byte `json:"b"`
	A byte `json:"a"`
}

type S1BattleProto2022SiloDevStatus struct {
	Online                  byte    `json:"online"`
	DoorStatus              byte    `json:"doorStatus"`
	FloorStatus             byte    `json:"floorStatus"`
	Speed                   uint16  `json:"speed"`
	Angle                   uint16  `json:"angle"`
	Target                  byte    `json:"target"`
	CountdownTime           float32 `json:"countdownTime"`
	ShootCheckCountdownTime float32 `json:"shootCheckCountdownTime"`
	OpenedTimeCount         int32   `json:"openedTimeCount"`
	IsDetectionWindows      byte    `json:"isDetectionWindows"`
	MissileCount            int32   `json:"missileCount"`
	IsCanOpen               byte    `json:"isCanOpen"`
	SiloCooldown            float32 `json:"siloCooldown"`
	SiloErrorCode           byte    `json:"siloErrorCode"`
	MissileHitCount         byte    `json:"missileHitCount"`
	ErrorGateSensorUp       byte    `json:"errorGateSensorUp"`
	ErrorGateSensorDown     byte    `json:"errorGateSensorDown"`
	ErrorBaseBoardSensor    byte    `json:"errorBaseBoardSensor"`
	ErrorGateSensorBoth     byte    `json:"errorGateSensorBoth"`
	ErrorMotorBroken        byte    `json:"errorMotorBroken"`
	ErrorBrakeConstNo       byte    `json:"errorBrakeConstNo"`
	ErrorBrakeConstYes      byte    `json:"errorBrakeConstYes"`
	ErrorFacingObstacle     byte    `json:"errorFacingObstacle"`
	ErrorEmergencyStop      byte    `json:"errorEmergencyStop"`
	ErrorMotorOverHeat      byte    `json:"errorMotorOverHeat"`
	ErrorCloseOverTime      byte    `json:"errorCloseOverTime"`
	ErrorOpenOverTime       byte    `json:"errorOpenOverTime"`
}

type S1BattleProto2022SiloDevStatusSyncData struct {
	SiloDevStatusSyncDataListLen int32                            `json:"siloDevStatusSyncDataListLen"`
	SiloDevStatusDataList        []S1BattleProto2022SiloDevStatus `json:"siloDevStatusDataList"`
}

type S1BattleProto2022StuAreialEnergy struct {
	Cmd        uint16 `json:"cmd"`
	RemainTime byte   `json:"remainTime"`
}

type S1BattleProto2022StuClientRecvInfo struct {
	TargetRobotId uint16  `json:"targetRobotId"`
	TargetPosX    float32 `json:"targetPosX"`
	TargetPosY    float32 `json:"targetPosY"`
}

type S1BattleProto2022StuClientSendCmd struct {
	Cmd           uint16  `json:"cmd"`
	TargetPosX    float32 `json:"targetPosX"`
	TargetPosY    float32 `json:"targetPosY"`
	TargetPosZ    float32 `json:"targetPosZ"`
	CmdKeyboard   byte    `json:"cmdKeyboard"`
	TargetRobotId uint16  `json:"targetRobotId"`
}

type S1BattleProto2022StuCommunication struct {
	DataId     uint16 `json:"dataId"`
	SenderId   uint16 `json:"senderId"`
	ReceiverId uint16 `json:"receiverId"`
}

type S1BattleProto2022StuCommunicationByteData struct {
	Cmd             uint16 `json:"cmd"`
	ByteDataListLen byte   `json:"byteDataListLen"`
	ByteDataList    []byte `json:"byteDataList"`
	S0HeaderBodyLen uint16 `json:"s0HeaderBodyLen"`
	Dataid          uint16 `json:"dataid"`
	Senderid        uint16 `json:"senderid"`
	Receiverid      uint16 `json:"receiverid"`
}

type S1BattleProto2022StuCustomControlData struct {
	Cmd     uint16 `json:"cmd"`
	ListLen byte   `json:"listLen"`
	Data    []byte `json:"data"`
}

type S1BattleProto2022StuEnvStatus struct {
	Cmd                     uint16 `json:"cmd"`
	Num1AddBloodPointStatus uint32 `json:"num1AddBloodPointStatus"`
	Num2AddBloodPointStatus uint32 `json:"num2AddBloodPointStatus"`
	Num3AddBloodPointStatus uint32 `json:"num3AddBloodPointStatus"`
	RuneRfidStatus          uint32 `json:"runeRfidStatus"`
	SmallRuneStatus         uint32 `json:"smallRuneStatus"`
	BigRuneStatus           uint32 `json:"bigRuneStatus"`
	RingUpland              uint32 `json:"ringUpland"`
	TrapezoidR3             uint32 `json:"trapezoidR3"`
	TrapezoidR4             uint32 `json:"trapezoidR4"`
	BaseShieldStatus        uint32 `json:"baseShieldStatus"`
	OutpostStatus           uint32 `json:"outpostStatus"`
}

type S1BattleProto2022StuGameResult struct {
	Cmd    uint16 `json:"cmd"`
	Winner byte   `json:"winner"`
}

type S1BattleProto2022StuGameStatus struct {
	Cmd           uint16 `json:"cmd"`
	GameMode      byte   `json:"gameMode"`
	GameStatus    byte   `json:"gameStatus"`
	RemainTime    uint16 `json:"remainTime"`
	SyncTimeStamp uint64 `json:"syncTimeStamp"`
}

type S1BattleProto2022StuIcrabuffDebuffZoneStatus struct {
	Cmd          uint16 `json:"cmd"`
	F1Status     byte   `json:"f1Status"`
	F1StatusInfo byte   `json:"f1StatusInfo"`
	F2Status     byte   `json:"f2Status"`
	F2StatusInfo byte   `json:"f2StatusInfo"`
	F3Status     byte   `json:"f3Status"`
	F3StatusInfo byte   `json:"f3StatusInfo"`
	F4Status     byte   `json:"f4Status"`
	F4StatusInfo byte   `json:"f4StatusInfo"`
	F5Status     byte   `json:"f5Status"`
	F5StatusInfo byte   `json:"f5StatusInfo"`
	F6Status     byte   `json:"f6Status"`
	F6StatusInfo byte   `json:"f6StatusInfo"`
}

type S1BattleProto2022StuLeftBullet struct {
	Cmd                      uint16 `json:"cmd"`
	SmallAvaliableBulletsNum uint16 `json:"smallAvaliableBulletsNum"`
	BigAvaliableBulletsNum   uint16 `json:"bigAvaliableBulletsNum"`
	LeftCoinsNum             uint16 `json:"leftCoinsNum"`
}

type S1BattleProto2022StuMissileRemainingTime struct {
	Cmdid uint16 `json:"cmdid"`
	Time  byte   `json:"time"`
}

type S1BattleProto2022StuRfidstatus struct {
	Cmd         uint16 `json:"cmd"`
	BaseArea    uint32 `json:"baseArea"`
	UplandArea  uint32 `json:"uplandArea"`
	RuneArea    uint32 `json:"runeArea"`
	FlyArea     uint32 `json:"flyArea"`
	OutpostArea uint32 `json:"outpostArea"`
	IslandArea  uint32 `json:"islandArea"`
	HomeArea    uint32 `json:"homeArea"`
	SapperRfid  uint32 `json:"sapperRfid"`
}

type S1BattleProto2022StuRobotBuff struct {
	Cmd         uint16 `json:"cmd"`
	HealBuff    byte   `json:"healBuff"`
	CoolBuff    byte   `json:"coolBuff"`
	DefenceBuff byte   `json:"defenceBuff"`
	AttackBuff  byte   `json:"attackBuff"`
}

type S1BattleProto2022StuRobotCurrentHp struct {
	Cmd           uint16 `json:"cmd"`
	RedHero       uint16 `json:"redHero"`
	RedSapper     uint16 `json:"redSapper"`
	RedInfantry1  uint16 `json:"redInfantry1"`
	RedInfantry2  uint16 `json:"redInfantry2"`
	RedInfantry3  uint16 `json:"redInfantry3"`
	RedGuard      uint16 `json:"redGuard"`
	RedOutpost    uint16 `json:"redOutpost"`
	RedBase       uint16 `json:"redBase"`
	BlueHero      uint16 `json:"blueHero"`
	BlueSapper    uint16 `json:"blueSapper"`
	BlueInfantry1 uint16 `json:"blueInfantry1"`
	BlueInfantry2 uint16 `json:"blueInfantry2"`
	BlueInfantry3 uint16 `json:"blueInfantry3"`
	BlueGuard     uint16 `json:"blueGuard"`
	BlueOutpost   uint16 `json:"blueOutpost"`
	BlueBase      uint16 `json:"blueBase"`
}

type S1BattleProto2022StuRobotHurt struct {
	Cmd      uint16 `json:"cmd"`
	Armor    byte   `json:"armor"`
	Hurttype byte   `json:"hurttype"`
}

type S1BattleProto2022StuSiloInfo struct {
	Cmd                  uint16 `json:"cmd"`
	LaunchOpeningStatus  byte   `json:"launchOpeningStatus"`
	AttackTarget         byte   `json:"attackTarget"`
	TargetChangeTime     uint16 `json:"targetChangeTime"`
	OperateLaunchCmdTime uint16 `json:"operateLaunchCmdTime"`
}

type S1BattleProto2022StuSupplyStatus struct {
	Cmd             uint16 `json:"cmd"`
	SupplyEnvId     byte   `json:"supplyEnvId"`
	SupplyRobotId   byte   `json:"supplyRobotId"`
	SupplyEnvStatus byte   `json:"supplyEnvStatus"`
	BulletsNum      byte   `json:"bulletsNum"`
}

type S1BattleProto2022StuWarning struct {
	Cmd          uint16 `json:"cmd"`
	WarningLevel byte   `json:"warningLevel"`
	WarningRobot byte   `json:"warningRobot"`
}

type S1BattleProto2022SupplyClearScaleAck struct {
	Placeholder byte `json:"placeholder"`
}

type S1BattleProto2022SupplyClearScaleReq struct {
	Placeholder byte `json:"placeholder"`
}

type S1BattleProto2022SupplyExportAck struct {
	LastState byte `json:"lastState"`
	NowState  byte `json:"nowState"`
}

type S1BattleProto2022SupplyExportReq struct {
	ExportEnable byte `json:"exportEnable"`
}

type S1BattleProto2022SupplyFreedAck struct {
	IsAccept byte `json:"isAccept"`
}

type S1BattleProto2022SupplyFreedReq struct {
	BoxFreedNum byte `json:"boxFreedNum"`
}

type S1BattleProto2022SupplyReloadAck struct {
	IsAccept byte `json:"isAccept"`
	Action   byte `json:"action"`
}

type S1BattleProto2022SupplyReloadReq struct {
	ReloadEnable byte `json:"reloadEnable"`
}

type S1BattleProto2022SupplyReport struct {
	BulletGear0   byte `json:"bulletGear0"`
	BulletGear1   byte `json:"bulletGear1"`
	Loadselector0 byte `json:"loadselector0"`
	Box0          byte `json:"box0"`
	Box1          byte `json:"box1"`
	Box2          byte `json:"box2"`
	Relase        byte `json:"relase"`
	Scales        byte `json:"scales"`
	State         byte `json:"state"`
	BoxReadyNum   byte `json:"boxReadyNum"`
	BoxFreedNum   byte `json:"boxFreedNum"`
}

type S1BattleProto2022SupplyResetReq struct {
	Placeholder byte `json:"placeholder"`
}

type S1BattleProto2022SupplyRestartReq struct {
	Placeholder byte `json:"placeholder"`
}

type S1BattleProto2022TestTimeDelayDownComplete struct {
	Count uint32 `json:"count"`
}

type S1BattleProto2022TestTimeDelayDownData struct {
	Seq      uint32 `json:"seq"`
	CurTime  int64  `json:"curTime"`
	WifiTime uint32 `json:"wifiTime"`
	DataLen  int32  `json:"dataLen"`
	Datas    []byte `json:"datas"`
}

type S1BattleProto2022TestTimeDelayDownRecord struct {
	GlobalDelay      uint32   `json:"globalDelay"`
	UartDelay        uint32   `json:"uartDelay"`
	UartLossCnt      byte     `json:"uartLossCnt"`
	WifiLossCnt      byte     `json:"wifiLossCnt"`
	WifiLossTabCount int32    `json:"wifiLossTabCount"`
	WifiLossTabS     []uint16 `json:"wifiLossTabS"`
}

type S1BattleProto2022UwbData struct {
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	Z          float32 `json:"z"`
	Compass    float32 `json:"compass"`
	AnchorMask byte    `json:"anchorMask"`
	WifiEn     byte    `json:"wifiEn"`
	Rsv0       byte    `json:"rsv0"`
	Rsv1       byte    `json:"rsv1"`
}

type S1BattleProto2022VtmSpeedSync struct {
	Clientid           int32 `json:"clientid"`
	CurrentFreq        int32 `json:"currentFreq"`
	TxConnect          int32 `json:"txConnect"`
	CurrentSpeedRate   int32 `json:"currentSpeedRate"`
	CurrentCountryCode int32 `json:"currentCountryCode"`
}

type S1BattleProto2022IoctrCfgReq struct {
	Res        byte `json:"res"`
	ModuleId   byte `json:"moduleId"`
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrRmMotorsStatus struct {
	RealPositionListLen    int32   `json:"realPositionListLen"`
	RealPositions          []int64 `json:"realPositions"`
	RealSpeedListLen       int32   `json:"realSpeedListLen"`
	RealSpeeds             []int16 `json:"realSpeeds"`
	RealTemperatureListLen int32   `json:"realTemperatureListLen"`
	RealTemperatures       []byte  `json:"realTemperatures"`
	ModuleId               byte    `json:"moduleId"`
	ModuleType             byte    `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType1 struct {
	StartLedIndex uint16                    `json:"startLedIndex"`
	LedNum        uint16                    `json:"ledNum"`
	Color         S1BattleProto2022LedColor `json:"color"`
	Seq           uint32                    `json:"seq"`
	ModuleId      byte                      `json:"moduleId"`
	ModuleType    byte                      `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType1Ack struct {
	ErrCode    byte   `json:"errCode"`
	SeqAck     uint32 `json:"seqAck"`
	ModuleId   byte   `json:"moduleId"`
	ModuleType byte   `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType2 struct {
	StartLedIndex uint16                      `json:"startLedIndex"`
	LedNum        uint16                      `json:"ledNum"`
	ListLen       int32                       `json:"listLen"`
	LedsColors    []S1BattleProto2022LedColor `json:"ledsColors"`
	Seq           uint32                      `json:"seq"`
	ModuleId      byte                        `json:"moduleId"`
	ModuleType    byte                        `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType2Ack struct {
	ErrCode    byte   `json:"errCode"`
	SeqAck     uint32 `json:"seqAck"`
	ModuleId   byte   `json:"moduleId"`
	ModuleType byte   `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType3 struct {
	LightEffect byte   `json:"lightEffect"`
	TimeSpan    uint16 `json:"timeSpan"`
	Seq         uint32 `json:"seq"`
	ModuleId    byte   `json:"moduleId"`
	ModuleType  byte   `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType3Ack struct {
	ErrCode    byte   `json:"errCode"`
	SeqAck     uint32 `json:"seqAck"`
	ModuleId   byte   `json:"moduleId"`
	ModuleType byte   `json:"moduleType"`
}

type S1BattleProto2022IoctrSetRmMotorsCfg struct {
	SpeedPidListLen       int32                              `json:"speedPidListLen"`
	SpeedPids             []S1BattleProto2022SpeedModePid    `json:"speedPids"`
	PositionPidListLen    int32                              `json:"positionPidListLen"`
	PositionPids          []S1BattleProto2022PositionModePid `json:"positionPids"`
	TransRatioListLen     int32                              `json:"transRatioListLen"`
	TransRatios           []float32                          `json:"transRatios"`
	IsStatusReturnListLen int32                              `json:"isStatusReturnListLen"`
	IsStatusReturn        []byte                             `json:"isStatusReturn"`
	Seq                   uint32                             `json:"seq"`
	ModuleId              byte                               `json:"moduleId"`
	ModuleType            byte                               `json:"moduleType"`
}

type S1BattleProto2022IoctrSetRmMotorsCfgAck struct {
	ErrCode    byte   `json:"errCode"`
	SeqAck     uint32 `json:"seqAck"`
	ModuleId   byte   `json:"moduleId"`
	ModuleType byte   `json:"moduleType"`
}

type S1BattleProto2022IoctrSetRmMotorsMoveVal struct {
	LoopModeListLen int32   `json:"loopModeListLen"`
	LoopMode        []byte  `json:"loopMode"`
	Len             int32   `json:"len"`
	ExpectVal       []int64 `json:"expectVal"`
	ModuleId        byte    `json:"moduleId"`
	ModuleType      byte    `json:"moduleType"`
}

type S1BattleProto2022LedColor struct {
	R byte `json:"r"`
	G byte `json:"g"`
	B byte `json:"b"`
}

type S1BattleProto2022PositionModePid struct {
	Kp             float32 `json:"kp"`
	Ki             float32 `json:"ki"`
	Kd             float32 `json:"kd"`
	NoResponse     byte    `json:"noResponse"`
	IntergralLimit uint16  `json:"intergralLimit"`
	MaxSpeed       uint16  `json:"maxSpeed"`
}

type S1BattleProto2022SiloCtrlDoor struct {
	Cmd byte `json:"cmd"`
}

type S1BattleProto2022SiloStatus struct {
	DoorStatus  byte `json:"doorStatus"`
	FloorStatus byte `json:"floorStatus"`
	Errorcode   byte `json:"errorcode"`
}

type S1BattleProto2022SpeedModePid struct {
	Kp        float32 `json:"kp"`
	Ki        float32 `json:"ki"`
	Kd        float32 `json:"kd"`
	MaxOutput uint16  `json:"maxOutput"`
}

type S1BattleProto2023ClientChangeTokenCmd struct {
	Len      uint16 `json:"len"`
	NewToken string `json:"newToken"`
}

type S1BattleProto2023ClientGameSystemInfoNotify struct {
	TokenLen      uint16 `json:"tokenLen"`
	CurMatchToken string `json:"curMatchToken"`
}

type S1BattleProto2023ClientGuardInPatrolInfoSync struct {
	RobotId                  byte    `json:"robotId"`
	IsCurrentInPatrolArea    byte    `json:"isCurrentInPatrolArea"`
	IsLeavePatrolAreaTimeout byte    `json:"isLeavePatrolAreaTimeout"`
	LeavePatrolAreaLeftTime  float32 `json:"leavePatrolAreaLeftTime"`
	IsCountDownActive        byte    `json:"isCountDownActive"`
}

type S1BattleProto2023ClientModuleVersionData struct {
	MoudleId          byte   `json:"moudleId"`
	MoudleType        byte   `json:"moudleType"`
	MoudleSubType     byte   `json:"moudleSubType"`
	NewMoudleVersion  string `json:"newMoudleVersion"`
	CurMoudleVersion  string `json:"curMoudleVersion"`
	VersionState      byte   `json:"versionState"`
	MoudleIsImportant byte   `json:"moudleIsImportant"`
}

type S1BattleProto2023ClientModuleVersionState struct {
	RobotId       byte                                       `json:"robotId"`
	ModuleNum     byte                                       `json:"moduleNum"`
	ModuleNumMax  byte                                       `json:"moduleNumMax"`
	ModuleVersion []S1BattleProto2023ClientModuleVersionData `json:"moduleVersion"`
}

type S1BattleProto2023ClientPenaltyInfo struct {
	RobotId       byte   `json:"robotId"`
	PenaltyType   byte   `json:"penaltyType"`
	PenaltyReason string `json:"penaltyReason"`
}

type S1BattleProto2023ClientPenaltyTableInfos struct {
	UploadType byte                                 `json:"uploadType"`
	InfosLen   byte                                 `json:"infosLen"`
	Infos      []S1BattleProto2023ClientPenaltyInfo `json:"infos"`
}

type S1BattleProto2023ClientPenaltyTableInfosAck struct {
	State      byte `json:"state"`
	UploadType byte `json:"uploadType"`
}

type S1BattleProto2023ClientRobotBattleInfo struct {
	Robots0Id                    byte    `json:"robots0Id"`
	BattleState                  byte    `json:"battleState"`
	CanRemoteHeal                byte    `json:"canRemoteHeal"`
	CanRemoteExchangeSmallBullet byte    `json:"canRemoteExchangeSmallBullet"`
	CanRemoteExchangeBigBullet   byte    `json:"canRemoteExchangeBigBullet"`
	RemaindBuyRevivalCount       byte    `json:"remaindBuyRevivalCount"`
	BuyRevivalPrice              int32   `json:"buyRevivalPrice"`
	BattleStateCountDown         float32 `json:"battleStateCountDown"`
	IsOnHpPoint                  byte    `json:"isOnHpPoint"`
	ChassisPowerCtrl             byte    `json:"chassisPowerCtrl"`
	GimbalPowerCtrl              byte    `json:"gimbalPowerCtrl"`
	ShooterPowerCtrl             byte    `json:"shooterPowerCtrl"`
}

type S1BattleProto2023ClientServerRobotsBattleInfoSync struct {
	RobotsNum                 byte                                     `json:"robotsNum"`
	RobotsBattleInfoSyncDatas []S1BattleProto2023ClientRobotBattleInfo `json:"robotsBattleInfoSyncDatas"`
}

type S1BattleProto2023ClientTeamSupplyInfoSync struct {
	Teamid                          byte    `json:"teamid"`
	ExchangeSmallBulletPackageCount byte    `json:"exchangeSmallBulletPackageCount"`
	ExchangeBigBulletPackageCount   byte    `json:"exchangeBigBulletPackageCount"`
	ExchangeRemoteHealCount         byte    `json:"exchangeRemoteHealCount"`
	SentryRemainRevivalCount        byte    `json:"sentryRemainRevivalCount"`
	SmallBulletPackageUnitPrice     float32 `json:"smallBulletPackageUnitPrice"`
	BigBulletPackageUnitPrice       float32 `json:"bigBulletPackageUnitPrice"`
	BigBulletUnitPrice              float32 `json:"bigBulletUnitPrice"`
	SmallBulletUnitPrice            float32 `json:"smallBulletUnitPrice"`
	CurrentRemoteHealPrice          float32 `json:"currentRemoteHealPrice"`
	SentryControlPrice              float32 `json:"sentryControlPrice"`
}

type S1BattleProto2023ExchangeProErrorInfoNotify struct {
	TeamId                   byte   `json:"teamId"`
	IsOnline                 byte   `json:"isOnline"`
	ArmErrorCode             uint16 `json:"armErrorCode"`
	RfidErrorCode            byte   `json:"rfidErrorCode"`
	RunningState             byte   `json:"runningState"`
	Ir1State                 byte   `json:"ir1State"`
	Ir2State                 byte   `json:"ir2State"`
	Ir3State                 byte   `json:"ir3State"`
	IrErrorCode              byte   `json:"irErrorCode"`
	TempAlertVal             int16  `json:"tempAlertVal"`
	TempRmMotorRoll          int16  `json:"tempRmMotorRoll"`
	TempRmMotorPitch         int16  `json:"tempRmMotorPitch"`
	TempRmMotorPush          int16  `json:"tempRmMotorPush"`
	MotorRollState           byte   `json:"motorRollState"`
	MotorPitchState          byte   `json:"motorPitchState"`
	MotorPushState           byte   `json:"motorPushState"`
	IsSoftwareEmergencyState byte   `json:"isSoftwareEmergencyState"`
	IsHardwareEmergencyState byte   `json:"isHardwareEmergencyState"`
}

type S1BattleProtoAddBulletNotify struct {
	Uid         uint64 `json:"uid"`
	BulletIndex byte   `json:"bulletIndex"`
}

type S1BattleProtoArmorHitAck struct {
	Result int32 `json:"result"`
}

type S1BattleProtoArmorHitReq struct {
	Index       int32  `json:"index"`
	AttackType  int32  `json:"attackType"`
	AttackerUid uint64 `json:"attackerUid"`
	AccValue    int32  `json:"accValue"`
	MicValue    int32  `json:"micValue"`
}

type S1BattleProtoAttrsNotify struct {
	Uid           uint64  `json:"uid"`
	Tid           uint32  `json:"tid"`
	Teamid        byte    `json:"teamid"`
	SeatIndex     byte    `json:"seatIndex"`
	Hp            int32   `json:"hp"`
	HpMax         int32   `json:"hpMax"`
	Armor         int32   `json:"armor"`
	ArmorMax      int32   `json:"armorMax"`
	Durability    int32   `json:"durability"`
	DurabilityMax int32   `json:"durabilityMax"`
	Bullet        int32   `json:"bullet"`
	BulletMax     int32   `json:"bulletMax"`
	Attack1       int32   `json:"attack1"`
	Attack2       float32 `json:"attack2"`
}

type S1BattleProtoAutoStateNotify struct {
	State int32 `json:"state"`
}

type S1BattleProtoBeAttackNotify struct {
	Index         int32  `json:"index"`
	AttackType    int32  `json:"attackType"`
	AttackerUid   uint64 `json:"attackerUid"`
	BeAttackerUid uint64 `json:"beAttackerUid"`
	Damage        int32  `json:"damage"`
}

type S1BattleProtoBuffAddNotify struct {
	PlayerUid    uint64  `json:"playerUid"`
	BuffUid      uint64  `json:"buffUid"`
	BuffTid      uint32  `json:"buffTid"`
	BuffLevel    uint32  `json:"buffLevel"`
	BuffVisible  byte    `json:"buffVisible"`
	BuffMaxTime  float32 `json:"buffMaxTime"`
	BuffLeftTime float32 `json:"buffLeftTime"`
	MsgParams    string  `json:"msgParams"`
}

type S1BattleProtoBuffDelNotify struct {
	PlayerUid uint64 `json:"playerUid"`
	BuffUid   uint64 `json:"buffUid"`
}

type S1BattleProtoBuffModifyNotify struct {
	PlayerUid      uint64  `json:"playerUid"`
	BuffUid        uint64  `json:"buffUid"`
	BuffTid        uint32  `json:"buffTid"`
	BuffLevel      uint32  `json:"buffLevel"`
	BuffVisible    byte    `json:"buffVisible"`
	BuffMaxTime    float32 `json:"buffMaxTime"`
	BuffLeftTime   float32 `json:"buffLeftTime"`
	UpdateProgress byte    `json:"updateProgress"`
	MsgParams      string  `json:"msgParams"`
}

type S1BattleProtoClientSyncNotify struct {
	Uid             uint64 `json:"uid"`
	ConnClientState int32  `json:"connClientState"`
	ConnRobotState  int32  `json:"connRobotState"`
	Battery         int32  `json:"battery"`
	SignalQuality   int32  `json:"signalQuality"`
	BattleMode      byte   `json:"battleMode"`
}

type S1BattleProtoClientSyncReq struct {
	Uid             uint64 `json:"uid"`
	ConnClientState int32  `json:"connClientState"`
	ConnRobotState  int32  `json:"connRobotState"`
	Battery         int32  `json:"battery"`
	SignalQuality   int32  `json:"signalQuality"`
	BattleMode      byte   `json:"battleMode"`
}

type S1BattleProtoDeviceBulletNotify struct {
	BulletBottleCount int32    `json:"bulletBottleCount"`
	BulletBottleState []byte   `json:"bulletBottleState"`
	PlayersUid        []uint64 `json:"playersUid"`
}

type S1BattleProtoDeviceModuleNotify struct {
	Uid                 uint64   `json:"uid"`
	ProductType         uint16   `json:"productType"`
	AllDeviceCount      uint16   `json:"allDeviceCount"`
	AllDevice           []uint64 `json:"allDevice"`
	AllDeviceConnection []byte   `json:"allDeviceConnection"`
	AllDevicePrority    []byte   `json:"allDevicePrority"`
}

type S1BattleProtoDeviceModuleReq struct {
	Uid                  uint64   `json:"uid"`
	ProductType          uint16   `json:"productType"`
	ConnectedDeviceCount uint16   `json:"connectedDeviceCount"`
	ConnectedDevice      []uint64 `json:"connectedDevice"`
}

type S1BattleProtoFullSceneDataReq struct {
	Nouse byte `json:"nouse"`
}

type S1BattleProtoGameSettlementNotify struct {
	TeamidWin       int32    `json:"teamidWin"`
	PlayersCount    int32    `json:"playersCount"`
	PlayersUid      []uint64 `json:"playersUid"`
	PlayersHp       []int32  `json:"playersHp"`
	PlayersHpMax    []int32  `json:"playersHpMax"`
	PlayersBekilled []int32  `json:"playersBekilled"`
	Winreason       int32    `json:"winreason"`
	TeamCount       int32    `json:"teamCount"`
	TeamHurtSum     []int32  `json:"teamHurtSum"`
	TeamAtkCount    []int32  `json:"teamAtkCount"`
	TeamPojiaCount  []int32  `json:"teamPojiaCount"`
	DurationTime    int32    `json:"durationTime"`
}

type S1BattleProtoGunBulletNotify struct {
	Uid          uint64 `json:"uid"`
	GunBulletMax int16  `json:"gunBulletMax"`
	GunBullet    int16  `json:"gunBullet"`
}

type S1BattleProtoGunFireAck struct {
	Result int32 `json:"result"`
}

type S1BattleProtoGunFireNotify struct {
	AttackerUid uint64  `json:"attackerUid"`
	GunType     byte    `json:"gunType"`
	GunSpeed    float32 `json:"gunSpeed"`
}

type S1BattleProtoGunFireReq struct {
	AttackerUid uint64  `json:"attackerUid"`
	GunType     byte    `json:"gunType"`
	GunSpeed    float32 `json:"gunSpeed"`
}

type S1BattleProtoGunHeatNotify struct {
	GunHeatMax float32 `json:"gunHeatMax"`
	GunHeat    float32 `json:"gunHeat"`
	PlayerUid  uint64  `json:"playerUid"`
}

type S1BattleProtoIsCanUseAtknotify struct {
	Teamcount int32   `json:"teamcount"`
	Teamid    []int32 `json:"teamid"`
	State     []int32 `json:"state"`
}

type S1BattleProtoIsCanUsePjnotify struct {
	Teamcount int32   `json:"teamcount"`
	Teamid    []int32 `json:"teamid"`
	State     []int32 `json:"state"`
}

type S1BattleProtoMarkerDetectAck struct {
	Result byte `json:"result"`
}

type S1BattleProtoMarkerDetectReq struct {
	Uid       uint64 `json:"uid"`
	Act       byte   `json:"act"`
	MarkerId  byte   `json:"markerId"`
	MarkerStr string `json:"markerStr"`
}

type S1BattleProtoModuleOffLineInfoNotify struct {
	Nouse byte `json:"nouse"`
}

type S1BattleProtoPlaceholderNotify struct {
	EventName string `json:"eventName"`
	State     byte   `json:"state"`
}

type S1BattleProtoPlayerCommandAutoControlNotify struct {
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandGunFireNotify struct {
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandLockScreenNotify struct {
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandMoveNotify struct {
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandSetDisconnectNotify struct {
	Uid   uint64 `json:"uid"`
	Value byte   `json:"value"`
}

type S1BattleProtoPlayerCommandUserCustomActionNotify struct {
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandUserInputNotify struct {
	Value byte `json:"value"`
}

type S1BattleProtoPlayerDeadNotify struct {
	Uid       uint64 `json:"uid"`
	KillerUid uint64 `json:"killerUid"`
	KillHonor int32  `json:"killHonor"`
	DeadType  int32  `json:"deadType"`
	Flag0     int32  `json:"flag0"`
	Flag1     int32  `json:"flag1"`
	Flag2     int32  `json:"flag2"`
}

type S1BattleProtoPlayerInitOkNotify struct {
	Uid uint64 `json:"uid"`
}

type S1BattleProtoPlayerReliveNotify struct {
	Uid uint64 `json:"uid"`
}

type S1BattleProtoPlayerResetNotify struct {
	PlayerUid uint64 `json:"playerUid"`
}

type S1BattleProtoPoJiaNotify struct {
	Uid      uint64  `json:"uid"`
	State    byte    `json:"state"`
	Timeleft float32 `json:"timeleft"`
	Timemax  float32 `json:"timemax"`
}

type S1BattleProtoProgressNotify struct {
	PlayerUid uint64  `json:"playerUid"`
	EventName string  `json:"eventName"`
	TimeMax   float32 `json:"timeMax"`
	TimeLeft  float32 `json:"timeLeft"`
	State     byte    `json:"state"`
}

type S1BattleProtoRfidnotify struct {
	S0Id  byte `json:"s0Id"`
	Flag  byte `json:"flag"`
	Type  byte `json:"type"`
	Area  byte `json:"area"`
	Data0 byte `json:"data0"`
	Data1 byte `json:"data1"`
	Data2 byte `json:"data2"`
	Data3 byte `json:"data3"`
	Data4 byte `json:"data4"`
	Data5 byte `json:"data5"`
}

type S1BattleProtoRfidreq struct {
	Type      byte   `json:"type"`
	Area      byte   `json:"area"`
	Data0     byte   `json:"data0"`
	Data1     byte   `json:"data1"`
	Data2     byte   `json:"data2"`
	Data3     byte   `json:"data3"`
	Data4     byte   `json:"data4"`
	Data5     byte   `json:"data5"`
	Timestamp uint64 `json:"timestamp"`
}

type S1BattleProtoReduceReliveTimeNotify struct {
	PlayerUid uint64  `json:"playerUid"`
	TimeMax   float32 `json:"timeMax"`
	TimeLeft  float32 `json:"timeLeft"`
}

type S1BattleProtoReloginNotify struct {
	SelfUid                         uint64      `json:"selfUid"`
	SelfTid                         uint32      `json:"selfTid"`
	SelfTeamId                      byte        `json:"selfTeamId"`
	SelfSeatIndex                   byte        `json:"selfSeatIndex"`
	SceneState                      int32       `json:"sceneState"`
	SceneStateTimeLeft              int32       `json:"sceneStateTimeLeft"`
	SceneTeamsCount                 int32       `json:"sceneTeamsCount"`
	SceneTeamsPlayerMax             int32       `json:"sceneTeamsPlayerMax"`
	ScenePlayersCount               int32       `json:"scenePlayersCount"`
	TeamsTotalDamage                []uint32    `json:"teamsTotalDamage"`
	PlayersTotalBuffCount           uint32      `json:"playersTotalBuffCount"`
	PlayersUid                      []uint64    `json:"playersUid"`
	PlayersTid                      []uint32    `json:"playersTid"`
	PlayersTeamId                   []byte      `json:"playersTeamId"`
	PlayersSeatIndex                []byte      `json:"playersSeatIndex"`
	PlayersState                    []byte      `json:"playersState"`
	PlayersCdLeftTime               []float32   `json:"playersCdLeftTime"`
	PlayersCdMaxTime                []float32   `json:"playersCdMaxTime"`
	PlayersHp                       []uint32    `json:"playersHp"`
	PlayersMaxHp                    []uint32    `json:"playersMaxHp"`
	PlayersHeat                     []float32   `json:"playersHeat"`
	PlayersMaxHeat                  []float32   `json:"playersMaxHeat"`
	PlayersBulletNum                []uint32    `json:"playersBulletNum"`
	PlayersMaxBulletNum             []uint32    `json:"playersMaxBulletNum"`
	PlayersArmor                    []int32     `json:"playersArmor"`
	PlayersMaxArmor                 []int32     `json:"playersMaxArmor"`
	PlayersDurability               []int32     `json:"playersDurability"`
	PlayersMaxDurability            []int32     `json:"playersMaxDurability"`
	PlayersBuffCount                []uint32    `json:"playersBuffCount"`
	PlayersBuffUid                  [][]uint64  `json:"playersBuffUid"`
	PlayersBuffTid                  [][]uint32  `json:"playersBuffTid"`
	PlayersBuffLevel                [][]uint32  `json:"playersBuffLevel"`
	PlayersBuffLeftTime             [][]float32 `json:"playersBuffLeftTime"`
	PlayersGunCtrlState             []byte      `json:"playersGunCtrlState"`
	PlayersMoveCtrlState            []byte      `json:"playersMoveCtrlState"`
	PlayersUserInputState           []byte      `json:"playersUserInputState"`
	PlayersLockScreenState          []byte      `json:"playersLockScreenState"`
	PlayersUserCustomActionState    []byte      `json:"playersUserCustomActionState"`
	PlayersAutoModeCtrlState        []byte      `json:"playersAutoModeCtrlState"`
	DevicesBaseStationState         []byte      `json:"devicesBaseStationState"`
	DevicesBaseStationStateLeftTime []float32   `json:"devicesBaseStationStateLeftTime"`
	DevicesRuneState                []byte      `json:"devicesRuneState"`
	DevicesRuneStateLeftTime        []float32   `json:"devicesRuneStateLeftTime"`
	WifiName                        string      `json:"wifiName"`
	WifiPassword                    string      `json:"wifiPassword"`
}

type S1BattleProtoSceneInfosNotify struct {
	PlayersCount     int32    `json:"playersCount"`
	PlayersUid       []uint64 `json:"playersUid"`
	PlayersTid       []uint32 `json:"playersTid"`
	PlayersTeamId    []byte   `json:"playersTeamId"`
	PlayersSeatIndex []byte   `json:"playersSeatIndex"`
}

type S1BattleProtoSceneStateNotify struct {
	State           ESceneStateType `json:"state"`
	TimeLeft        int32           `json:"timeLeft"`
	TeamCount       int32           `json:"teamCount"`
	TeamHurtSum     []int32         `json:"teamHurtSum"`
	TeamBekillCount []int32         `json:"teamBekillCount"`
}

type S1BattleProtoShareBulletReq struct {
	TeamId      byte `json:"teamId"`
	BulletIndex byte `json:"bulletIndex"`
}

type S1BattleProtoTeamInfoNotify struct {
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1BattleProtoUploadModuleInfo struct {
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1BattleProtoWarningNotify struct {
	PlayerUid    uint64 `json:"playerUid"`
	WarningLevel byte   `json:"warningLevel"`
}

type S1BattleProtoBaseArmorStateNotify struct {
	PlayerUid  uint64 `json:"playerUid"`
	ArmorId    int16  `json:"armorId"`
	ArmorCount int16  `json:"armorCount"`
	ArmorValue int16  `json:"armorValue"`
}

type S1BattleProtot2022RobotArmorCfgData struct {
	ArmorParaDataCount int32                                  `json:"armorParaDataCount"`
	ArmorParaDatas     []S1BattleProto2022ArmorParaConfigItem `json:"armorParaDatas"`
}

type S1ClientTransferRobotProto struct {
	DataLen int32  `json:"dataLen"`
	Data    []byte `json:"data"`
}

type S1ProtoClientExceptionInfo struct {
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1ProtoClientNetworkInfoNotify struct {
	Robotid         byte   `json:"robotid"`
	VtUplinkSpeed   uint32 `json:"vtUplinkSpeed"`
	VtDownlinkSpeed uint32 `json:"vtDownlinkSpeed"`
}

type S1ProtoClientSyncNotify struct {
	Uid               uint64   `json:"uid"`
	ConnClientState   int32    `json:"connClientState"`
	ConnRobotState    int32    `json:"connRobotState"`
	Battery           int32    `json:"battery"`
	SignalQuality     int32    `json:"signalQuality"`
	BattleMode        byte     `json:"battleMode"`
	OnlineModuleCount int32    `json:"onlineModuleCount"`
	OnlineModule      []uint64 `json:"onlineModule"`
}

type S1ProtoClientSyncReq struct {
	Uid               uint64   `json:"uid"`
	ConnClientState   int32    `json:"connClientState"`
	ConnRobotState    int32    `json:"connRobotState"`
	Battery           int32    `json:"battery"`
	SignalQuality     int32    `json:"signalQuality"`
	BattleMode        byte     `json:"battleMode"`
	OnlineModuleCount int32    `json:"onlineModuleCount"`
	OnlineModule      []uint64 `json:"onlineModule"`
}

type S1ProtoEnterRoomAck struct {
	ResultId EEnterRoomAckResultType `json:"resultId"`
	Uid      uint64                  `json:"uid"`
}

type S1ProtoEnterRoomNotify struct {
	Uid uint64 `json:"uid"`
}

type S1ProtoEnterRoomReq struct {
	Nouse1 byte `json:"nouse1"`
	Nouse2 byte `json:"nouse2"`
}

type S1ProtoGmcommandAck struct {
	ResultId int32 `json:"resultId"`
}

type S1ProtoGmcommandReq struct {
	Command string `json:"command"`
	Pars    string `json:"pars"`
}

type S1ProtoHeader struct {
	ProtoId    uint16 `json:"protoId"`
	ProtoType  uint16 `json:"protoType"`
	DataLen    uint32 `json:"dataLen"`
	SequenceId byte   `json:"sequenceId"`
	AckType    byte   `json:"ackType"`
	Nouse1     byte   `json:"nouse1"`
	Nouse2     byte   `json:"nouse2"`
}

type S1ProtoHeartBeatAck struct {
	ResultId   int32 `json:"resultId"`
	S0Clientid byte  `json:"s0Clientid"`
}

type S1ProtoHeartBeatReq struct {
	Nouse      byte `json:"nouse"`
	S0Clientid byte `json:"s0Clientid"`
}

type S1ProtoLeaveRoomAck struct {
	ResultId ELeaveRoomAckResultType `json:"resultId"`
}

type S1ProtoLeaveRoomNotify struct {
	Uid    uint64        `json:"uid"`
	Willgo ERoleLocation `json:"willgo"`
}

type S1ProtoLeaveRoomReq struct {
	Nouse byte `json:"nouse"`
}

type S1ProtoLoginAck struct {
	ResultId ELoginAckResultType `json:"resultId"`
	Uid      uint64              `json:"uid"`
	Token    string              `json:"token"`
}

type S1ProtoLoginReq struct {
	Account  string  `json:"account"`
	Password string  `json:"password"`
	Version  float32 `json:"version"`
	Tid      uint32  `json:"tid"`
	Teamid   uint32  `json:"teamid"`
	Hash     int64   `json:"hash"`
}

type S1ProtoLogoutAck struct {
	ResultId int32 `json:"resultId"`
}

type S1ProtoLogoutNotify struct {
	Uid    uint64 `json:"uid"`
	Reason byte   `json:"reason"`
}

type S1ProtoLogoutReq struct {
	Account string `json:"account"`
}

type S1ProtoMarkDetectResult struct {
	X        float32   `json:"x"`
	Y        float32   `json:"y"`
	W        float32   `json:"w"`
	H        float32   `json:"h"`
	Pitch    float32   `json:"pitch"`
	Yaw      float32   `json:"yaw"`
	Roll     float32   `json:"roll"`
	T44C2M   []float32 `json:"t44C2M"`
	Color    byte      `json:"color"`
	MarkerId uint16    `json:"markerId"`
	Distance uint16    `json:"distance"`
}

type S1ProtoMarkDetectResultAck struct {
	Result   int32   `json:"result"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	W        float32 `json:"w"`
	H        float32 `json:"h"`
	Color    byte    `json:"color"`
	MarkerId uint16  `json:"markerId"`
	Distance uint16  `json:"distance"`
}

type S1ProtoNetworkInfoNotify struct {
	Robotid           byte   `json:"robotid"`
	WifiUplinkSpeed   uint32 `json:"wifiUplinkSpeed"`
	WifiDownlinkSpeed uint32 `json:"wifiDownlinkSpeed"`
	VtUplinkSpeed     uint32 `json:"vtUplinkSpeed"`
	VtDownlinkSpeed   uint32 `json:"vtDownlinkSpeed"`
	WifiUplinkPlr     uint16 `json:"wifiUplinkPlr"`
	WifiDownlinkPlr   uint16 `json:"wifiDownlinkPlr"`
	VtUplinkPlr       uint16 `json:"vtUplinkPlr"`
	VtDownlinkPlr     uint16 `json:"vtDownlinkPlr"`
	OtherType         byte   `json:"otherType"`
	OtherPlr          uint16 `json:"otherPlr"`
	OtherDelay        uint16 `json:"otherDelay"`
	Delay             uint32 `json:"delay"`
}

type S1ProtoPingAck struct {
	ResultId int32 `json:"resultId"`
}

type S1ProtoPingReq struct {
	Nouse byte `json:"nouse"`
}

type S1ProtoPlayerCommandDisconnectNotify struct {
	Uid   uint64 `json:"uid"`
	Value byte   `json:"value"`
}

type S1ProtoPlayerCommandUserInputNotify struct {
	Value byte `json:"value"`
}

type S1ProtoReLoginAck struct {
	ResultId     ELoginAckResultType `json:"resultId"`
	Uid          uint64              `json:"uid"`
	Token        string              `json:"token"`
	WifiName     string              `json:"wifiName"`
	WifiPassword string              `json:"wifiPassword"`
}

type S1ProtoReLoginReq struct {
	Account  string  `json:"account"`
	Password string  `json:"password"`
	Version  float32 `json:"version"`
	Tid      uint32  `json:"tid"`
}

type S1ProtoReLoginRoomNotify struct {
	SelfUid            uint64   `json:"selfUid"`
	SelfTid            uint32   `json:"selfTid"`
	SelfTeamId         byte     `json:"selfTeamId"`
	SelfSeatIndex      byte     `json:"selfSeatIndex"`
	RoomState          int32    `json:"roomState"`
	RoomStateTimeLeft  int32    `json:"roomStateTimeLeft"`
	RoomSeatIslock     int32    `json:"roomSeatIslock"`
	RoomTeamsCount     int32    `json:"roomTeamsCount"`
	RoomTeamsPlayerMax int32    `json:"roomTeamsPlayerMax"`
	RoomPlayersCount   int32    `json:"roomPlayersCount"`
	PlayersUid         []uint64 `json:"playersUid"`
	PlayersTid         []uint32 `json:"playersTid"`
	PlayersTeamId      []byte   `json:"playersTeamId"`
	PlayersSeatIndex   []byte   `json:"playersSeatIndex"`
	WifiName           string   `json:"wifiName"`
	WifiPassword       string   `json:"wifiPassword"`
}

type S1ProtoRobotNetworkStatusAck struct {
	WifiDownlinkPlr uint16 `json:"wifiDownlinkPlr"`
	VtDownlinkPlr   uint16 `json:"vtDownlinkPlr"`
}

type S1ProtoRobotNetworkStatusReq struct {
	Nouse byte `json:"nouse"`
}

type S1ProtoRobotStudentSerialPortInfos struct {
	OtherType  byte   `json:"otherType"`
	OtherPlr   uint16 `json:"otherPlr"`
	OtherDelay uint16 `json:"otherDelay"`
}

type S1ProtoRoomInfosNotify struct {
	PlayersCount     int32    `json:"playersCount"`
	PlayersUid       []uint64 `json:"playersUid"`
	PlayersTid       []uint32 `json:"playersTid"`
	PlayersTeamId    []byte   `json:"playersTeamId"`
	PlayersSeatIndex []byte   `json:"playersSeatIndex"`
	WifiName         string   `json:"wifiName"`
	WifiPassword     string   `json:"wifiPassword"`
	CurrMatch        string   `json:"currMatch"`
}

type S1ProtoRoomPauseNotify struct {
	TimeElapse int32 `json:"timeElapse"`
}

type S1ProtoRoomStateNotify struct {
	State    ERoomStateType `json:"state"`
	TimeLeft int32          `json:"timeLeft"`
	IsPause  int32          `json:"isPause"`
}

type S1ProtoSetCurrMatchNotify struct {
	CurrMatch string `json:"currMatch"`
	CurrToken string `json:"currToken"`
}

type S1ProtoSetLockSeatNotify struct {
	LockSeat byte `json:"lockSeat"`
}

type S1ProtoSetReadyAck struct {
	ResultId int32 `json:"resultId"`
}

type S1ProtoSetReadyNotify struct {
	Uid   uint64 `json:"uid"`
	State byte   `json:"state"`
}

type S1ProtoSetReadyReq struct {
	State byte `json:"state"`
}

type S1ProtoSetSvrInfoAck struct {
	ResultId int32 `json:"resultId"`
}

type S1ProtoSetSvrInfoReq struct {
	SvrName string `json:"svrName"`
}

type S1ProtoSetTeamAck struct {
	ResultId  int32 `json:"resultId"`
	TeamId    byte  `json:"teamId"`
	SeatIndex byte  `json:"seatIndex"`
}

type S1ProtoSetTeamNotify struct {
	Uid       uint64 `json:"uid"`
	TeamId    byte   `json:"teamId"`
	SeatIndex byte   `json:"seatIndex"`
}

type S1ProtoSetTeamPlayerNumNotify struct {
	PlayerNum byte `json:"playerNum"`
}

type S1ProtoSetTeamReq struct {
	TeamId    byte `json:"teamId"`
	SeatIndex byte `json:"seatIndex"`
}

type S1ProtoSetTidAck struct {
	ResultId int32  `json:"resultId"`
	Tid      uint32 `json:"tid"`
}

type S1ProtoSetTidNotify struct {
	Uid uint64 `json:"uid"`
	Tid uint32 `json:"tid"`
}

type S1ProtoSetTidReq struct {
	Tid   uint32 `json:"tid"`
	Token string `json:"token"`
}

type S1ProtoSetTokenAck struct {
	ResultId int32 `json:"resultId"`
}

type S1ProtoSetTokenReq struct {
	Token string `json:"token"`
}

type S1ProtoSetWifiInfoNotify struct {
	WifiName     string `json:"wifiName"`
	WifiPassword string `json:"wifiPassword"`
}

type S1ProtoSvrStateAck struct {
	State     ERoomStateType `json:"state"`
	TimeLeft  int32          `json:"timeLeft"`
	IsPause   int32          `json:"isPause"`
	CurrMatch string         `json:"currMatch"`
}

type S1ProtoSvrStateReq struct {
	Nouse byte `json:"nouse"`
}

type S1ProtoTeamCampNotify struct {
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1ProtoTeamInfoNotify struct {
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1ProtoTeamLineupInfoNotify struct {
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1ProtoTeamLogoNotify struct {
	Teamid  int32  `json:"teamid"`
	DataLen int32  `json:"dataLen"`
	Data    string `json:"data"`
}

type S1ProtoTechnicalPauseInfoNotify struct {
	PauseTimeType           byte   `json:"pauseTimeType"`
	PauseSide               byte   `json:"pauseSide"`
	RedShortPauseLeftCount  byte   `json:"redShortPauseLeftCount"`
	RedLongPauseLeftCount   byte   `json:"redLongPauseLeftCount"`
	BlueShortPauseLeftCount byte   `json:"blueShortPauseLeftCount"`
	BlueLongPauseLeftCount  byte   `json:"blueLongPauseLeftCount"`
	PausedDuration          uint32 `json:"pausedDuration"`
}

type S1ProtoTestAck struct {
	ResultId int32  `json:"resultId"`
	Test     string `json:"test"`
}

type S1ProtoTestReq struct {
	Test string `json:"test"`
}

type S1ProtoUserStateAck struct {
	Online int32 `json:"online"`
}

type S1ProtoUserStateReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
