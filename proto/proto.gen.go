package proto

type S1BattleProto2022AirplaneStatusNotify struct {
	Robotid       byte
	Energy        int32
	IsFire        byte
	Curbullet     int16
	Maxbullet     int16
	Curshoottime  float32
	Fixshoottime  float32
	Cdleft        float32
	Countleft     uint16
	Cdrefreshcost int32
	Isincd        byte
}

type S1BattleProto2022ArmorParaConfigItem struct {
	TiggerPress    float32
	BulletMaxPress float32
	GolfMinPress   float32
}

type S1BattleProto2022BaseLightingEffect struct {
	LightColor uint32
}

type S1BattleProto2022BaseShellStatus struct {
	BaseStatus                byte
	BaseMotoOneStatus         byte
	BaseMotoOneBlockStatus    byte
	BaseMotoOneOnlineStatus   byte
	BaseMotoTwoStatus         byte
	BaseMotoTwoBlockStatus    byte
	BaseMotoTwoOnlineStatus   byte
	BaseMotoThreeStatus       byte
	BaseMotoThreeBlockStatus  byte
	BaseMotoThreeOnlineStatus byte
	Reserved0                 byte
	Reserved2                 byte
}

type S1BattleProto2022ControlCmdT struct {
	OptCode uint16
	Data    uint16
}

type S1BattleProto2022ClientAirplaneCtrlReq struct {
	Robotid     byte
	ControlCode byte
}

type S1BattleProto2022ClientAllClientStatusSync struct {
	Listlength byte
	Status     []byte
}

type S1BattleProto2022ClientArmorLifeInfo struct {
	ModuleId  int16
	LifeState byte
}

type S1BattleProto2022ClientBaseRobotDevStatus struct {
	ShellStatus                    byte
	Online                         byte
	DartTargetPosition             int16
	DartTargetIsInplace            byte
	BaseShellOneSensorStatus       byte
	BaseShellTwoSensorStatus       byte
	BaseShellThreeSensorStatus     byte
	BaseDartboardSensorStatus      byte
	BaseShellMotoOneBlockStatus    byte
	BaseShellMotoOneOnlineStatus   byte
	BaseShellMotoTwoBlockStatus    byte
	BaseShellMotoTwoOnlineStatus   byte
	BaseShellMotoThreeBlockStatus  byte
	BaseShellMotoThreeOnlineStatus byte
	BaseDartboardMotoBlockStatus   byte
	BaseDartboardMotoOnlineStatus  byte
	BaseShellSelfcheckStatus       byte
	BaseDartboardSelfcheckStatus   byte
	BaseDartboardMoveTimeout       byte
	BaseShellOpenTimeout           byte
	BaseShellCloseTimeout          byte
	DartBoardIronline              byte
	DartBoardBrokenErr             byte
	IrledStatusLeft                uint32
	IrledStatusRight               uint32
	BaseShellMotoOneOverHeat       byte
	BaseShellMotoTwoOverHeat       byte
	BaseShellMotoThreeOverHeat     byte
	BaseDartBoardMotoOverHeat      byte
	BaseShellMotoOneOverCurrent    byte
	BaseShellMotoTwoOverCurrent    byte
	BaseShellMotoThreeOverCurrent  byte
	BaseDartBoardMotoOverCurrent   byte
	BaseShellMotoOneOverSpeed      byte
	BaseShellMotoTwoOverSpeed      byte
	BaseShellMotoThreeOverSpeed    byte
	BaseDartBoardMotoOverSpeed     byte
}

type S1BattleProto2022ClientBaseRobotDevStatusSyncData struct {
	BaseRobotDevStatusListLen byte
	BaseRobotDevStatusList    []S1BattleProto2022ClientBaseRobotDevStatus
}

type S1BattleProto2022ClientBattleFirstData struct {
	Progress byte
	IsPaused byte
}

type S1BattleProto2022ClientBuffItem struct {
	BuffId byte
	CdTime float32
}

type S1BattleProto2022ClientBuffSlot struct {
	BuffItemsLen int32
	BuffItems    []S1BattleProto2022ClientBuffItem
}

type S1BattleProto2022ClientBuyBulletAck struct {
	Result int32
}

type S1BattleProto2022ClientBuyBulletReq struct {
	Type  byte
	Count byte
}

type S1BattleProto2022ClientCheckInTimeStampNotify struct {
	RobotidListLen byte
	RobotidList    []uint32
}

type S1BattleProto2022ClientCurrentProgress struct {
	CurProgress int32
	IsPaused    byte
}

type S1BattleProto2022ClientCustomControlData struct {
	RobotId     int32
	DataLen     int32
	DataListLen byte
	DataList    []byte
}

type S1BattleProto2022ClientEconomyNotify struct {
	Clientid byte
	Type     int32
	Money    int32
}

type S1BattleProto2022ClientExceptionCapDataNotify struct {
	RobotidListLen byte
	RobotidList    []byte
}

type S1BattleProto2022ClientExceptionRecoverCardNotify struct {
	Teamid    byte
	Robotid   byte
	Exception byte
}

type S1BattleProto2022ClientExchangeProStateNotify struct {
	RobotId          byte
	IsOnline         byte
	ExchangeLevel    byte
	IrState          byte
	Ir1              byte
	Ir2              byte
	X                float32
	Y                float32
	Z                float32
	Pitch            float32
	Roll             float32
	Yaw              float32
	ErrorCode        byte
	GoldCount        byte
	SilverCount      byte
	State            byte
	MineralRfidState byte
	CoinsExchange    int32
	CoinsAddRate     float32
	EngineerLevel    byte
}

type S1BattleProto2022ClientFlycatStatusSync struct {
	TeamId                     byte
	Online                     byte
	SysState                   byte
	CtrlState                  byte
	WorkState                  byte
	Battery                    byte
	MotorOneOnlineError        byte
	MotorTwoOnlineError        byte
	MotorOneOverheatError      byte
	MotorTwoOverheatError      byte
	LowBatteryWarning          byte
	LowBatteryWarningThreshold byte
	BatteryRenewThreshold      byte
}

type S1BattleProto2022ClientGmcommand struct {
	Len byte
	Cmd string
}

type S1BattleProto2022ClientHitNotify struct {
	RobotId     byte
	OnHitType   byte
	ArmorNumber byte
	HpReduce    int32
	HpCurr      int32
	HpMax       int32
}

type S1BattleProto2022ClientHoldCenterPointCompleteNotify struct {
	Teamid    int32
	RewardExp float32
}

type S1BattleProto2022ClientHpChangeDetailVal struct {
	ByBullet        int32
	ByGolf          int32
	ByMissile       int32
	ByOverSpeed     int32
	ByoverFreq      int32
	ByOverPower     int32
	ByOverHeat      int32
	ByModuleOffline int32
	ByPunish        int32
	ByKill          int32
	ByCrash         int32
	ByWifiOffline   int32
	All             int32
}

type S1BattleProto2022ClientKickOffRobotNotify struct {
	RobotId int32
	Reason  byte
}

type S1BattleProto2022ClientMineralExchangeNotify struct {
	TeamId      byte
	MineralType byte
	Value       int32
}

type S1BattleProto2022ClientMineralInfoSync struct {
	IsRedConnect             byte
	IsBlueConnect            byte
	RedGoldCount             byte
	RedSilverCount           byte
	BlueGoldCount            byte
	BlueSilverCount          byte
	RedInfraredStateListLen  byte
	RedInfraredState         []byte
	BlueInfraredStateListLen byte
	BlueInfraredState        []byte
}

type S1BattleProto2022ClientOutpostDeviceStatusSync struct {
	RobotId               byte
	Online                byte
	StatusCode            byte
	CurTemperature        int32
	TemperatureWarning    byte
	SpinSpeedRatio        float32
	MotorOnline           byte
	MotorCurrentWarning   byte
	MotorSpeedoverWarning byte
	IsMagnetOn            byte
	ActionTimeout         byte
	DartBoardIrOnline     byte
	DartBoardBrokenErr    byte
	IrledStatusLeft       uint32
	IrledStatusRight      uint32
}

type S1BattleProto2022ClientRaderInfoToClient struct {
	TargetRobotId uint16
	TargetPosX    float32
	TargetPosY    float32
	TorwardAngle  float32
}

type S1BattleProto2022ClientRecoverCardStatus struct {
	Redcard  byte
	Bluecard byte
}

type S1BattleProto2022ClientResultDataInfo struct {
	WebGameId                    int32
	GameId                       int64
	GameOverReasonId             byte
	GameOverReasonListLen        int32
	GameOverReason               []byte
	StartTime                    int64
	DurationTime                 int32
	GameOrder                    byte
	TotalRound                   byte
	RecordsListLen               int32
	Records                      []int32
	Winner                       byte
	ScoreRed                     int32
	ScoreBlue                    int32
	RedCurrentHp                 int32
	BlueCurrentHp                int32
	RedTotalHp                   int32
	BlueTotalHp                  int32
	RedWarning1                  int32
	RedWarning2                  int32
	RedWarning3                  int32
	BlueWarning1                 int32
	BlueWarning2                 int32
	BlueWarning3                 int32
	RedHits                      int32
	BlueHits                     int32
	RedRuneCount                 int32
	BlueRuneCount                int32
	RedKillCount                 int32
	BlueKillCount                int32
	RedDartHitCount              int32
	BlueDartHitCount             int32
	RedShootSmallCount           int32
	RedShootBigCount             int32
	BlueShootSmallCount          int32
	BlueShootBigCount            int32
	RedGuardLeftLives            byte
	BlueGuardLeftLives           byte
	GuardFixedLives              byte
	RobotsNum                    int32
	RobotsInfos                  []S1BattleProto2022ClientRobotDataInfo
	RedSnipeSuccCount            byte
	BlueSnipeSuccCount           byte
	RedUavCallCount              byte
	BlueUavCallCount             byte
	RedRadarBuffDoubleUsedCount  byte
	BlueRadarBuffDoubleUsedCount byte
}

type S1BattleProto2022ClientRobotArmorLifeQueryResult struct {
	IsWindMill         byte
	WindMillTeamId     byte
	RobotId            byte
	RobotTotalArmorNum int32
	InfosLen           byte
	LifeInfos          []S1BattleProto2022ClientArmorLifeInfo
}

type S1BattleProto2022ClientRobotBaseDataSync struct {
	UserId                      uint16
	Id                          byte
	Type                        byte
	Level                       byte
	CpuId                       uint32
	FixHp                       int32
	FixPower                    float32
	FixPowerBufferRecoverable   float32
	FixPowerBufferUnrecoverable float32
	FixSmallSpeed               float32
	FixSmallFreq                float32
	FixSmallHeatEnergy          float32
	FixSmallHeatEnergyCoolRate  float32
	FixSmallSpeed2              float32
	FixSmallFreq2               float32
	FixSmallHeatEnergy2         float32
	FixSmallHeatEnergyCoolRate2 float32
	FixBigSpeed                 float32
	FixBigFreq                  float32
	FixBigHeatEnergy            float32
	FixBigHeatEnergyCoolRate    float32
	ExpOffer                    float32
	CurExpAddRate               float32
	CurRebornCd                 int32
	FixExpUpgradeNeed           float32
	FixExpAddRate               float32
	FixRebornCd                 int32
}

type S1BattleProto2022ClientRobotDataInfo struct {
	RobotId     int32
	RobotType   int32
	CurHp       int32
	MaxHp       int32
	IsConnected byte
	DeadReason  int32
	BehitedVal  int32
	TotalAttack int32
}

type S1BattleProto2022ClientRobotDeathNotify struct {
	RobotidDeath  byte
	RobotidKiller byte
	DeathReason   byte
	BFirstBlood   byte
	KillCount     int32
}

type S1BattleProto2022ClientRobotFirstData struct {
	Userid uint16
}

type S1BattleProto2022ClientRobotMapData struct {
	Joinstate byte
	IsAlive   byte
	X         float32
	Y         float32
	Yaw       float32
}

type S1BattleProto2022ClientRobotStatusNotify struct {
	RobotUserId         int32
	ConnState           byte
	JoinState           byte
	SurvivalState       byte
	DeathSubState       int32
	WifiWeak            byte
	ModuleOnline        byte
	BatteryLow          byte
	IdConflict          byte
	IsCanReliveByClient byte
}

type S1BattleProto2022ClientRobotSyncData struct {
	UserId                      uint16
	CurHp                       uint16
	Voltage                     float32
	Current                     float32
	Rssi                        byte
	CurPower                    float32
	YellowWarningCount          byte
	PowerBuffer                 float32
	CurSmallSpeed1              float32
	CurSmallFreq1               float32
	CurSmallHeatEnergy1         float32
	CurSmallHeatEnergyCoolRate1 float32
	SmallShootNum1              int32
	SmallLeftBulletsNum1        int32
	CurSmallSpeed2              float32
	CurSmallFreq2               float32
	CurSmallHeatEnergy2         float32
	CurSmallHeatEnergyCoolRate2 float32
	SmallShootNum2              int32
	SmallLeftBulletsNum2        int32
	CurBigSpeed                 float32
	CurBigFreq                  float32
	CurBigHeatEnergy            float32
	CurBigHeatEnergyCoolRate    float32
	BigBulletsShootCnt          int32
	BigLeftBulletCnt            int32
	CurExp                      float32
	ByAttackTotal               int32
	MurderId                    byte
	Power                       byte
	RfidNum                     byte
	Rfid0                       byte
	Rfid1                       byte
	BloodNum                    byte
	Blood0                      byte
	Blood1                      byte
	Blood2                      byte
	ShooterDetectNum            byte
	ShooterDetect0              byte
	ShooterDetect1              byte
	ShooterDetect2              byte
	HasshooterDetect0           byte
	HasshooterDetect1           byte
	HasshooterDetect2           byte
	UwbNum                      byte
	Uwb                         byte
	StrikeNum                   byte
	Strike0                     byte
	Strike1                     byte
	Strike2                     byte
	Strike3                     byte
	Strike4                     byte
	Strike5                     byte
	Strike6                     byte
	Strike7                     byte
	Strike8                     byte
	Strike9                     byte
	CameraNum                   byte
	Camera                      byte
	CapNum                      byte
	Cap                         byte
	RebornRfid                  byte
	X                           float32
	Y                           float32
	Z                           float32
	Compass                     float32
	CoustomData1                float32
	CoustomData2                float32
	CoustomData3                float32
	Mask                        byte
	RtsDmgData                  S1BattleProto2022ClientHpChangeDetailVal
	CurrentPerformancePoint     int32
	HpLevel                     int32
	ChassisPowerLevel           int32
	ShootSpeedLevel             int32
}

type S1BattleProto2022ClientShielddata struct {
	RedHasShield  byte
	BlueHasShield byte
	RedShield     int32
	BlueShield    int32
	ShieldMax     int32
}

type S1BattleProto2022ClientSafetySync struct {
	Id            byte
	Lift1Connect  byte
	Lift2Connect  byte
	FlycatConnect byte
	Lift1State    byte
	Lift2State    byte
	FlycatState   byte
	Lift1Error    byte
	Lift2Error    byte
	CatError      byte
	FlycatBattery byte
}

type S1BattleProto2022ClientServerMapSync struct {
	YawOffset      float32
	AnchorMask     byte
	RobotidListLen byte
	RobotidList    []S1BattleProto2022ClientRobotMapData
}

type S1BattleProto2022ClientServerUimessage struct {
	Msg string
}

type S1BattleProto2022ClientSeverAlertNotify struct {
	YellowRateWarningCount          int32
	RedRateWarningCount             int32
	LastYellowTimestamp             uint32
	LastRedTimestamp                uint32
	LastYellowGameLeftTime          int32
	LastRedGameLeftTime             int32
	CurrAlertLevel                  byte
	GetQingflowData                 byte
	BalancedInfantryUltraLimitError byte
}

type S1BattleProto2022ClientShowMessage struct {
	Teamid    byte
	MsgType   byte
	MsgEnum   uint32
	MsgParams string
}

type S1BattleProto2022ClientSiloEnvCtrReq struct {
	Teamid      byte
	ControlCode byte
	Target      byte
}

type S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify struct {
	TeamId      byte
	DoorState   byte
	DoorOpenCnt byte
}

type S1BattleProto2022ClientStatusSync struct {
	ClientId int32
	Status   int32
}

type S1BattleProto2022ClientSupplySync struct {
	SupplyId            int32
	Connect             int32
	State               int32
	UsableBullets       int32
	UsedBullets         int32
	ReadyBox            int32
	FreedBox            int32
	Timespan            float32
	NextaddBullets      int32
	MakerRobotid        int32
	MakeBullets         int32
	ErrorCodeArrListLen int32
	ErrorCodeArrList    []int32
	RfidRobotIdListLen  int32
	RfidRobotIdList     []int32
}

type S1BattleProto2022ClientTalentDataAck struct {
	Result int32
}

type S1BattleProto2022ClientTalentDataNotify struct {
	Robotid             byte
	Data                []性能体系_数据
	IsBalance           byte
	IsSemiAutomaticCtrl byte
	IsTempManualCtrl    byte
}

type S1BattleProto2022ClientTalentDataReq struct {
	Robotid byte
	目标类型    string
	性能类型    string
	是否设置副武器 byte
}

type S1BattleProto2022ClientTeamInfo struct {
	TotalHp                 uint32
	CurrentHp               uint32
	Warning1Count           uint16
	Warning2Count           uint16
	Warning3Count           uint16
	TotalHurt               uint16
	BigBulletAvailableNum   int16
	BigBulletQuotaNum       int16
	BigBulletCanBuyNum      int16
	SmallBulletAvailableNum int16
	SmallBulletQuotaNum     int16
	SmallBulletCanBuyNum    int16
	AirSupportNum           uint16
	BloodPack               uint32
	CenterPointEnergy       uint32
	TeamAvaliableCoinsNum   int32
	TotalCoins              int32
}

type S1BattleProto2022ClientVtmSpeedSync struct {
	Clientid             int32
	CurrentFreq          int32
	TxConnect            int32
	CurrentSpeedRate     int32
	CurrentTxTemperature int32
}

type S1BattleProto2022ClientWarningNotify struct {
	Type            byte
	Team            byte
	RobotId         byte
	Leftcredit      byte
	HpChangePercent byte
	MaskTimeSelf    byte
	MaskTimeOthers  byte
}

type S1BattleProto2022ClientWeaponFireNotify struct {
	Robotid    byte
	BulletType byte
	Speed      float32
	Angle      float32
}

type S1BattleProto2022ClientsFirstSync struct {
	GameMode        byte
	BattleFirstData S1BattleProto2022ClientBattleFirstData
	RobotsNum       int32
	RobotsFirstData []S1BattleProto2022ClientRobotFirstData
	ClientNum       int32
	ClientsStatus   []byte
	RobotsNum1      int32
	RobotStatus     []S1BattleProto2022ClientRobotStatusNotify
	CourtYawOffset  float32
}

type S1BattleProto2022ClientsRobotModuleErrorInfo struct {
	RobotId               int32
	PowerNum              byte
	Power                 byte
	RfidNum               byte
	Rfid0                 byte
	Rfid1                 byte
	BloodNum              byte
	Blood0                byte
	Blood1                byte
	Blood2                byte
	IshasExShooter        byte
	SamllShooterDetectNum byte
	SamllshooterDetect0   byte
	SamllshooterDetect1   byte
	BigShooterDetectNum   byte
	BigshooterDetect      byte
	UwbNum                byte
	Uwb                   byte
	StrikeNum             byte
	Strike0               byte
	Strike1               byte
	Strike2               byte
	Strike3               byte
	Strike4               byte
	Strike5               byte
	Strike6               byte
	Strike7               byte
	Strike8               byte
	Strike9               byte
	CameraNum             byte
	Camera                byte
	CapNum                byte
	Cap                   byte
	Spower                byte
	Srfid0                byte
	Srfid1                byte
	Scamera               byte
	Sblood0               byte
	Sblood1               byte
	SshooterDetect0       byte
	SshooterDetect1       byte
	SshooterDetect2       byte
	Suwb                  byte
	Sarmor0               byte
	Sarmor1               byte
	Sarmor2               byte
	Sarmor3               byte
	Sarmor4               byte
	Sarmor5               byte
	Sarmor6               byte
	Sarmor7               byte
	Sarmor8               byte
	Sarmor9               byte
	Scap                  byte
	SmainController       byte
}

type S1BattleProto2022ClientsServerBaseSync struct {
	RobotsSyncDatasLen int32
	RobotsBaseSyncData []S1BattleProto2022ClientRobotBaseDataSync
}

type S1BattleProto2022ClientsServerSync struct {
	PassTime            float32
	LeftTime            float32
	CenterPointCoolTime float32
	TeamInfosLen        int32
	TeamInfos           []S1BattleProto2022ClientTeamInfo
	RobotsSyncDatasLen  int32
	RobotsSyncDatas     []S1BattleProto2022ClientRobotSyncData
}

type S1BattleProto2022ConfigTabAck struct {
	Magic              uint32
	RobotConfigVersion byte
	Color              byte
	ModuleNum          S1BattleProto2022ModuleNum
	HealthCalc         S1BattleProto2022HealthCalc
	GameLimit          S1BattleProto2022GameLimit
	ArmorData          S1BattleProtot2022RobotArmorCfgData
}

type S1BattleProto2022Energy2019AmorSelfCheck struct {
	Nouse byte
}

type S1BattleProto2022Energy2019ArmorHitUpload struct {
	Row     byte
	ArmorId byte
	HitType uint16
}

type S1BattleProto2022Energy2019SetArmLight struct {
	ArmorColor []uint32
	ArmColor   []uint32
	Effect     []byte
	ExtVar     []uint16
	Row        byte
}

type S1BattleProto2022Energy2019StateReport struct {
	State       byte
	Armors      []byte
	Reserve     byte
	RotateDeg   uint16
	RotateSpeed float32
}

type S1BattleProto2022Energy2020SetArmRotate struct {
	CtrlMode byte
	Direct   byte
}

type S1BattleProto2022EnergyReset struct {
	Res byte
}

type S1BattleProto2022EnergySetId struct {
	Res byte
}

type S1BattleProto2022EnergyStateChangeNotify struct {
	RuneId      byte
	State       byte
	RingSum     byte
	AtkBuffVal  uint16
	DefBuffVal  byte
	BeforeState byte
}

type S1BattleProto2022EnergyStateSync struct {
	RuneId            byte
	Connect           byte
	State             byte
	Round             byte
	Time              byte
	Error             byte
	Armor0            byte
	Armor1            byte
	Armor2            byte
	Armor3            byte
	Armor4            byte
	Armor5            byte
	Motor             byte
	AvailbleCountdown byte
}

type S1BattleProto2022EnerySetLogoLight struct {
	IconRColor uint32
	Row        byte
}

type S1BattleProto2022EnvBaseShellControl struct {
	BaseShellControl byte
	Rgb              uint16
}

type S1BattleProto2022EnvDevicesDescripeAck struct {
	Name S1BattleProto2022RobotDynamicUistr
}

type S1BattleProto2022EnvDevicesDescripeReq struct {
	Res byte
}

type S1BattleProto2022EnvHeartBeatAck struct {
	Time uint32
}

type S1BattleProto2022EnvHeartBeatReq struct {
	Status byte
}

type S1BattleProto2022ExchangeProClearOreRes struct {
	Res byte
}

type S1BattleProto2022ExchangeProCtrCmd struct {
	Command    byte
	Reserved1  byte
	Reserved2  byte
	Seq        uint32
	ModuleType byte
	ModuleId   byte
}

type S1BattleProto2022ExchangeProCtrCmdack struct {
	ErrorCode  byte
	Seq        uint32
	ModuleType byte
	ModuleId   byte
}

type S1BattleProto2022ExchangeProLidarInfo struct {
	State      byte
	ModuleType byte
	ModuleId   byte
}

type S1BattleProto2022ExchangeProPosition struct {
	X          int32
	Y          int32
	Z          int32
	Pitch      int32
	Roll       int32
	Yaw        int32
	Seq        uint32
	ModuleType byte
	ModuleId   byte
}

type S1BattleProto2022ExchangeProRealPosition struct {
	X          int32
	Y          int32
	Z          int32
	Pitch      int32
	Roll       int32
	Yaw        int32
	Seq        uint32
	ModuleType byte
	ModuleId   byte
}

type S1BattleProto2022FlyCatStatus struct {
	SysState    byte
	CtrlState   byte
	WorkState   byte
	Battery     byte
	SensorState byte
}

type S1BattleProto2022GameLimit struct {
	PowerThreshold        uint16
	PowerBuffer1          uint16
	PowerBuffer2          uint16
	PowerMaxhurt          uint16
	PowerHurtTabCount     int32
	PowerHurtTabs         []byte
	ShooterLimitsCount    int32
	ShooterLimits         []S1BattleProto2022ShooterLimit
	HeatLimitsCount       int32
	HeatLimits            []S1BattleProto2022HeatLimit
	ShooterFreqLimitCount int32
	FreqLimits            []S1BattleProto2022ShooterFreqLimit
}

type S1BattleProto2022GripperStateNotify struct {
	IsConnect             byte
	IsHasMineralsListLen  byte
	IsHasMineralsList     []byte
	ErrorsListLen         byte
	ErrorsList            []byte
	GrippersStatesListLen byte
	GripperStates         []byte
}

type S1BattleProto2022HealthCalc struct {
	HealthPointStart uint16
	HealthPointFull  uint16
	BulletHurt       uint16
	GolfHurt         uint16
	ImpactHurt       uint16
}

type S1BattleProto2022HeatLimit struct {
	HeatSpdThreshold uint16
	HeatSpdMaxHurt   uint16
	HeatHurtTabCount int32
	HeatHurtTab      []byte
	HeatCdFreq       byte
	HeatCdStep       uint16
}

type S1BattleProto2022IoctrCfgSet struct {
	IsDp2SpiMode    byte
	IsPwmDp2Io2     byte
	IsPwmDp2Io5     byte
	IsPwmDp2Io6     byte
	IsPwmDp2Io7     byte
	Dp2PwmFrq       uint16
	IsOutputDp1     byte
	IsTriggerDp1Io0 byte
	IsTriggerDp1Io1 byte
	IsTriggerDp1Io2 byte
	IsTriggerDp1Io3 byte
	IsTriggerDp1Io4 byte
	IsTriggerDp1Io5 byte
	IsTriggerDp1Io6 byte
	IsTriggerDp1Io7 byte
	IsTriggerDp3Io0 byte
	IsTriggerDp3Io1 byte
	IsTriggerDp3Io2 byte
	IsTriggerDp3Io3 byte
	IsTriggerDp3Io4 byte
	IsTriggerDp3Io5 byte
	IsTriggerDp3Io6 byte
	IsTriggerDp3Io7 byte
	ModuleId        byte
	ModuleType      byte
}

type S1BattleProto2022IoctrCfgSetAck struct {
	ErrCode    byte
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetActuator struct {
	Actuator0  byte
	Actuator1  byte
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetActuatorAck struct {
	Actuator0  byte
	Actuator1  byte
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetVal struct {
	Dp2OutputValsLen int32
	Dp2OutputVal     []uint16
	OutDp1Io0        byte
	OutDp1Io1        byte
	OutDp1Io2        byte
	OutDp1Io3        byte
	OutDp1Io4        byte
	OutDp1Io5        byte
	OutDp1Io6        byte
	OutDp1Io7        byte
	ModuleId         byte
	ModuleType       byte
}

type S1BattleProto2022IoctrSetValAck struct {
	ErrCode    byte
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrState struct {
	Dp1Gpio0          byte
	Dp1Gpio1          byte
	Dp1Gpio2          byte
	Dp1Gpio3          byte
	Dp1Gpio4          byte
	Dp1Gpio5          byte
	Dp1Gpio6          byte
	Dp1Gpio7          byte
	Dp3Adc1In0        byte
	Dp3Adc1In1        byte
	Dp3Adc1In2        byte
	Dp3Adc1In3        byte
	Dp3Adc1In4        byte
	Dp3Adc1In5        byte
	Dp3Adc1In6        byte
	Dp3Adc1In7        byte
	Dp3AdcValsListLen int32
	Dp3AdcVals        []uint16
	IsPwmDp2Io0       byte
	IsPwmDp2Io1       byte
	IsPwmDp2Io2       byte
	IsPwmDp2Io3       byte
	IsPwmDp2Io4       byte
	IsPwmDp2Io5       byte
	IsPwmDp2Io6       byte
	IsPwmDp2Io7       byte
	Dp2OutValsListLen int32
	Dp2OutVars        []uint16
	Dp2PwmFrq         uint16
	Error             byte
	ModuleId          byte
	ModuleType        byte
}

type S1BattleProto2022IoctrTriggerVal struct {
	Dp1InputValBefore byte
	Dp3InputValBefore byte
	Dp1InputValAfter  byte
	Dp3InputValAfter  byte
	SysTickMs         uint32
	ModuleId          byte
	ModuleType        byte
}

type S1BattleProto2022MapClickInfoNotify struct {
	TeamId         byte
	IsSendAll      byte
	RobotidListLen byte
	RobotidList    []byte
	Mode           byte
	EnemyId        byte
	Ascii          byte
	Type           byte
	ScreenX        uint16
	ScreenY        uint16
	MapX           float32
	MapY           float32
}

type S1BattleProto2022ModuleNum struct {
	SmallStrikeNum byte
	BigStrikeNum   byte
	LightBarNum    byte
	PowerNum       byte
	SmallShootNum  byte
	BigShootNum    byte
	CameraNum      byte
	RfidNum        byte
	UwbNum         byte
	SmallArmorImp  byte
	BigArmorImp    byte
	LightBrdImp    byte
	PowerBrdImp    byte
	SmallShotImp   byte
	BigShotImp     byte
	CameraImp      byte
	RfidImp        byte
	UwbImp         byte
	CapNum         byte
	CapImp         byte
	ExtTypeAnum    byte
	ExtTypeAimp    byte
	ExtTypeBnum    byte
	ExtTypeBimp    byte
	ExtTypeCnum    byte
	ExtTypeCimp    byte
	Res            byte
}

type S1BattleProto2022PowerSwitchState struct {
	Chassis byte
	Gimbal  byte
	Shooter byte
}

type S1BattleProto2022QueryRobotInfo struct {
	ModuleType byte
	ModuleId   byte
}

type S1BattleProto2022QueryRobotInfoResult struct {
	LoaderVersion   uint32
	AppVersion      uint32
	DeviceIdListLen int32
	DeviceIdList    []int32
	Reserved        uint32
	ModuleType      byte
	ModuleId        byte
}

type S1BattleProto2022ReLoginDataSync struct {
	RoleS0Id     uint64
	BuffCount    byte
	BuffUids     []uint64
	BuffTids     []uint32
	BuffLevels   []uint32
	BuffVisibles []byte
	BuffMaxTime  []float32
	BuffLeftTime []float32
}

type S1BattleProto2022RobotBaseDevCtlCmdAck struct {
	Placeholder byte
}

type S1BattleProto2022RobotCheckinTimestamp struct {
	Magic     uint32
	DataVer   byte
	Timestamp uint32
}

type S1BattleProto2022RobotDataSync struct {
	CurHp                 int32
	MaxHp                 int32
	Level                 byte
	LightEffectMask       uint32
	ShooterHeatsCount     int32
	ShooterHeats          []float32
	ShooterHeatThresCount int32
	ShooterHeatThres      []float32
	CurHeatCoolCount      int32
	CurHeatCool           []float32
	IsGuardAlive          uint16
	InventedShieldValue   uint16
	FixPower              uint16
	FixSamllShootSpeed1   uint16
	FixBigShootSpeed      uint16
	FixSamllShootSpeed2   uint16
}

type S1BattleProto2022RobotDynamicUistr struct {
	EnvDeviceXnameEn string
	EnvDeviceXnameZh string
	EnvDeviceYnameEn string
	EnvDeviceYnameZh string
	EnvDeviceZnameEn string
	EnvDeviceZnameZh string
	ExtModuleAnameEn string
	ExtModuleAnameZh string
	ExtModuleBnameEn string
	ExtModuleBnameZh string
	ExtModuleCnameEn string
	ExtModuleCnameZh string
}

type S1BattleProto2022RobotGetCheckCapAck struct {
	CheckedCap       float32
	VoltageFirst     float32
	VoltageFinal     float32
	MeasureTime      float32
	OutputQEnergy    float32
	RecharingQEnergy float32
	Date             uint32
	Time             uint32
}

type S1BattleProto2022RobotGetCheckCapReq struct {
	Res byte
}

type S1BattleProto2022RobotHited struct {
	OnHitType   byte
	AttackSpeed uint16
	ArmorNumber byte
	Press       float32
	TimeStamp   uint64
}

type S1BattleProto2022RobotIrcheckReq struct {
	Check byte
}

type S1BattleProto2022RobotInitCfgTab struct {
	Ext17MmSpeed byte
	ResCount     int32
	Res          []byte
}

type S1BattleProto2022RobotMeasureStartReq struct {
	MaxPrriod  uint32
	PushPrriod uint32
}

type S1BattleProto2022RobotMeasureStartRsp struct {
	Type   byte
	Id     byte
	Status byte
}

type S1BattleProto2022RobotMeasureStopReq struct {
	Res byte
}

type S1BattleProto2022RobotMeasureStopRsp struct {
	Res byte
}

type S1BattleProto2022RobotMoudleLife struct {
	ModuleId      byte
	ModuleType    byte
	AckModuleId   byte
	AckModuleType byte
	Status        string
}

type S1BattleProto2022RobotNewHeartBeatAck struct {
	Nouse byte
}

type S1BattleProto2022RobotNewHeartBeatReq struct {
	Nouse byte
}

type S1BattleProto2022RobotNtpServerIpAck struct {
	Ip uint32
}

type S1BattleProto2022RobotNtpServerIpReq struct {
	Res byte
}

type S1BattleProto2022RobotPushCapinfo struct {
	VoltageFirst     float32
	VoltageFinal     float32
	MeasureTime      float32
	OutputQEnergy    float32
	RecharingQEnergy float32
	MeasureCap       float32
	CheckCap         float32
}

type S1BattleProto2022RobotPushCaprtinfo struct {
	Voltage           float32
	OutputCurrent     float32
	RechargingCurrent float32
	VoltageFirst      float32
	MeasureTime       float32
	MeasureCap        float32
	RecharingQEnergy  float32
	OutputQEnergy     float32
}

type S1BattleProto2022RobotPushSensorInfo struct {
	Voltage           float32
	OutputCurrent     float32
	RechargingCurrent float32
	Status            byte
}

type S1BattleProto2022RobotResurgenceNotify struct {
	RobotId byte
}

type S1BattleProto2022RobotStatus struct {
	RobotStatusData S1BattleProto2022RobotStatusData
}

type S1BattleProto2022RobotStatusData struct {
	Hp             uint16
	Voltage        float32
	Current        float32
	ChassisPower   float32
	PowerBuffer    float32
	GimbalVoltage  float32
	GimbalPower    float32
	ShooterVoltage float32
	ShooterPower   float32
	Rssi           byte
	HwId           uint32
	ShooterPitch   int16
	ShooterRoll    int16
	ShooterYaw     int16
	Status         S1BattleProto2022RobotSystemStatus
	Uwb            S1BattleProto2022UwbData
	PowerState     S1BattleProto2022PowerSwitchState
	TimeStamp      uint64
}

type S1BattleProto2022RobotSystemStatus struct {
	Power          byte
	Rfid0          byte
	Rfid1          byte
	Camera         byte
	Blood0         byte
	Blood1         byte
	ShooterDetect0 byte
	ShooterDetect1 byte
	ShooterDetect2 byte
	Uwb            byte
	Armor0         byte
	Armor1         byte
	Armor2         byte
	Armor3         byte
	Armor4         byte
	Armor5         byte
	Armor6         byte
	Armor7         byte
	Armor8         byte
	Armor9         byte
	Cap            byte
	ExtA0          byte
	ExtA1          byte
	ExtA2          byte
	ExtA3          byte
	ExtA4          byte
	ExtA5          byte
	ExtA6          byte
	ExtB0          byte
	ExtB1          byte
	ExtB2          byte
	ExtC0          byte
	ExtC1          byte
	ExtC2          byte
	ReservedCount  int32
	ReservedList   []byte
}

type S1BattleProto2022RobotVtmsettingFull struct {
	ListLen byte
	Tds1    []byte
	Tds2    []byte
	Tds3    []byte
}

type S1BattleProto2022RobotWeaponFire struct {
	ShooterId  byte
	BulletType byte
	Speed      float32
	Frequency  float32
	Yaw        float32
	Pitch      float32
	RollAngle  float32
	TimeStamp  uint64
}

type S1BattleProto2022RobotsDataSync struct {
	Progress   byte
	TimeRemain uint16
	Utc        uint32
	RobotData  S1BattleProto2022RobotDataSync
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022S2CPowerStateNotify struct {
	ClientId          byte
	ChassisPowerState int32
	GimbalPowerState  int32
	ShooterPowerState int32
}

type S1BattleProto2022SendDartRobotStatus struct {
	Connectstatus byte
}

type S1BattleProto2022SetFlyCatLight struct {
	Lightcode uint32
}

type S1BattleProto2022SetFlyCatLightAck struct {
	RecvedLightcode uint32
}

type S1BattleProto2022SetFlyCatStatus struct {
	WorkState byte
}

type S1BattleProto2022SetFlyCatStatusAck struct {
	RecvedWorkState byte
}

type S1BattleProto2022SetSupplyStateReq struct {
	ReportTime uint16
	ReportReq  byte
}

type S1BattleProto2022ShooterFreqLimit struct {
	ShooterFreqThreshold    float32
	ShooterFreqMaxHurt      uint16
	ShooterFreqHurtTabCount int32
	ShooterFreqHurtTab      []byte
}

type S1BattleProto2022ShooterLimit struct {
	ShooterSpdThreshold    float32
	ShooterSpdMaxHurt      uint16
	ShooterSqdHurtTabCount int32
	ShooterSqdHurtTab      []byte
}

type S1BattleProto2022SiloCtrLedData struct {
	R byte
	G byte
	B byte
	A byte
}

type S1BattleProto2022SiloDevStatus struct {
	Online                  byte
	DoorStatus              byte
	FloorStatus             byte
	Speed                   uint16
	Angle                   uint16
	Target                  byte
	CountdownTime           float32
	ShootCheckCountdownTime float32
	OpenedTimeCount         int32
	IsDetectionWindows      byte
	MissileCount            int32
	IsCanOpen               byte
	SiloCooldown            float32
	SiloErrorCode           byte
	MissileHitCount         byte
	ErrorGateSensorUp       byte
	ErrorGateSensorDown     byte
	ErrorBaseBoardSensor    byte
	ErrorGateSensorBoth     byte
	ErrorMotorBroken        byte
	ErrorBrakeConstNo       byte
	ErrorBrakeConstYes      byte
	ErrorFacingObstacle     byte
	ErrorEmergencyStop      byte
	ErrorMotorOverHeat      byte
	ErrorCloseOverTime      byte
	ErrorOpenOverTime       byte
}

type S1BattleProto2022SiloDevStatusSyncData struct {
	SiloDevStatusSyncDataListLen int32
	SiloDevStatusDataList        []S1BattleProto2022SiloDevStatus
}

type S1BattleProto2022StuAreialEnergy struct {
	Cmd        uint16
	RemainTime byte
}

type S1BattleProto2022StuClientRecvInfo struct {
	TargetRobotId uint16
	TargetPosX    float32
	TargetPosY    float32
}

type S1BattleProto2022StuClientSendCmd struct {
	Cmd           uint16
	TargetPosX    float32
	TargetPosY    float32
	TargetPosZ    float32
	CmdKeyboard   byte
	TargetRobotId uint16
}

type S1BattleProto2022StuCommunication struct {
	DataId     uint16
	SenderId   uint16
	ReceiverId uint16
}

type S1BattleProto2022StuCommunicationByteData struct {
	Cmd             uint16
	ByteDataListLen byte
	ByteDataList    []byte
	S0HeaderBodyLen uint16
	Dataid          uint16
	Senderid        uint16
	Receiverid      uint16
}

type S1BattleProto2022StuCustomControlData struct {
	Cmd     uint16
	ListLen byte
	Data    []byte
}

type S1BattleProto2022StuEnvStatus struct {
	Cmd                     uint16
	Num1AddBloodPointStatus uint32
	Num2AddBloodPointStatus uint32
	Num3AddBloodPointStatus uint32
	RuneRfidStatus          uint32
	SmallRuneStatus         uint32
	BigRuneStatus           uint32
	RingUpland              uint32
	TrapezoidR3             uint32
	TrapezoidR4             uint32
	BaseShieldStatus        uint32
	OutpostStatus           uint32
}

type S1BattleProto2022StuGameResult struct {
	Cmd    uint16
	Winner byte
}

type S1BattleProto2022StuGameStatus struct {
	Cmd           uint16
	GameMode      byte
	GameStatus    byte
	RemainTime    uint16
	SyncTimeStamp uint64
}

type S1BattleProto2022StuIcrabuffDebuffZoneStatus struct {
	Cmd          uint16
	F1Status     byte
	F1StatusInfo byte
	F2Status     byte
	F2StatusInfo byte
	F3Status     byte
	F3StatusInfo byte
	F4Status     byte
	F4StatusInfo byte
	F5Status     byte
	F5StatusInfo byte
	F6Status     byte
	F6StatusInfo byte
}

type S1BattleProto2022StuLeftBullet struct {
	Cmd                      uint16
	SmallAvaliableBulletsNum uint16
	BigAvaliableBulletsNum   uint16
	LeftCoinsNum             uint16
}

type S1BattleProto2022StuMissileRemainingTime struct {
	Cmdid uint16
	Time  byte
}

type S1BattleProto2022StuRfidstatus struct {
	Cmd         uint16
	BaseArea    uint32
	UplandArea  uint32
	RuneArea    uint32
	FlyArea     uint32
	OutpostArea uint32
	IslandArea  uint32
	HomeArea    uint32
	SapperRfid  uint32
}

type S1BattleProto2022StuRobotBuff struct {
	Cmd         uint16
	HealBuff    byte
	CoolBuff    byte
	DefenceBuff byte
	AttackBuff  byte
}

type S1BattleProto2022StuRobotCurrentHp struct {
	Cmd           uint16
	RedHero       uint16
	RedSapper     uint16
	RedInfantry1  uint16
	RedInfantry2  uint16
	RedInfantry3  uint16
	RedGuard      uint16
	RedOutpost    uint16
	RedBase       uint16
	BlueHero      uint16
	BlueSapper    uint16
	BlueInfantry1 uint16
	BlueInfantry2 uint16
	BlueInfantry3 uint16
	BlueGuard     uint16
	BlueOutpost   uint16
	BlueBase      uint16
}

type S1BattleProto2022StuRobotHurt struct {
	Cmd      uint16
	Armor    byte
	Hurttype byte
}

type S1BattleProto2022StuSiloInfo struct {
	Cmd                  uint16
	LaunchOpeningStatus  byte
	AttackTarget         byte
	TargetChangeTime     uint16
	OperateLaunchCmdTime uint16
}

type S1BattleProto2022StuSupplyStatus struct {
	Cmd             uint16
	SupplyEnvId     byte
	SupplyRobotId   byte
	SupplyEnvStatus byte
	BulletsNum      byte
}

type S1BattleProto2022StuWarning struct {
	Cmd          uint16
	WarningLevel byte
	WarningRobot byte
}

type S1BattleProto2022SupplyClearScaleAck struct {
	Placeholder byte
}

type S1BattleProto2022SupplyClearScaleReq struct {
	Placeholder byte
}

type S1BattleProto2022SupplyExportAck struct {
	LastState byte
	NowState  byte
}

type S1BattleProto2022SupplyExportReq struct {
	ExportEnable byte
}

type S1BattleProto2022SupplyFreedAck struct {
	IsAccept byte
}

type S1BattleProto2022SupplyFreedReq struct {
	BoxFreedNum byte
}

type S1BattleProto2022SupplyReloadAck struct {
	IsAccept byte
	Action   byte
}

type S1BattleProto2022SupplyReloadReq struct {
	ReloadEnable byte
}

type S1BattleProto2022SupplyReport struct {
	BulletGear0   byte
	BulletGear1   byte
	Loadselector0 byte
	Box0          byte
	Box1          byte
	Box2          byte
	Relase        byte
	Scales        byte
	State         byte
	BoxReadyNum   byte
	BoxFreedNum   byte
}

type S1BattleProto2022SupplyResetReq struct {
	Placeholder byte
}

type S1BattleProto2022SupplyRestartReq struct {
	Placeholder byte
}

type S1BattleProto2022TestTimeDelayDownComplete struct {
	Count uint32
}

type S1BattleProto2022TestTimeDelayDownData struct {
	Seq      uint32
	CurTime  int64
	WifiTime uint32
	DataLen  int32
	Datas    []byte
}

type S1BattleProto2022TestTimeDelayDownRecord struct {
	GlobalDelay      uint32
	UartDelay        uint32
	UartLossCnt      byte
	WifiLossCnt      byte
	WifiLossTabCount int32
	WifiLossTabS     []uint16
}

type S1BattleProto2022UwbData struct {
	X          float32
	Y          float32
	Z          float32
	Compass    float32
	AnchorMask byte
	WifiEn     byte
	Rsv0       byte
	Rsv1       byte
}

type S1BattleProto2022VtmSpeedSync struct {
	Clientid           int32
	CurrentFreq        int32
	TxConnect          int32
	CurrentSpeedRate   int32
	CurrentCountryCode int32
}

type S1BattleProto2022IoctrCfgReq struct {
	Res        byte
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrRmMotorsStatus struct {
	RealPositionListLen    int32
	RealPositions          []int64
	RealSpeedListLen       int32
	RealSpeeds             []int16
	RealTemperatureListLen int32
	RealTemperatures       []byte
	ModuleId               byte
	ModuleType             byte
}

type S1BattleProto2022IoctrSetLightsRgbType1 struct {
	StartLedIndex uint16
	LedNum        uint16
	Color         S1BattleProto2022LedColor
	Seq           uint32
	ModuleId      byte
	ModuleType    byte
}

type S1BattleProto2022IoctrSetLightsRgbType1Ack struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetLightsRgbType2 struct {
	StartLedIndex uint16
	LedNum        uint16
	ListLen       int32
	LedsColors    []S1BattleProto2022LedColor
	Seq           uint32
	ModuleId      byte
	ModuleType    byte
}

type S1BattleProto2022IoctrSetLightsRgbType2Ack struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetLightsRgbType3 struct {
	LightEffect byte
	TimeSpan    uint16
	Seq         uint32
	ModuleId    byte
	ModuleType  byte
}

type S1BattleProto2022IoctrSetLightsRgbType3Ack struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetRmMotorsCfg struct {
	SpeedPidListLen       int32
	SpeedPids             []S1BattleProto2022SpeedModePid
	PositionPidListLen    int32
	PositionPids          []S1BattleProto2022PositionModePid
	TransRatioListLen     int32
	TransRatios           []float32
	IsStatusReturnListLen int32
	IsStatusReturn        []byte
	Seq                   uint32
	ModuleId              byte
	ModuleType            byte
}

type S1BattleProto2022IoctrSetRmMotorsCfgAck struct {
	ErrCode    byte
	SeqAck     uint32
	ModuleId   byte
	ModuleType byte
}

type S1BattleProto2022IoctrSetRmMotorsMoveVal struct {
	LoopModeListLen int32
	LoopMode        []byte
	Len             int32
	ExpectVal       []int64
	ModuleId        byte
	ModuleType      byte
}

type S1BattleProto2022LedColor struct {
	R byte
	G byte
	B byte
}

type S1BattleProto2022PositionModePid struct {
	Kp             float32
	Ki             float32
	Kd             float32
	NoResponse     byte
	IntergralLimit uint16
	MaxSpeed       uint16
}

type S1BattleProto2022SiloCtrlDoor struct {
	Cmd byte
}

type S1BattleProto2022SiloStatus struct {
	DoorStatus  byte
	FloorStatus byte
	Errorcode   byte
}

type S1BattleProto2022SpeedModePid struct {
	Kp        float32
	Ki        float32
	Kd        float32
	MaxOutput uint16
}

type S1BattleProto2023ClientChangeTokenCmd struct {
	Len      uint16
	NewToken string
}

type S1BattleProto2023ClientGameSystemInfoNotify struct {
	TokenLen      uint16
	CurMatchToken string
}

type S1BattleProto2023ClientGuardInPatrolInfoSync struct {
	RobotId                  byte
	IsCurrentInPatrolArea    byte
	IsLeavePatrolAreaTimeout byte
	LeavePatrolAreaLeftTime  float32
	IsCountDownActive        byte
}

type S1BattleProto2023ClientModuleVersionData struct {
	MoudleId          byte
	MoudleType        byte
	MoudleSubType     byte
	NewMoudleVersion  string
	CurMoudleVersion  string
	VersionState      byte
	MoudleIsImportant byte
}

type S1BattleProto2023ClientModuleVersionState struct {
	RobotId       byte
	ModuleNum     byte
	ModuleNumMax  byte
	ModuleVersion []S1BattleProto2023ClientModuleVersionData
}

type S1BattleProto2023ClientPenaltyInfo struct {
	RobotId       byte
	PenaltyType   byte
	PenaltyReason string
}

type S1BattleProto2023ClientPenaltyTableInfos struct {
	UploadType byte
	InfosLen   byte
	Infos      []S1BattleProto2023ClientPenaltyInfo
}

type S1BattleProto2023ClientPenaltyTableInfosAck struct {
	State      byte
	UploadType byte
}

type S1BattleProto2023ClientRobotBattleInfo struct {
	Robots0Id                    byte
	BattleState                  byte
	CanRemoteHeal                byte
	CanRemoteExchangeSmallBullet byte
	CanRemoteExchangeBigBullet   byte
	RemaindBuyRevivalCount       byte
	BuyRevivalPrice              int32
	BattleStateCountDown         float32
	IsOnHpPoint                  byte
	ChassisPowerCtrl             byte
	GimbalPowerCtrl              byte
	ShooterPowerCtrl             byte
}

type S1BattleProto2023ClientServerRobotsBattleInfoSync struct {
	RobotsNum                 byte
	RobotsBattleInfoSyncDatas []S1BattleProto2023ClientRobotBattleInfo
}

type S1BattleProto2023ClientTeamSupplyInfoSync struct {
	Teamid                          byte
	ExchangeSmallBulletPackageCount byte
	ExchangeBigBulletPackageCount   byte
	ExchangeRemoteHealCount         byte
	SentryRemainRevivalCount        byte
	SmallBulletPackageUnitPrice     float32
	BigBulletPackageUnitPrice       float32
	BigBulletUnitPrice              float32
	SmallBulletUnitPrice            float32
	CurrentRemoteHealPrice          float32
	SentryControlPrice              float32
}

type S1BattleProto2023ExchangeProErrorInfoNotify struct {
	TeamId                   byte
	IsOnline                 byte
	ArmErrorCode             uint16
	RfidErrorCode            byte
	RunningState             byte
	Ir1State                 byte
	Ir2State                 byte
	Ir3State                 byte
	IrErrorCode              byte
	TempAlertVal             int16
	TempRmMotorRoll          int16
	TempRmMotorPitch         int16
	TempRmMotorPush          int16
	MotorRollState           byte
	MotorPitchState          byte
	MotorPushState           byte
	IsSoftwareEmergencyState byte
	IsHardwareEmergencyState byte
}

type S1BattleProtoAddBulletNotify struct {
	Uid         uint64
	BulletIndex byte
}

type S1BattleProtoArmorHitAck struct {
	Result int32
}

type S1BattleProtoArmorHitReq struct {
	Index       int32
	AttackType  int32
	AttackerUid uint64
	AccValue    int32
	MicValue    int32
}

type S1BattleProtoAttrsNotify struct {
	Uid           uint64
	Tid           uint32
	Teamid        byte
	SeatIndex     byte
	Hp            int32
	HpMax         int32
	Armor         int32
	ArmorMax      int32
	Durability    int32
	DurabilityMax int32
	Bullet        int32
	BulletMax     int32
	Attack1       int32
	Attack2       float32
}

type S1BattleProtoAutoStateNotify struct {
	State int32
}

type S1BattleProtoBeAttackNotify struct {
	Index         int32
	AttackType    int32
	AttackerUid   uint64
	BeAttackerUid uint64
	Damage        int32
}

type S1BattleProtoBuffAddNotify struct {
	PlayerUid    uint64
	BuffUid      uint64
	BuffTid      uint32
	BuffLevel    uint32
	BuffVisible  byte
	BuffMaxTime  float32
	BuffLeftTime float32
	MsgParams    string
}

type S1BattleProtoBuffDelNotify struct {
	PlayerUid uint64
	BuffUid   uint64
}

type S1BattleProtoBuffModifyNotify struct {
	PlayerUid      uint64
	BuffUid        uint64
	BuffTid        uint32
	BuffLevel      uint32
	BuffVisible    byte
	BuffMaxTime    float32
	BuffLeftTime   float32
	UpdateProgress byte
	MsgParams      string
}

type S1BattleProtoClientSyncNotify struct {
	Uid             uint64
	ConnClientState int32
	ConnRobotState  int32
	Battery         int32
	SignalQuality   int32
	BattleMode      byte
}

type S1BattleProtoClientSyncReq struct {
	Uid             uint64
	ConnClientState int32
	ConnRobotState  int32
	Battery         int32
	SignalQuality   int32
	BattleMode      byte
}

type S1BattleProtoDeviceBulletNotify struct {
	BulletBottleCount int32
	BulletBottleState []byte
	PlayersUid        []uint64
}

type S1BattleProtoDeviceModuleNotify struct {
	Uid                 uint64
	ProductType         uint16
	AllDeviceCount      uint16
	AllDevice           []uint64
	AllDeviceConnection []byte
	AllDevicePrority    []byte
}

type S1BattleProtoDeviceModuleReq struct {
	Uid                  uint64
	ProductType          uint16
	ConnectedDeviceCount uint16
	ConnectedDevice      []uint64
}

type S1BattleProtoFullSceneDataReq struct {
	Nouse byte
}

type S1BattleProtoGameSettlementNotify struct {
	TeamidWin       int32
	PlayersCount    int32
	PlayersUid      []uint64
	PlayersHp       []int32
	PlayersHpMax    []int32
	PlayersBekilled []int32
	Winreason       int32
	TeamCount       int32
	TeamHurtSum     []int32
	TeamAtkCount    []int32
	TeamPojiaCount  []int32
	DurationTime    int32
}

type S1BattleProtoGunBulletNotify struct {
	Uid          uint64
	GunBulletMax int16
	GunBullet    int16
}

type S1BattleProtoGunFireAck struct {
	Result int32
}

type S1BattleProtoGunFireNotify struct {
	AttackerUid uint64
	GunType     byte
	GunSpeed    float32
}

type S1BattleProtoGunFireReq struct {
	AttackerUid uint64
	GunType     byte
	GunSpeed    float32
}

type S1BattleProtoGunHeatNotify struct {
	GunHeatMax float32
	GunHeat    float32
	PlayerUid  uint64
}

type S1BattleProtoIsCanUseAtknotify struct {
	Teamcount int32
	Teamid    []int32
	State     []int32
}

type S1BattleProtoIsCanUsePjnotify struct {
	Teamcount int32
	Teamid    []int32
	State     []int32
}

type S1BattleProtoMarkerDetectAck struct {
	Result byte
}

type S1BattleProtoMarkerDetectReq struct {
	Uid       uint64
	Act       byte
	MarkerId  byte
	MarkerStr string
}

type S1BattleProtoModuleOffLineInfoNotify struct {
	Nouse byte
}

type S1BattleProtoPlaceholderNotify struct {
	EventName string
	State     byte
}

type S1BattleProtoPlayerCommandAutoControlNotify struct {
	Value byte
}

type S1BattleProtoPlayerCommandGunFireNotify struct {
	Value byte
}

type S1BattleProtoPlayerCommandLockScreenNotify struct {
	Value byte
}

type S1BattleProtoPlayerCommandMoveNotify struct {
	Value byte
}

type S1BattleProtoPlayerCommandSetDisconnectNotify struct {
	Uid   uint64
	Value byte
}

type S1BattleProtoPlayerCommandUserCustomActionNotify struct {
	Value byte
}

type S1BattleProtoPlayerCommandUserInputNotify struct {
	Value byte
}

type S1BattleProtoPlayerDeadNotify struct {
	Uid       uint64
	KillerUid uint64
	KillHonor int32
	DeadType  int32
	Flag0     int32
	Flag1     int32
	Flag2     int32
}

type S1BattleProtoPlayerInitOkNotify struct {
	Uid uint64
}

type S1BattleProtoPlayerReliveNotify struct {
	Uid uint64
}

type S1BattleProtoPlayerResetNotify struct {
	PlayerUid uint64
}

type S1BattleProtoPoJiaNotify struct {
	Uid      uint64
	State    byte
	Timeleft float32
	Timemax  float32
}

type S1BattleProtoProgressNotify struct {
	PlayerUid uint64
	EventName string
	TimeMax   float32
	TimeLeft  float32
	State     byte
}

type S1BattleProtoRfidnotify struct {
	S0Id  byte
	Flag  byte
	Type  byte
	Area  byte
	Data0 byte
	Data1 byte
	Data2 byte
	Data3 byte
	Data4 byte
	Data5 byte
}

type S1BattleProtoRfidreq struct {
	Type      byte
	Area      byte
	Data0     byte
	Data1     byte
	Data2     byte
	Data3     byte
	Data4     byte
	Data5     byte
	Timestamp uint64
}

type S1BattleProtoReduceReliveTimeNotify struct {
	PlayerUid uint64
	TimeMax   float32
	TimeLeft  float32
}

type S1BattleProtoReloginNotify struct {
	SelfUid                         uint64
	SelfTid                         uint32
	SelfTeamId                      byte
	SelfSeatIndex                   byte
	SceneState                      int32
	SceneStateTimeLeft              int32
	SceneTeamsCount                 int32
	SceneTeamsPlayerMax             int32
	ScenePlayersCount               int32
	TeamsTotalDamage                []uint32
	PlayersTotalBuffCount           uint32
	PlayersUid                      []uint64
	PlayersTid                      []uint32
	PlayersTeamId                   []byte
	PlayersSeatIndex                []byte
	PlayersState                    []byte
	PlayersCdLeftTime               []float32
	PlayersCdMaxTime                []float32
	PlayersHp                       []uint32
	PlayersMaxHp                    []uint32
	PlayersHeat                     []float32
	PlayersMaxHeat                  []float32
	PlayersBulletNum                []uint32
	PlayersMaxBulletNum             []uint32
	PlayersArmor                    []int32
	PlayersMaxArmor                 []int32
	PlayersDurability               []int32
	PlayersMaxDurability            []int32
	PlayersBuffCount                []uint32
	PlayersBuffUid                  [][]uint64
	PlayersBuffTid                  [][]uint32
	PlayersBuffLevel                [][]uint32
	PlayersBuffLeftTime             [][]float32
	PlayersGunCtrlState             []byte
	PlayersMoveCtrlState            []byte
	PlayersUserInputState           []byte
	PlayersLockScreenState          []byte
	PlayersUserCustomActionState    []byte
	PlayersAutoModeCtrlState        []byte
	DevicesBaseStationState         []byte
	DevicesBaseStationStateLeftTime []float32
	DevicesRuneState                []byte
	DevicesRuneStateLeftTime        []float32
	WifiName                        string
	WifiPassword                    string
}

type S1BattleProtoSceneInfosNotify struct {
	PlayersCount     int32
	PlayersUid       []uint64
	PlayersTid       []uint32
	PlayersTeamId    []byte
	PlayersSeatIndex []byte
}

type S1BattleProtoSceneStateNotify struct {
	State           ESceneStateType
	TimeLeft        int32
	TeamCount       int32
	TeamHurtSum     []int32
	TeamBekillCount []int32
}

type S1BattleProtoShareBulletReq struct {
	TeamId      byte
	BulletIndex byte
}

type S1BattleProtoTeamInfoNotify struct {
	DataLen int32
	Data    string
}

type S1BattleProtoUploadModuleInfo struct {
	DataLen int32
	Data    string
}

type S1BattleProtoWarningNotify struct {
	PlayerUid    uint64
	WarningLevel byte
}

type S1BattleProtoBaseArmorStateNotify struct {
	PlayerUid  uint64
	ArmorId    int16
	ArmorCount int16
	ArmorValue int16
}

type S1BattleProtot2022RobotArmorCfgData struct {
	ArmorParaDataCount int32
	ArmorParaDatas     []S1BattleProto2022ArmorParaConfigItem
}

type S1ClientTransferRobotProto struct {
	DataLen int32
	Data    []byte
}

type S1ProtoClientExceptionInfo struct {
	DataLen int32
	Data    string
}

type S1ProtoClientNetworkInfoNotify struct {
	Robotid         byte
	VtUplinkSpeed   uint32
	VtDownlinkSpeed uint32
}

type S1ProtoClientSyncNotify struct {
	Uid               uint64
	ConnClientState   int32
	ConnRobotState    int32
	Battery           int32
	SignalQuality     int32
	BattleMode        byte
	OnlineModuleCount int32
	OnlineModule      []uint64
}

type S1ProtoClientSyncReq struct {
	Uid               uint64
	ConnClientState   int32
	ConnRobotState    int32
	Battery           int32
	SignalQuality     int32
	BattleMode        byte
	OnlineModuleCount int32
	OnlineModule      []uint64
}

type S1ProtoEnterRoomAck struct {
	ResultId EEnterRoomAckResultType
	Uid      uint64
}

type S1ProtoEnterRoomNotify struct {
	Uid uint64
}

type S1ProtoEnterRoomReq struct {
	Nouse1 byte
	Nouse2 byte
}

type S1ProtoGmcommandAck struct {
	ResultId int32
}

type S1ProtoGmcommandReq struct {
	Command string
	Pars    string
}

type S1ProtoHeader struct {
	ProtoId    uint16
	ProtoType  uint16
	DataLen    uint32
	SequenceId byte
	AckType    byte
	Nouse1     byte
	Nouse2     byte
}

type S1ProtoHeartBeatAck struct {
	ResultId   int32
	S0Clientid byte
}

type S1ProtoHeartBeatReq struct {
	Nouse      byte
	S0Clientid byte
}

type S1ProtoLeaveRoomAck struct {
	ResultId ELeaveRoomAckResultType
}

type S1ProtoLeaveRoomNotify struct {
	Uid    uint64
	Willgo ERoleLocation
}

type S1ProtoLeaveRoomReq struct {
	Nouse byte
}

type S1ProtoLoginAck struct {
	ResultId ELoginAckResultType
	Uid      uint64
	Token    string
}

type S1ProtoLoginReq struct {
	Account  string
	Password string
	Version  float32
	Tid      uint32
	Teamid   uint32
	Hash     int64
}

type S1ProtoLogoutAck struct {
	ResultId int32
}

type S1ProtoLogoutNotify struct {
	Uid    uint64
	Reason byte
}

type S1ProtoLogoutReq struct {
	Account string
}

type S1ProtoMarkDetectResult struct {
	X        float32
	Y        float32
	W        float32
	H        float32
	Pitch    float32
	Yaw      float32
	Roll     float32
	T44C2M   []float32
	Color    byte
	MarkerId uint16
	Distance uint16
}

type S1ProtoMarkDetectResultAck struct {
	Result   int32
	X        float32
	Y        float32
	W        float32
	H        float32
	Color    byte
	MarkerId uint16
	Distance uint16
}

type S1ProtoNetworkInfoNotify struct {
	Robotid           byte
	WifiUplinkSpeed   uint32
	WifiDownlinkSpeed uint32
	VtUplinkSpeed     uint32
	VtDownlinkSpeed   uint32
	WifiUplinkPlr     uint16
	WifiDownlinkPlr   uint16
	VtUplinkPlr       uint16
	VtDownlinkPlr     uint16
	OtherType         byte
	OtherPlr          uint16
	OtherDelay        uint16
	Delay             uint32
}

type S1ProtoPingAck struct {
	ResultId int32
}

type S1ProtoPingReq struct {
	Nouse byte
}

type S1ProtoPlayerCommandDisconnectNotify struct {
	Uid   uint64
	Value byte
}

type S1ProtoPlayerCommandUserInputNotify struct {
	Value byte
}

type S1ProtoReLoginAck struct {
	ResultId     ELoginAckResultType
	Uid          uint64
	Token        string
	WifiName     string
	WifiPassword string
}

type S1ProtoReLoginReq struct {
	Account  string
	Password string
	Version  float32
	Tid      uint32
}

type S1ProtoReLoginRoomNotify struct {
	SelfUid            uint64
	SelfTid            uint32
	SelfTeamId         byte
	SelfSeatIndex      byte
	RoomState          int32
	RoomStateTimeLeft  int32
	RoomSeatIslock     int32
	RoomTeamsCount     int32
	RoomTeamsPlayerMax int32
	RoomPlayersCount   int32
	PlayersUid         []uint64
	PlayersTid         []uint32
	PlayersTeamId      []byte
	PlayersSeatIndex   []byte
	WifiName           string
	WifiPassword       string
}

type S1ProtoRobotNetworkStatusAck struct {
	WifiDownlinkPlr uint16
	VtDownlinkPlr   uint16
}

type S1ProtoRobotNetworkStatusReq struct {
	Nouse byte
}

type S1ProtoRobotStudentSerialPortInfos struct {
	OtherType  byte
	OtherPlr   uint16
	OtherDelay uint16
}

type S1ProtoRoomInfosNotify struct {
	PlayersCount     int32
	PlayersUid       []uint64
	PlayersTid       []uint32
	PlayersTeamId    []byte
	PlayersSeatIndex []byte
	WifiName         string
	WifiPassword     string
	CurrMatch        string
}

type S1ProtoRoomPauseNotify struct {
	TimeElapse int32
}

type S1ProtoRoomStateNotify struct {
	State    ERoomStateType
	TimeLeft int32
	IsPause  int32
}

type S1ProtoSetCurrMatchNotify struct {
	CurrMatch string
	CurrToken string
}

type S1ProtoSetLockSeatNotify struct {
	LockSeat byte
}

type S1ProtoSetReadyAck struct {
	ResultId int32
}

type S1ProtoSetReadyNotify struct {
	Uid   uint64
	State byte
}

type S1ProtoSetReadyReq struct {
	State byte
}

type S1ProtoSetSvrInfoAck struct {
	ResultId int32
}

type S1ProtoSetSvrInfoReq struct {
	SvrName string
}

type S1ProtoSetTeamAck struct {
	ResultId  int32
	TeamId    byte
	SeatIndex byte
}

type S1ProtoSetTeamNotify struct {
	Uid       uint64
	TeamId    byte
	SeatIndex byte
}

type S1ProtoSetTeamPlayerNumNotify struct {
	PlayerNum byte
}

type S1ProtoSetTeamReq struct {
	TeamId    byte
	SeatIndex byte
}

type S1ProtoSetTidAck struct {
	ResultId int32
	Tid      uint32
}

type S1ProtoSetTidNotify struct {
	Uid uint64
	Tid uint32
}

type S1ProtoSetTidReq struct {
	Tid   uint32
	Token string
}

type S1ProtoSetTokenAck struct {
	ResultId int32
}

type S1ProtoSetTokenReq struct {
	Token string
}

type S1ProtoSetWifiInfoNotify struct {
	WifiName     string
	WifiPassword string
}

type S1ProtoSvrStateAck struct {
	State     ERoomStateType
	TimeLeft  int32
	IsPause   int32
	CurrMatch string
}

type S1ProtoSvrStateReq struct {
	Nouse byte
}

type S1ProtoTeamCampNotify struct {
	DataLen int32
	Data    string
}

type S1ProtoTeamInfoNotify struct {
	DataLen int32
	Data    string
}

type S1ProtoTeamLineupInfoNotify struct {
	DataLen int32
	Data    string
}

type S1ProtoTeamLogoNotify struct {
	Teamid  int32
	DataLen int32
	Data    string
}

type S1ProtoTechnicalPauseInfoNotify struct {
	PauseTimeType           byte
	PauseSide               byte
	RedShortPauseLeftCount  byte
	RedLongPauseLeftCount   byte
	BlueShortPauseLeftCount byte
	BlueLongPauseLeftCount  byte
	PausedDuration          uint32
}

type S1ProtoTestAck struct {
	ResultId int32
	Test     string
}

type S1ProtoTestReq struct {
	Test string
}

type S1ProtoUserStateAck struct {
	Online int32
}

type S1ProtoUserStateReq struct {
	Account  string
	Password string
}
