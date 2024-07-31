package proto

// S1BattleProto2022AirplaneStatusNotify 无人机状态协议-客户端   137-AirplaneStatusNotify
type S1BattleProto2022AirplaneStatusNotify struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// Energy
	Energy int32 `json:"energy"`
	// IsFire
	IsFire byte `json:"isFire"`
	// Curbullet
	Curbullet int16 `json:"curbullet"`
	// Maxbullet
	Maxbullet int16 `json:"maxbullet"`
	// Curshoottime
	Curshoottime float32 `json:"curshoottime"`
	// Fixshoottime
	Fixshoottime float32 `json:"fixshoottime"`
	// Cdleft
	Cdleft float32 `json:"cdleft"`
	// Countleft
	Countleft uint16 `json:"countleft"`
	// Cdrefreshcost
	Cdrefreshcost int32 `json:"cdrefreshcost"`
	// Isincd
	Isincd byte `json:"isincd"`
}

type S1BattleProto2022ArmorParaConfigItem struct {
	// TiggerPress
	TiggerPress float32 `json:"tiggerPress"`
	// BulletMaxPress
	BulletMaxPress float32 `json:"bulletMaxPress"`
	// GolfMinPress
	GolfMinPress float32 `json:"golfMinPress"`
}

type S1BattleProto2022BaseLightingEffect struct {
	// LightColor
	LightColor uint32 `json:"lightColor"`
}

// S1BattleProto2022BaseShellStatus 基地外壳状态(上行)-BaseShellStatus
type S1BattleProto2022BaseShellStatus struct {
	// BaseStatus 0:闭合状态  1:正在展开 2.正在闭合 3.展开状态
	BaseStatus byte `json:"baseStatus"`
	// BaseMotoOneStatus 一号电机状态; 0:正常  1：不正常
	BaseMotoOneStatus byte `json:"baseMotoOneStatus"`
	// BaseMotoOneBlockStatus 0:堵转  1：不堵转
	BaseMotoOneBlockStatus byte `json:"baseMotoOneBlockStatus"`
	// BaseMotoOneOnlineStatus 在线 1.不在线
	BaseMotoOneOnlineStatus byte `json:"baseMotoOneOnlineStatus"`
	// BaseMotoTwoStatus 二号电机状态; 0:正常  1：不正常
	BaseMotoTwoStatus byte `json:"baseMotoTwoStatus"`
	// BaseMotoTwoBlockStatus 0:堵转  1：不堵转
	BaseMotoTwoBlockStatus byte `json:"baseMotoTwoBlockStatus"`
	// BaseMotoTwoOnlineStatus 在线 1.不在线
	BaseMotoTwoOnlineStatus byte `json:"baseMotoTwoOnlineStatus"`
	// BaseMotoThreeStatus 三号电机状态； 0:正常  1：不正常
	BaseMotoThreeStatus byte `json:"baseMotoThreeStatus"`
	// BaseMotoThreeBlockStatus 0:堵转  1：不堵转
	BaseMotoThreeBlockStatus byte `json:"baseMotoThreeBlockStatus"`
	// BaseMotoThreeOnlineStatus 在线 1.不在线
	BaseMotoThreeOnlineStatus byte `json:"baseMotoThreeOnlineStatus"`
	// Reserved0
	Reserved0 byte `json:"reserved0"`
	// Reserved2
	Reserved2 byte `json:"reserved2"`
}

type S1BattleProto2022ControlCmdT struct {
	// OptCode
	OptCode uint16 `json:"optCode"`
	// Data
	Data uint16 `json:"data"`
}

type S1BattleProto2022ClientAirplaneCtrlReq struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// ControlCode
	ControlCode byte `json:"controlCode"`
}

// S1BattleProto2022ClientAllClientStatusSync 全部客户端在线状态同步裁判端
type S1BattleProto2022ClientAllClientStatusSync struct {
	// Listlength
	Listlength byte `json:"listlength"`
	// Status
	Status []byte `json:"status"`
}

// S1BattleProto2022ClientArmorLifeInfo 裁判端寿命查询的单条结果
type S1BattleProto2022ClientArmorLifeInfo struct {
	// ModuleId 装甲模块ID 默认-1，因装甲ID从0开始
	ModuleId int16 `json:"moduleId"`
	// LifeState 查询的寿命状态 0查询无果 1无需更换 2建议更换 3需要更换
	LifeState byte `json:"lifeState"`
}

// S1BattleProto2022ClientBaseRobotDevStatus 基地信息同步客户端-BaseRobotDevStatus
type S1BattleProto2022ClientBaseRobotDevStatus struct {
	// ShellStatus 基地设备外壳当前状态
	ShellStatus byte `json:"shellStatus"`
	// Online 基地设备是否在线  1：在线  2.离线
	Online byte `json:"online"`
	// DartTargetPosition 飞镖靶实时位置
	DartTargetPosition int16 `json:"dartTargetPosition"`
	// DartTargetIsInplace 飞镖靶是否到达点位
	DartTargetIsInplace byte `json:"dartTargetIsInplace"`
	// BaseShellOneSensorStatus 基地护甲一号传感器 0 正常 1 异常
	BaseShellOneSensorStatus byte `json:"baseShellOneSensorStatus"`
	// BaseShellTwoSensorStatus 基地护甲2号传感器 0 正常 1 异常
	BaseShellTwoSensorStatus byte `json:"baseShellTwoSensorStatus"`
	// BaseShellThreeSensorStatus 基地护甲3号传感器 0 正常 1 异常
	BaseShellThreeSensorStatus byte `json:"baseShellThreeSensorStatus"`
	// BaseDartboardSensorStatus 飞镖靶传感器 0 正常 1 异常
	BaseDartboardSensorStatus byte `json:"baseDartboardSensorStatus"`
	// BaseShellMotoOneBlockStatus 基地护甲1号电机堵转状态 0 正常 1 异常
	BaseShellMotoOneBlockStatus byte `json:"baseShellMotoOneBlockStatus"`
	// BaseShellMotoOneOnlineStatus 基地护甲1号电机在线状态 0 在线 1 离线
	BaseShellMotoOneOnlineStatus byte `json:"baseShellMotoOneOnlineStatus"`
	// BaseShellMotoTwoBlockStatus 基地护甲2号电机堵转状态  0 正常 1 堵转
	BaseShellMotoTwoBlockStatus byte `json:"baseShellMotoTwoBlockStatus"`
	// BaseShellMotoTwoOnlineStatus 基地护甲2号电机在线状态 0 在线 1 离线
	BaseShellMotoTwoOnlineStatus byte `json:"baseShellMotoTwoOnlineStatus"`
	// BaseShellMotoThreeBlockStatus 基地护甲3号电机堵转状态 0 正常 1 异常
	BaseShellMotoThreeBlockStatus byte `json:"baseShellMotoThreeBlockStatus"`
	// BaseShellMotoThreeOnlineStatus 基地护甲3号电机在线状态 0 在线 1 离线
	BaseShellMotoThreeOnlineStatus byte `json:"baseShellMotoThreeOnlineStatus"`
	// BaseDartboardMotoBlockStatus 飞镖靶电机堵转状态 0 正常 1 异常
	BaseDartboardMotoBlockStatus byte `json:"baseDartboardMotoBlockStatus"`
	// BaseDartboardMotoOnlineStatus 飞镖靶电机在线状态 0 在线 1 离线
	BaseDartboardMotoOnlineStatus byte `json:"baseDartboardMotoOnlineStatus"`
	// BaseShellSelfcheckStatus 基地护甲自检结果状态 0 正常 1 失败
	BaseShellSelfcheckStatus byte `json:"baseShellSelfcheckStatus"`
	// BaseDartboardSelfcheckStatus 飞镖靶自检结果状态 0 正常 1 失败
	BaseDartboardSelfcheckStatus byte `json:"baseDartboardSelfcheckStatus"`
	// BaseDartboardMoveTimeout 飞镖靶移动超时 0 正常 1 超时
	BaseDartboardMoveTimeout byte `json:"baseDartboardMoveTimeout"`
	// BaseShellOpenTimeout 外壳开启超时  0 正常 1 超时
	BaseShellOpenTimeout byte `json:"baseShellOpenTimeout"`
	// BaseShellCloseTimeout 外壳关闭超时  0 正常 1 超时
	BaseShellCloseTimeout byte `json:"baseShellCloseTimeout"`
	// DartBoardIronline 飞镖靶红外模块离线 0 在线 1为离线
	DartBoardIronline byte `json:"dartBoardIronline"`
	// DartBoardBrokenErr 飞镖检测靶自检后的异常码 0 正常 1 接收板红外灯珠损坏 2 飞镖引导灯损坏 3 红外灯珠和引导灯均损坏
	DartBoardBrokenErr byte `json:"dartBoardBrokenErr"`
	// IrledStatusLeft 飞镖检测靶自检后的 左接收板灯珠损坏情况，0-31bit代表对应灯珠ID
	IrledStatusLeft uint32 `json:"irledStatusLeft"`
	// IrledStatusRight 飞镖检测靶自检后的 右接收板灯珠损坏情况，0-31bit代表对应灯珠ID
	IrledStatusRight uint32 `json:"irledStatusRight"`
	// BaseShellMotoOneOverHeat 基地护甲1号电机过热  0 正常 1 异常
	BaseShellMotoOneOverHeat byte `json:"baseShellMotoOneOverHeat"`
	// BaseShellMotoTwoOverHeat 基地护甲2号电机过热 0 正常 1 异常
	BaseShellMotoTwoOverHeat byte `json:"baseShellMotoTwoOverHeat"`
	// BaseShellMotoThreeOverHeat 基地护甲3号电机过热 0 正常 1 异常
	BaseShellMotoThreeOverHeat byte `json:"baseShellMotoThreeOverHeat"`
	// BaseDartBoardMotoOverHeat 基地靶子电机过热 0 正常 1 异常
	BaseDartBoardMotoOverHeat byte `json:"baseDartBoardMotoOverHeat"`
	// BaseShellMotoOneOverCurrent 基地护甲1号电机电流过大 0 正常 1 异常
	BaseShellMotoOneOverCurrent byte `json:"baseShellMotoOneOverCurrent"`
	// BaseShellMotoTwoOverCurrent 基地护甲2号电机电流过大 0 正常 1 异常
	BaseShellMotoTwoOverCurrent byte `json:"baseShellMotoTwoOverCurrent"`
	// BaseShellMotoThreeOverCurrent 基地护甲3号电机电流过大 0 正常 1 异常
	BaseShellMotoThreeOverCurrent byte `json:"baseShellMotoThreeOverCurrent"`
	// BaseDartBoardMotoOverCurrent 基地飞镖靶电机电流过大 0 正常 1 异常
	BaseDartBoardMotoOverCurrent byte `json:"baseDartBoardMotoOverCurrent"`
	// BaseShellMotoOneOverSpeed 基地护甲1号电机差速过大 0 正常 1 异常
	BaseShellMotoOneOverSpeed byte `json:"baseShellMotoOneOverSpeed"`
	// BaseShellMotoTwoOverSpeed 基地护甲2号电机差速过大 0 正常 1 异常
	BaseShellMotoTwoOverSpeed byte `json:"baseShellMotoTwoOverSpeed"`
	// BaseShellMotoThreeOverSpeed 基地护甲3号电机差速过大 0 正常 1 异常
	BaseShellMotoThreeOverSpeed byte `json:"baseShellMotoThreeOverSpeed"`
	// BaseDartBoardMotoOverSpeed 基地飞镖靶电机差速过大 0 正常 1 异常
	BaseDartBoardMotoOverSpeed byte `json:"baseDartBoardMotoOverSpeed"`
}

// S1BattleProto2022ClientBaseRobotDevStatusSyncData 基地信息同步客户端-BaseRobotDevStatusSyncData
type S1BattleProto2022ClientBaseRobotDevStatusSyncData struct {
	// BaseRobotDevStatusListLen
	BaseRobotDevStatusListLen byte `json:"baseRobotDevStatusListLen"`
	// BaseRobotDevStatusList
	BaseRobotDevStatusList []S1BattleProto2022ClientBaseRobotDevStatus `json:"baseRobotDevStatusList"`
}

// S1BattleProto2022ClientBattleFirstData 客户端协议-BattleFirstData
type S1BattleProto2022ClientBattleFirstData struct {
	// Progress
	Progress byte `json:"progress"`
	// IsPaused
	IsPaused byte `json:"isPaused"`
}

// S1BattleProto2022ClientBuffItem buff详细信息-BuffItem
type S1BattleProto2022ClientBuffItem struct {
	// BuffId BuffID
	BuffId byte `json:"buffId"`
	// CdTime CdTime
	CdTime float32 `json:"cdTime"`
}

// S1BattleProto2022ClientBuffSlot 机器人BUFF槽-BuffSlot
type S1BattleProto2022ClientBuffSlot struct {
	// BuffItemsLen
	BuffItemsLen int32 `json:"buffItemsLen"`
	// BuffItems
	BuffItems []S1BattleProto2022ClientBuffItem `json:"buffItems"`
}

type S1BattleProto2022ClientBuyBulletAck struct {
	// Result
	Result int32 `json:"result"`
}

type S1BattleProto2022ClientBuyBulletReq struct {
	// Type
	Type byte `json:"type"`
	// Count
	Count byte `json:"count"`
}

// S1BattleProto2022ClientCheckInTimeStampNotify 检录时间戳同步给客户都安-CheckInTimeStampNotifyData
type S1BattleProto2022ClientCheckInTimeStampNotify struct {
	// RobotidListLen
	RobotidListLen byte `json:"robotidListLen"`
	// RobotidList
	RobotidList []uint32 `json:"robotidList"`
}

// S1BattleProto2022ClientCurrentProgress 比赛当前进程-CurrentProgress
type S1BattleProto2022ClientCurrentProgress struct {
	// CurProgress
	CurProgress int32 `json:"curProgress"`
	// IsPaused
	IsPaused byte `json:"isPaused"`
}

// S1BattleProto2022ClientCustomControlData 自定义控制器-CustomControlData
type S1BattleProto2022ClientCustomControlData struct {
	// RobotId
	RobotId int32 `json:"robotId"`
	// DataLen
	DataLen int32 `json:"dataLen"`
	// DataListLen
	DataListLen byte `json:"dataListLen"`
	// DataList
	DataList []byte `json:"dataList"`
}

type S1BattleProto2022ClientEconomyNotify struct {
	// Clientid
	Clientid byte `json:"clientid"`
	// Type
	Type int32 `json:"type"`
	// Money
	Money int32 `json:"money"`
}

// S1BattleProto2022ClientExceptionCapDataNotify 电容检测-电容异常同步给裁判端-ExceptionCapData
type S1BattleProto2022ClientExceptionCapDataNotify struct {
	// RobotidListLen
	RobotidListLen byte `json:"robotidListLen"`
	// RobotidList
	RobotidList []byte `json:"robotidList"`
}

// S1BattleProto2022ClientExceptionRecoverCardNotify 补血卡位置异常-ExceptionRecoverCardData
type S1BattleProto2022ClientExceptionRecoverCardNotify struct {
	// Teamid 队伍id
	Teamid byte `json:"teamid"`
	// Robotid 机器人id
	Robotid byte `json:"robotid"`
	// Exception 异常标识符（备用）
	Exception byte `json:"exception"`
}

// S1BattleProto2022ClientExchangeProStateNotify 兑换站pro状态同步协议-ExchangeProStateNotify
type S1BattleProto2022ClientExchangeProStateNotify struct {
	// RobotId 兑换站id
	RobotId byte `json:"robotId"`
	// IsOnline
	IsOnline byte `json:"isOnline"`
	// ExchangeLevel 选手最近一次选择的档位
	ExchangeLevel byte `json:"exchangeLevel"`
	// IrState 红外状态，结合下面两个字段给的结果
	IrState byte `json:"irState"`
	// Ir1 红外1状态
	Ir1 byte `json:"ir1"`
	// Ir2 红外2状态
	Ir2 byte `json:"ir2"`
	// X 机械臂坐标
	X float32 `json:"x"`
	// Y
	Y float32 `json:"y"`
	// Z
	Z float32 `json:"z"`
	// Pitch
	Pitch float32 `json:"pitch"`
	// Roll
	Roll float32 `json:"roll"`
	// Yaw
	Yaw float32 `json:"yaw"`
	// ErrorCode 错误码
	ErrorCode byte `json:"errorCode"`
	// GoldCount 金矿兑换数量
	GoldCount byte `json:"goldCount"`
	// SilverCount 银矿兑换数量
	SilverCount byte `json:"silverCount"`
	// State 兑换站pro兑换流程状态
	State byte `json:"state"`
	// MineralRfidState 兑换站pro矿石rfid状态
	MineralRfidState byte `json:"mineralRfidState"`
	// CoinsExchange 累计兑换得到的金币数
	CoinsExchange int32 `json:"coinsExchange"`
	// CoinsAddRate 金币加成倍率
	CoinsAddRate float32 `json:"coinsAddRate"`
	// EngineerLevel 工程等级
	EngineerLevel byte `json:"engineerLevel"`
}

// S1BattleProto2022ClientFlycatStatusSync 飞猫状态同步-FlycatStatusSync
type S1BattleProto2022ClientFlycatStatusSync struct {
	// TeamId 0:红 1:蓝
	TeamId byte `json:"teamId"`
	// Online 1：在线  0.离线
	Online byte `json:"online"`
	// SysState 0正常 1警告
	SysState byte `json:"sysState"`
	// CtrlState 控制方式 0 遥控器 1 服务器
	CtrlState byte `json:"ctrlState"`
	// WorkState 工作状态 对应控制指令枚举
	WorkState byte `json:"workState"`
	// Battery 当前电量
	Battery byte `json:"battery"`
	// MotorOneOnlineError 一号电机在线状态 0 正常 1 异常
	MotorOneOnlineError byte `json:"motorOneOnlineError"`
	// MotorTwoOnlineError 二号电机在线状态 0 正常 1 异常
	MotorTwoOnlineError byte `json:"motorTwoOnlineError"`
	// MotorOneOverheatError 一号电机温度过高 0 正常 1 异常
	MotorOneOverheatError byte `json:"motorOneOverheatError"`
	// MotorTwoOverheatError 二号电机温度过高 0 正常 1 异常
	MotorTwoOverheatError byte `json:"motorTwoOverheatError"`
	// LowBatteryWarning 飞猫低电量警告 0 正常 1 警告 2 立即更换
	LowBatteryWarning byte `json:"lowBatteryWarning"`
	// LowBatteryWarningThreshold 飞猫低电量警告阈值30
	LowBatteryWarningThreshold byte `json:"lowBatteryWarningThreshold"`
	// BatteryRenewThreshold 飞猫电量耗尽立刻更换警告阈值10
	BatteryRenewThreshold byte `json:"batteryRenewThreshold"`
}

type S1BattleProto2022ClientGmcommand struct {
	// Len
	Len byte `json:"len"`
	// Cmd
	Cmd string `json:"cmd"`
}

// S1BattleProto2022ClientHitNotify 受击通知-HitNotify
type S1BattleProto2022ClientHitNotify struct {
	// RobotId 机器人id
	RobotId byte `json:"robotId"`
	// OnHitType 受击类型
	OnHitType byte `json:"onHitType"`
	// ArmorNumber 受击的装甲号
	ArmorNumber byte `json:"armorNumber"`
	// HpReduce HP变化值
	HpReduce int32 `json:"hpReduce"`
	// HpCurr HP当前值
	HpCurr int32 `json:"hpCurr"`
	// HpMax HP最大值
	HpMax int32 `json:"hpMax"`
}

// S1BattleProto2022ClientHoldCenterPointCompleteNotify 中心增益点-完全占领同步协议
type S1BattleProto2022ClientHoldCenterPointCompleteNotify struct {
	// Teamid 完全占领的队伍id
	Teamid int32 `json:"teamid"`
	// RewardExp 奖励经验
	RewardExp float32 `json:"rewardExp"`
}

// S1BattleProto2022ClientHpChangeDetailVal 客户端协议-各伤害数据值hp_change_detail_val
type S1BattleProto2022ClientHpChangeDetailVal struct {
	// ByBullet 小子弹
	ByBullet int32 `json:"byBullet"`
	// ByGolf 大子弹
	ByGolf int32 `json:"byGolf"`
	// ByMissile 导弹
	ByMissile int32 `json:"byMissile"`
	// ByOverSpeed 超速掉血
	ByOverSpeed int32 `json:"byOverSpeed"`
	// ByoverFreq 超频掉血
	ByoverFreq int32 `json:"byoverFreq"`
	// ByOverPower 超功率
	ByOverPower int32 `json:"byOverPower"`
	// ByOverHeat 超热量
	ByOverHeat int32 `json:"byOverHeat"`
	// ByModuleOffline 模块离线
	ByModuleOffline int32 `json:"byModuleOffline"`
	// ByPunish 惩罚
	ByPunish int32 `json:"byPunish"`
	// ByKill 服务器杀死
	ByKill int32 `json:"byKill"`
	// ByCrash 撞击
	ByCrash int32 `json:"byCrash"`
	// ByWifiOffline wifi离线
	ByWifiOffline int32 `json:"byWifiOffline"`
	// All 总和
	All int32 `json:"all"`
}

// S1BattleProto2022ClientKickOffRobotNotify 踢出机器人通知-KickOffRobotNotify
type S1BattleProto2022ClientKickOffRobotNotify struct {
	// RobotId 机器人ID
	RobotId int32 `json:"robotId"`
	// Reason 1机动枪口超限  2主裁踢出
	Reason byte `json:"reason"`
}

type S1BattleProto2022ClientMineralExchangeNotify struct {
	// TeamId
	TeamId byte `json:"teamId"`
	// MineralType
	MineralType byte `json:"mineralType"`
	// Value
	Value int32 `json:"value"`
}

type S1BattleProto2022ClientMineralInfoSync struct {
	// IsRedConnect
	IsRedConnect byte `json:"isRedConnect"`
	// IsBlueConnect
	IsBlueConnect byte `json:"isBlueConnect"`
	// RedGoldCount
	RedGoldCount byte `json:"redGoldCount"`
	// RedSilverCount
	RedSilverCount byte `json:"redSilverCount"`
	// BlueGoldCount
	BlueGoldCount byte `json:"blueGoldCount"`
	// BlueSilverCount
	BlueSilverCount byte `json:"blueSilverCount"`
	// RedInfraredStateListLen
	RedInfraredStateListLen byte `json:"redInfraredStateListLen"`
	// RedInfraredState
	RedInfraredState []byte `json:"redInfraredState"`
	// BlueInfraredStateListLen
	BlueInfraredStateListLen byte `json:"blueInfraredStateListLen"`
	// BlueInfraredState
	BlueInfraredState []byte `json:"blueInfraredState"`
}

// S1BattleProto2022ClientOutpostDeviceStatusSync 前哨站道具状态同步
type S1BattleProto2022ClientOutpostDeviceStatusSync struct {
	// RobotId 前哨战id
	RobotId byte `json:"robotId"`
	// Online 1：在线  0.离线
	Online byte `json:"online"`
	// StatusCode 状态码 0；空闲  1：顺时针旋转 2：逆时针旋转 3:复位中 4:复位完成
	StatusCode byte `json:"statusCode"`
	// CurTemperature 当前温度
	CurTemperature int32 `json:"curTemperature"`
	// TemperatureWarning 温度预警  0：正常  1：警告
	TemperatureWarning byte `json:"temperatureWarning"`
	// SpinSpeedRatio 转速比例 0-1 0:不转 0.5:半速 1:满速
	SpinSpeedRatio float32 `json:"spinSpeedRatio"`
	// MotorOnline 电机在线 0 离线 1在线
	MotorOnline byte `json:"motorOnline"`
	// MotorCurrentWarning 电机电流过大 0 正常 1异常
	MotorCurrentWarning byte `json:"motorCurrentWarning"`
	// MotorSpeedoverWarning 电机转速差过大 0 正常 1 异常
	MotorSpeedoverWarning byte `json:"motorSpeedoverWarning"`
	// IsMagnetOn 电磁铁上电状态 0断电 1上电
	IsMagnetOn byte `json:"isMagnetOn"`
	// ActionTimeout 动作超时，暂时预留，0 正常 其他为异常
	ActionTimeout byte `json:"actionTimeout"`
	// DartBoardIrOnline 飞镖靶红外模块离线 0 离线 1为在线
	DartBoardIrOnline byte `json:"dartBoardIrOnline"`
	// DartBoardBrokenErr 前哨战飞镖检测靶自检后的异常码 0 正常 1 接收板红外灯珠损坏 2 飞镖引导灯损坏 3 红外灯珠和引导灯均损坏
	DartBoardBrokenErr byte `json:"dartBoardBrokenErr"`
	// IrledStatusLeft 飞镖检测靶自检后的 左接收板灯珠损坏情况，0-31bit代表对应灯珠ID
	IrledStatusLeft uint32 `json:"irledStatusLeft"`
	// IrledStatusRight 飞镖检测靶自检后的 右接收板灯珠损坏情况，0-31bit代表对应灯珠ID
	IrledStatusRight uint32 `json:"irledStatusRight"`
}

// S1BattleProto2022ClientRaderInfoToClient 雷达站信息发给客户端-RaderInfoToClient
type S1BattleProto2022ClientRaderInfoToClient struct {
	// TargetRobotId
	TargetRobotId uint16 `json:"targetRobotId"`
	// TargetPosX 单位m
	TargetPosX float32 `json:"targetPosX"`
	// TargetPosY
	TargetPosY float32 `json:"targetPosY"`
	// TorwardAngle
	TorwardAngle float32 `json:"torwardAngle"`
}

// S1BattleProto2022ClientRecoverCardStatus 工程回血卡同步协议-RecoverCardStatus
type S1BattleProto2022ClientRecoverCardStatus struct {
	// Redcard 1.启用  2.停用
	Redcard byte `json:"redcard"`
	// Bluecard 1.启用  2.停用
	Bluecard byte `json:"bluecard"`
}

// S1BattleProto2022ClientResultDataInfo 比赛结算协议-ResultDataInfo
type S1BattleProto2022ClientResultDataInfo struct {
	// WebGameId 流程系统的game_id
	WebGameId int32 `json:"webGameId"`
	// GameId
	GameId int64 `json:"gameId"`
	// GameOverReasonId
	GameOverReasonId byte `json:"gameOverReasonId"`
	// GameOverReasonListLen
	GameOverReasonListLen int32 `json:"gameOverReasonListLen"`
	// GameOverReason 比赛结束原因
	GameOverReason []byte `json:"gameOverReason"`
	// StartTime 开始时间
	StartTime int64 `json:"startTime"`
	// DurationTime 比赛时长
	DurationTime int32 `json:"durationTime"`
	// GameOrder 第几局
	GameOrder byte `json:"gameOrder"`
	// TotalRound 总共几局
	TotalRound byte `json:"totalRound"`
	// RecordsListLen
	RecordsListLen int32 `json:"recordsListLen"`
	// Records 比分记录
	Records []int32 `json:"records"`
	// Winner 胜利ID，1红，2蓝，3平局
	Winner byte `json:"winner"`
	// ScoreRed 红方比分
	ScoreRed int32 `json:"scoreRed"`
	// ScoreBlue 蓝方比分
	ScoreBlue int32 `json:"scoreBlue"`
	// RedCurrentHp 红方当前血量
	RedCurrentHp int32 `json:"redCurrentHp"`
	// BlueCurrentHp 蓝方当前血量
	BlueCurrentHp int32 `json:"blueCurrentHp"`
	// RedTotalHp 红方总血量
	RedTotalHp int32 `json:"redTotalHp"`
	// BlueTotalHp 蓝方总血量
	BlueTotalHp int32 `json:"blueTotalHp"`
	// RedWarning1 红方一级警告
	RedWarning1 int32 `json:"redWarning1"`
	// RedWarning2 红方二级警告
	RedWarning2 int32 `json:"redWarning2"`
	// RedWarning3 红方三级警告
	RedWarning3 int32 `json:"redWarning3"`
	// BlueWarning1 蓝方一级警告
	BlueWarning1 int32 `json:"blueWarning1"`
	// BlueWarning2 蓝方二级警告
	BlueWarning2 int32 `json:"blueWarning2"`
	// BlueWarning3 蓝方三级警告
	BlueWarning3 int32 `json:"blueWarning3"`
	// RedHits 红方造成的伤害
	RedHits int32 `json:"redHits"`
	// BlueHits 蓝方造成的伤害
	BlueHits int32 `json:"blueHits"`
	// RedRuneCount 红方激活能量机关次数
	RedRuneCount int32 `json:"redRuneCount"`
	// BlueRuneCount 蓝方激活能量机关次数
	BlueRuneCount int32 `json:"blueRuneCount"`
	// RedKillCount 红方击杀对方次数
	RedKillCount int32 `json:"redKillCount"`
	// BlueKillCount 蓝方击杀对方次数
	BlueKillCount int32 `json:"blueKillCount"`
	// RedDartHitCount 红方飞镖击中数目
	RedDartHitCount int32 `json:"redDartHitCount"`
	// BlueDartHitCount 蓝方飞镖击中数目
	BlueDartHitCount int32 `json:"blueDartHitCount"`
	// RedShootSmallCount 红方发射小子弹数量（包括哨兵+飞机）
	RedShootSmallCount int32 `json:"redShootSmallCount"`
	// RedShootBigCount 红方发射大弹丸数量
	RedShootBigCount int32 `json:"redShootBigCount"`
	// BlueShootSmallCount 蓝方发射小弹丸数量（包括哨兵+飞机）
	BlueShootSmallCount int32 `json:"blueShootSmallCount"`
	// BlueShootBigCount 蓝方发射大弹丸数量
	BlueShootBigCount int32 `json:"blueShootBigCount"`
	// RedGuardLeftLives 红方哨兵剩余命数
	RedGuardLeftLives byte `json:"redGuardLeftLives"`
	// BlueGuardLeftLives 蓝方哨兵剩余命数
	BlueGuardLeftLives byte `json:"blueGuardLeftLives"`
	// GuardFixedLives 哨兵固定总命数
	GuardFixedLives byte `json:"guardFixedLives"`
	// RobotsNum
	RobotsNum int32 `json:"robotsNum"`
	// RobotsInfos
	RobotsInfos []S1BattleProto2022ClientRobotDataInfo `json:"robotsInfos"`
	// RedSnipeSuccCount 红方狙击命中次数
	RedSnipeSuccCount byte `json:"redSnipeSuccCount"`
	// BlueSnipeSuccCount 蓝方狙击命中次数
	BlueSnipeSuccCount byte `json:"blueSnipeSuccCount"`
	// RedUavCallCount 红方红方空中支援次数
	RedUavCallCount byte `json:"redUavCallCount"`
	// BlueUavCallCount 蓝方红方空中支援次数
	BlueUavCallCount byte `json:"blueUavCallCount"`
	// RedRadarBuffDoubleUsedCount 红方雷达双倍易伤使用次数
	RedRadarBuffDoubleUsedCount byte `json:"redRadarBuffDoubleUsedCount"`
	// BlueRadarBuffDoubleUsedCount 蓝方雷达双倍易伤使用次数
	BlueRadarBuffDoubleUsedCount byte `json:"blueRadarBuffDoubleUsedCount"`
}

// S1BattleProto2022ClientRobotArmorLifeQueryResult 裁判端寿命查询结果汇总返回
type S1BattleProto2022ClientRobotArmorLifeQueryResult struct {
	// IsWindMill 是否能量机关(因robotid为1
	IsWindMill byte `json:"isWindMill"`
	// WindMillTeamId 能量机关红蓝边 0红1蓝
	WindMillTeamId byte `json:"windMillTeamId"`
	// RobotId 机器人s0id
	RobotId byte `json:"robotId"`
	// RobotTotalArmorNum 机器人装甲总数
	RobotTotalArmorNum int32 `json:"robotTotalArmorNum"`
	// InfosLen 信息列表固定个数
	InfosLen byte `json:"infosLen"`
	// LifeInfos 装甲寿命信息
	LifeInfos []S1BattleProto2022ClientArmorLifeInfo `json:"lifeInfos"`
}

// S1BattleProto2022ClientRobotBaseDataSync 客户端协议-RobotBaseDataSync
type S1BattleProto2022ClientRobotBaseDataSync struct {
	// UserId 唯一ID号
	UserId uint16 `json:"userId"`
	// Id 机器人ID号
	Id byte `json:"id"`
	// Type 机器人类型
	Type byte `json:"type"`
	// Level 机器人等级
	Level byte `json:"level"`
	// CpuId 机器人的CPU ID（登录的时候发送）
	CpuId uint32 `json:"cpuId"`
	// FixHp 机器人总血量
	FixHp int32 `json:"fixHp"`
	// FixPower 机器人最大功率
	FixPower float32 `json:"fixPower"`
	// FixPowerBufferRecoverable 可恢复功率缓冲值
	FixPowerBufferRecoverable float32 `json:"fixPowerBufferRecoverable"`
	// FixPowerBufferUnrecoverable 不可恢复功率缓冲值
	FixPowerBufferUnrecoverable float32 `json:"fixPowerBufferUnrecoverable"`
	// FixSmallSpeed 小枪口最大射速
	FixSmallSpeed float32 `json:"fixSmallSpeed"`
	// FixSmallFreq 小枪口最大射频
	FixSmallFreq float32 `json:"fixSmallFreq"`
	// FixSmallHeatEnergy 小枪口最大热量
	FixSmallHeatEnergy float32 `json:"fixSmallHeatEnergy"`
	// FixSmallHeatEnergyCoolRate 小枪口最大冷却
	FixSmallHeatEnergyCoolRate float32 `json:"fixSmallHeatEnergyCoolRate"`
	// FixSmallSpeed2 额外小枪口最大射速
	FixSmallSpeed2 float32 `json:"fixSmallSpeed2"`
	// FixSmallFreq2 额外小枪口最大射频
	FixSmallFreq2 float32 `json:"fixSmallFreq2"`
	// FixSmallHeatEnergy2 额外小枪口最大热量
	FixSmallHeatEnergy2 float32 `json:"fixSmallHeatEnergy2"`
	// FixSmallHeatEnergyCoolRate2 额外小枪口最大冷却
	FixSmallHeatEnergyCoolRate2 float32 `json:"fixSmallHeatEnergyCoolRate2"`
	// FixBigSpeed 大枪口最大射速
	FixBigSpeed float32 `json:"fixBigSpeed"`
	// FixBigFreq 大枪口最大射频
	FixBigFreq float32 `json:"fixBigFreq"`
	// FixBigHeatEnergy 大枪口最大热量
	FixBigHeatEnergy float32 `json:"fixBigHeatEnergy"`
	// FixBigHeatEnergyCoolRate 大枪口最大冷却
	FixBigHeatEnergyCoolRate float32 `json:"fixBigHeatEnergyCoolRate"`
	// ExpOffer 死亡提供经验
	ExpOffer float32 `json:"expOffer"`
	// CurExpAddRate 经验增长速率
	CurExpAddRate float32 `json:"curExpAddRate"`
	// CurRebornCd 复活间隔
	CurRebornCd int32 `json:"curRebornCd"`
	// FixExpUpgradeNeed 升级经验
	FixExpUpgradeNeed float32 `json:"fixExpUpgradeNeed"`
	// FixExpAddRate 经验增长速率
	FixExpAddRate float32 `json:"fixExpAddRate"`
	// FixRebornCd 复活间隔
	FixRebornCd int32 `json:"fixRebornCd"`
}

type S1BattleProto2022ClientRobotDataInfo struct {
	// RobotId
	RobotId int32 `json:"robotId"`
	// RobotType
	RobotType int32 `json:"robotType"`
	// CurHp
	CurHp int32 `json:"curHp"`
	// MaxHp
	MaxHp int32 `json:"maxHp"`
	// IsConnected
	IsConnected byte `json:"isConnected"`
	// DeadReason
	DeadReason int32 `json:"deadReason"`
	// BehitedVal
	BehitedVal int32 `json:"behitedVal"`
	// TotalAttack
	TotalAttack int32 `json:"totalAttack"`
}

// S1BattleProto2022ClientRobotDeathNotify 死亡通知-HitNotify
type S1BattleProto2022ClientRobotDeathNotify struct {
	// RobotidDeath
	RobotidDeath byte `json:"robotidDeath"`
	// RobotidKiller
	RobotidKiller byte `json:"robotidKiller"`
	// DeathReason
	DeathReason byte `json:"deathReason"`
	// BFirstBlood
	BFirstBlood byte `json:"bfirstBlood"`
	// KillCount
	KillCount int32 `json:"killCount"`
}

// S1BattleProto2022ClientRobotFirstData 客户端协议-RobotFirstData
type S1BattleProto2022ClientRobotFirstData struct {
	// Userid
	Userid uint16 `json:"userid"`
}

// S1BattleProto2022ClientRobotMapData 机器人小地图数据-RobotMapData
type S1BattleProto2022ClientRobotMapData struct {
	// Joinstate
	Joinstate byte `json:"joinstate"`
	// IsAlive
	IsAlive byte `json:"isAlive"`
	// X
	X float32 `json:"x"`
	// Y
	Y float32 `json:"y"`
	// Yaw
	Yaw float32 `json:"yaw"`
}

// S1BattleProto2022ClientRobotStatusNotify 客户端协议-RobotStatusNotify-每 0.2s 同步一次机器人状态
type S1BattleProto2022ClientRobotStatusNotify struct {
	// RobotUserId
	RobotUserId int32 `json:"robotUserId"`
	// ConnState
	ConnState byte `json:"connState"`
	// JoinState
	JoinState byte `json:"joinState"`
	// SurvivalState
	SurvivalState byte `json:"survivalState"`
	// DeathSubState
	DeathSubState int32 `json:"deathSubState"`
	// WifiWeak
	WifiWeak byte `json:"wifiWeak"`
	// ModuleOnline
	ModuleOnline byte `json:"moduleOnline"`
	// BatteryLow
	BatteryLow byte `json:"batteryLow"`
	// IdConflict
	IdConflict byte `json:"idConflict"`
	// IsCanReliveByClient
	IsCanReliveByClient byte `json:"isCanReliveByClient"`
}

// S1BattleProto2022ClientRobotSyncData 客户端协议-RobotSyncData
type S1BattleProto2022ClientRobotSyncData struct {
	// UserId 唯一ID号
	UserId uint16 `json:"userId"`
	// CurHp 机器人血量
	CurHp uint16 `json:"curHp"`
	// Voltage 机器人电压
	Voltage float32 `json:"voltage"`
	// Current 机器人电流
	Current float32 `json:"current"`
	// Rssi 机器人信号强度
	Rssi byte `json:"rssi"`
	// CurPower 机器人功率
	CurPower float32 `json:"curPower"`
	// YellowWarningCount 机器人黄牌计数
	YellowWarningCount byte `json:"yellowWarningCount"`
	// PowerBuffer 60J能量
	PowerBuffer float32 `json:"powerBuffer"`
	// CurSmallSpeed1 小枪口1射速
	CurSmallSpeed1 float32 `json:"curSmallSpeed1"`
	// CurSmallFreq1 小枪口1射频
	CurSmallFreq1 float32 `json:"curSmallFreq1"`
	// CurSmallHeatEnergy1 小枪口1热量
	CurSmallHeatEnergy1 float32 `json:"curSmallHeatEnergy1"`
	// CurSmallHeatEnergyCoolRate1 小枪口1冷却值（秒）
	CurSmallHeatEnergyCoolRate1 float32 `json:"curSmallHeatEnergyCoolRate1"`
	// SmallShootNum1 小枪口1发弹量
	SmallShootNum1 int32 `json:"smallShootNum1"`
	// SmallLeftBulletsNum1 小枪口1剩余发弹量
	SmallLeftBulletsNum1 int32 `json:"smallLeftBulletsNum1"`
	// CurSmallSpeed2 小枪口2射速
	CurSmallSpeed2 float32 `json:"curSmallSpeed2"`
	// CurSmallFreq2 小枪口2射频
	CurSmallFreq2 float32 `json:"curSmallFreq2"`
	// CurSmallHeatEnergy2 小枪口2热量
	CurSmallHeatEnergy2 float32 `json:"curSmallHeatEnergy2"`
	// CurSmallHeatEnergyCoolRate2 小枪口2冷却值（秒）
	CurSmallHeatEnergyCoolRate2 float32 `json:"curSmallHeatEnergyCoolRate2"`
	// SmallShootNum2 小枪口2发弹量
	SmallShootNum2 int32 `json:"smallShootNum2"`
	// SmallLeftBulletsNum2 小枪口2剩余发弹量
	SmallLeftBulletsNum2 int32 `json:"smallLeftBulletsNum2"`
	// CurBigSpeed 大枪口射速
	CurBigSpeed float32 `json:"curBigSpeed"`
	// CurBigFreq 大枪口射频
	CurBigFreq float32 `json:"curBigFreq"`
	// CurBigHeatEnergy 大枪口热量
	CurBigHeatEnergy float32 `json:"curBigHeatEnergy"`
	// CurBigHeatEnergyCoolRate 大枪口冷却值（秒）
	CurBigHeatEnergyCoolRate float32 `json:"curBigHeatEnergyCoolRate"`
	// BigBulletsShootCnt 大枪口发弹量
	BigBulletsShootCnt int32 `json:"bigBulletsShootCnt"`
	// BigLeftBulletCnt 大枪口剩余发弹量
	BigLeftBulletCnt int32 `json:"bigLeftBulletCnt"`
	// CurExp 当前经验值
	CurExp float32 `json:"curExp"`
	// ByAttackTotal 机器人在比赛中被打击的伤害值（不会随着机器人复活而重置；一场比赛结束会重置）
	ByAttackTotal int32 `json:"byAttackTotal"`
	// MurderId 凶手机器人ID
	MurderId byte `json:"murderId"`
	// Power power模块在线状态(注：仅用于标识wifi模块是否在线， 不用于机器人与服务器的连接状态)
	Power byte `json:"power"`
	// RfidNum RFID在线状态
	RfidNum byte `json:"rfidNum"`
	// Rfid0
	Rfid0 byte `json:"rfid0"`
	// Rfid1
	Rfid1 byte `json:"rfid1"`
	// BloodNum 灯条在线状态
	BloodNum byte `json:"bloodNum"`
	// Blood0
	Blood0 byte `json:"blood0"`
	// Blood1
	Blood1 byte `json:"blood1"`
	// Blood2
	Blood2 byte `json:"blood2"`
	// ShooterDetectNum 测试模块在线状态
	ShooterDetectNum byte `json:"shooterDetectNum"`
	// ShooterDetect0
	ShooterDetect0 byte `json:"shooterDetect0"`
	// ShooterDetect1
	ShooterDetect1 byte `json:"shooterDetect1"`
	// ShooterDetect2
	ShooterDetect2 byte `json:"shooterDetect2"`
	// HasshooterDetect0 测试模块存在状态
	HasshooterDetect0 byte `json:"hasshooterDetect0"`
	// HasshooterDetect1
	HasshooterDetect1 byte `json:"hasshooterDetect1"`
	// HasshooterDetect2
	HasshooterDetect2 byte `json:"hasshooterDetect2"`
	// UwbNum uwb在线状态
	UwbNum byte `json:"uwbNum"`
	// Uwb
	Uwb byte `json:"uwb"`
	// StrikeNum 装甲在线情况
	StrikeNum byte `json:"strikeNum"`
	// Strike0
	Strike0 byte `json:"strike0"`
	// Strike1
	Strike1 byte `json:"strike1"`
	// Strike2
	Strike2 byte `json:"strike2"`
	// Strike3
	Strike3 byte `json:"strike3"`
	// Strike4
	Strike4 byte `json:"strike4"`
	// Strike5
	Strike5 byte `json:"strike5"`
	// Strike6
	Strike6 byte `json:"strike6"`
	// Strike7
	Strike7 byte `json:"strike7"`
	// Strike8
	Strike8 byte `json:"strike8"`
	// Strike9
	Strike9 byte `json:"strike9"`
	// CameraNum 图传在线情况
	CameraNum byte `json:"cameraNum"`
	// Camera
	Camera byte `json:"camera"`
	// CapNum 电容
	CapNum byte `json:"capNum"`
	// Cap
	Cap byte `json:"cap"`
	// RebornRfid 占领RFID满足激活条件
	RebornRfid byte `json:"rebornRfid"`
	// X uwb数据
	X float32 `json:"x"`
	// Y
	Y float32 `json:"y"`
	// Z
	Z float32 `json:"z"`
	// Compass
	Compass float32 `json:"compass"`
	// CoustomData1 学生自定义数据
	CoustomData1 float32 `json:"coustomData1"`
	// CoustomData2
	CoustomData2 float32 `json:"coustomData2"`
	// CoustomData3
	CoustomData3 float32 `json:"coustomData3"`
	// Mask
	Mask byte `json:"mask"`
	// RtsDmgData 扣血实时数据统计
	RtsDmgData S1BattleProto2022ClientHpChangeDetailVal `json:"rtsDmgData"`
	// CurrentPerformancePoint 当前技能点
	CurrentPerformancePoint int32 `json:"currentPerformancePoint"`
	// HpLevel 血量等级
	HpLevel int32 `json:"hpLevel"`
	// ChassisPowerLevel 底盘功率等级
	ChassisPowerLevel int32 `json:"chassisPowerLevel"`
	// ShootSpeedLevel 射速等级
	ShootSpeedLevel int32 `json:"shootSpeedLevel"`
}

type S1BattleProto2022ClientShielddata struct {
	// RedHasShield 红方是否有护盾
	RedHasShield byte `json:"redHasShield"`
	// BlueHasShield 蓝方是否有护盾
	BlueHasShield byte `json:"blueHasShield"`
	// RedShield 红方护盾值
	RedShield int32 `json:"redShield"`
	// BlueShield 蓝方护盾值
	BlueShield int32 `json:"blueShield"`
	// ShieldMax 护盾值上限
	ShieldMax int32 `json:"shieldMax"`
}

// S1BattleProto2022ClientSafetySync 给serverUI同步飞猫状态-SafetySync
type S1BattleProto2022ClientSafetySync struct {
	// Id 1红 2蓝
	Id byte `json:"id"`
	// Lift1Connect 升降机设备取消
	Lift1Connect byte `json:"lift1Connect"`
	// Lift2Connect 升降机设备取消
	Lift2Connect byte `json:"lift2Connect"`
	// FlycatConnect
	FlycatConnect byte `json:"flycatConnect"`
	// Lift1State 升降机设备取消
	Lift1State byte `json:"lift1State"`
	// Lift2State
	Lift2State byte `json:"lift2State"`
	// FlycatState
	FlycatState byte `json:"flycatState"`
	// Lift1Error 升降机设备取消
	Lift1Error byte `json:"lift1Error"`
	// Lift2Error
	Lift2Error byte `json:"lift2Error"`
	// CatError
	CatError byte `json:"catError"`
	// FlycatBattery
	FlycatBattery byte `json:"flycatBattery"`
}

// S1BattleProto2022ClientServerMapSync Server端小地图-ServerMapSync
type S1BattleProto2022ClientServerMapSync struct {
	// YawOffset
	YawOffset float32 `json:"yawOffset"`
	// AnchorMask
	AnchorMask byte `json:"anchorMask"`
	// RobotidListLen
	RobotidListLen byte `json:"robotidListLen"`
	// RobotidList
	RobotidList []S1BattleProto2022ClientRobotMapData `json:"robotidList"`
}

// S1BattleProto2022ClientServerUimessage 发送信息到ServerUI-ServerUIMessage
type S1BattleProto2022ClientServerUimessage struct {
	// Msg
	Msg string `json:"msg"`
}

type S1BattleProto2022ClientSeverAlertNotify struct {
	// YellowRateWarningCount 黄色告警累计个数
	YellowRateWarningCount int32 `json:"yellowRateWarningCount"`
	// RedRateWarningCount 红色告警累计个数
	RedRateWarningCount int32 `json:"redRateWarningCount"`
	// LastYellowTimestamp 最后一次黄色告警时间戳(秒)
	LastYellowTimestamp uint32 `json:"lastYellowTimestamp"`
	// LastRedTimestamp 最后一次红色告警时间戳(秒)
	LastRedTimestamp uint32 `json:"lastRedTimestamp"`
	// LastYellowGameLeftTime 最后一次黄色告警比赛剩余时间
	LastYellowGameLeftTime int32 `json:"lastYellowGameLeftTime"`
	// LastRedGameLeftTime 最后一次红色告警比赛剩余时间
	LastRedGameLeftTime int32 `json:"lastRedGameLeftTime"`
	// CurrAlertLevel 当前告警等级:0绿 1黄 2红
	CurrAlertLevel byte `json:"currAlertLevel"`
	// GetQingflowData 是否从呈现系统拉到轻流数据
	GetQingflowData byte `json:"getQingflowData"`
	// BalancedInfantryUltraLimitError 平衡步兵超限错误，0 无错误， 1 红方， 2 蓝方，3 双方
	BalancedInfantryUltraLimitError byte `json:"balancedInfantryUltraLimitError"`
}

type S1BattleProto2022ClientShowMessage struct {
	// Teamid
	Teamid byte `json:"teamid"`
	// MsgType
	MsgType byte `json:"msgType"`
	// MsgEnum
	MsgEnum uint32 `json:"msgEnum"`
	// MsgParams
	MsgParams string `json:"msgParams"`
}

// S1BattleProto2022ClientSiloEnvCtrReq SiloEnvCtrReq
type S1BattleProto2022ClientSiloEnvCtrReq struct {
	// Teamid 队伍ID 1 红 2蓝
	Teamid byte `json:"teamid"`
	// ControlCode 控制命令  0发射 1切换目标 2开门
	ControlCode byte `json:"controlCode"`
	// Target 目标 0前哨站 1基地
	Target byte `json:"target"`
}

// S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify 飞镖发射站闸门动作同步-BaseRobotDevStatusSyncData
type S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify struct {
	// TeamId 0红方 1蓝方
	TeamId byte `json:"teamId"`
	// DoorState 0开始开启，1开始关闭，2完全开启
	DoorState byte `json:"doorState"`
	// DoorOpenCnt 飞镖闸门开启次数
	DoorOpenCnt byte `json:"doorOpenCnt"`
}

type S1BattleProto2022ClientStatusSync struct {
	// ClientId
	ClientId int32 `json:"clientId"`
	// Status
	Status int32 `json:"status"`
}

// S1BattleProto2022ClientSupplySync 更新补给站信息-SupplySync
type S1BattleProto2022ClientSupplySync struct {
	// SupplyId 补给站ID
	SupplyId int32 `json:"supplyId"`
	// Connect 连接状态
	Connect int32 `json:"connect"`
	// State 0空闲 1释放 2装填 3未初始化
	State int32 `json:"state"`
	// UsableBullets 可用子弹
	UsableBullets int32 `json:"usableBullets"`
	// UsedBullets 已用子弹
	UsedBullets int32 `json:"usedBullets"`
	// ReadyBox 装填完毕的盒子
	ReadyBox int32 `json:"readyBox"`
	// FreedBox 释放的盒子
	FreedBox int32 `json:"freedBox"`
	// Timespan 下次加可用弹的剩余时间
	Timespan float32 `json:"timespan"`
	// NextaddBullets 下一次加可用弹数
	NextaddBullets int32 `json:"nextaddBullets"`
	// MakerRobotid 当前加弹的机器人id
	MakerRobotid int32 `json:"makerRobotid"`
	// MakeBullets 当前加弹的数量
	MakeBullets int32 `json:"makeBullets"`
	// ErrorCodeArrListLen
	ErrorCodeArrListLen int32 `json:"errorCodeArrListLen"`
	// ErrorCodeArrList 连接状态
	ErrorCodeArrList []int32 `json:"errorCodeArrList"`
	// RfidRobotIdListLen
	RfidRobotIdListLen int32 `json:"rfidRobotIdListLen"`
	// RfidRobotIdList 连接状态
	RfidRobotIdList []int32 `json:"rfidRobotIdList"`
}

type S1BattleProto2022ClientTalentDataAck struct {
	// Result
	Result int32 `json:"result"`
}

type S1BattleProto2022ClientTalentDataNotify struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// Data
	Data []性能体系_数据 `json:"data"`
	// IsBalance
	IsBalance byte `json:"isBalance"`
	// IsSemiAutomaticCtrl
	IsSemiAutomaticCtrl byte `json:"isSemiAutomaticCtrl"`
	// IsTempManualCtrl
	IsTempManualCtrl byte `json:"isTempManualCtrl"`
}

type S1BattleProto2022ClientTalentDataReq struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// 目标类型
	目标类型 string `json:""`
	// 性能类型
	性能类型 string `json:""`
	// 是否设置副武器
	是否设置副武器 byte `json:""`
}

// S1BattleProto2022ClientTeamInfo 客户端协议-TeamInfo
type S1BattleProto2022ClientTeamInfo struct {
	// TotalHp 战队总血量
	TotalHp uint32 `json:"totalHp"`
	// CurrentHp 战队当前血量
	CurrentHp uint32 `json:"currentHp"`
	// Warning1Count 一级警告
	Warning1Count uint16 `json:"warning1Count"`
	// Warning2Count 二级警告
	Warning2Count uint16 `json:"warning2Count"`
	// Warning3Count 三级警告
	Warning3Count uint16 `json:"warning3Count"`
	// TotalHurt 战队总伤害
	TotalHurt uint16 `json:"totalHurt"`
	// BigBulletAvailableNum 剩余可用大弹丸数量(英雄)
	BigBulletAvailableNum int16 `json:"bigBulletAvailableNum"`
	// BigBulletQuotaNum
	BigBulletQuotaNum int16 `json:"bigBulletQuotaNum"`
	// BigBulletCanBuyNum
	BigBulletCanBuyNum int16 `json:"bigBulletCanBuyNum"`
	// SmallBulletAvailableNum 剩余可用小弹丸数量(英雄+步兵)
	SmallBulletAvailableNum int16 `json:"smallBulletAvailableNum"`
	// SmallBulletQuotaNum
	SmallBulletQuotaNum int16 `json:"smallBulletQuotaNum"`
	// SmallBulletCanBuyNum
	SmallBulletCanBuyNum int16 `json:"smallBulletCanBuyNum"`
	// AirSupportNum 剩余可用空中支援次数
	AirSupportNum uint16 `json:"airSupportNum"`
	// BloodPack 血包数量
	BloodPack uint32 `json:"bloodPack"`
	// CenterPointEnergy 中心增益点能量
	CenterPointEnergy uint32 `json:"centerPointEnergy"`
	// TeamAvaliableCoinsNum 全队可用金币数量
	TeamAvaliableCoinsNum int32 `json:"teamAvaliableCoinsNum"`
	// TotalCoins 全队总经济
	TotalCoins int32 `json:"totalCoins"`
}

// S1BattleProto2022ClientVtmSpeedSync 图传速率同步-vtm_speed_sync
type S1BattleProto2022ClientVtmSpeedSync struct {
	// Clientid 客户端id
	Clientid int32 `json:"clientid"`
	// CurrentFreq 当前的图传的频道
	CurrentFreq int32 `json:"currentFreq"`
	// TxConnect 当前与发射端的连接状态
	TxConnect int32 `json:"txConnect"`
	// CurrentSpeedRate 当前的速率
	CurrentSpeedRate int32 `json:"currentSpeedRate"`
	// CurrentTxTemperature 当前的发射端温度
	CurrentTxTemperature int32 `json:"currentTxTemperature"`
}

type S1BattleProto2022ClientWarningNotify struct {
	// Type
	Type byte `json:"type"`
	// Team
	Team byte `json:"team"`
	// RobotId
	RobotId byte `json:"robotId"`
	// Leftcredit
	Leftcredit byte `json:"leftcredit"`
	// HpChangePercent
	HpChangePercent byte `json:"hpChangePercent"`
	// MaskTimeSelf
	MaskTimeSelf byte `json:"maskTimeSelf"`
	// MaskTimeOthers
	MaskTimeOthers byte `json:"maskTimeOthers"`
}

// S1BattleProto2022ClientWeaponFireNotify 子弹发射-WeaponFireNotify
type S1BattleProto2022ClientWeaponFireNotify struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// BulletType
	BulletType byte `json:"bulletType"`
	// Speed
	Speed float32 `json:"speed"`
	// Angle
	Angle float32 `json:"angle"`
}

// S1BattleProto2022ClientsFirstSync FirstSync登录同步包（首次更新数据，完成断线重连的状态更新）
type S1BattleProto2022ClientsFirstSync struct {
	// GameMode
	GameMode byte `json:"gameMode"`
	// BattleFirstData
	BattleFirstData S1BattleProto2022ClientBattleFirstData `json:"battleFirstData"`
	// RobotsNum
	RobotsNum int32 `json:"robotsNum"`
	// RobotsFirstData
	RobotsFirstData []S1BattleProto2022ClientRobotFirstData `json:"robotsFirstData"`
	// ClientNum
	ClientNum int32 `json:"clientNum"`
	// ClientsStatus
	ClientsStatus []byte `json:"clientsStatus"`
	// RobotsNum1
	RobotsNum1 int32 `json:"robotsNum1"`
	// RobotStatus
	RobotStatus []S1BattleProto2022ClientRobotStatusNotify `json:"robotStatus"`
	// CourtYawOffset
	CourtYawOffset float32 `json:"courtYawOffset"`
}

// S1BattleProto2022ClientsRobotModuleErrorInfo 同步机器人模块离线详细 RobotModuleErrorInfo
type S1BattleProto2022ClientsRobotModuleErrorInfo struct {
	// RobotId
	RobotId int32 `json:"robotId"`
	// PowerNum
	PowerNum byte `json:"powerNum"`
	// Power
	Power byte `json:"power"`
	// RfidNum
	RfidNum byte `json:"rfidNum"`
	// Rfid0
	Rfid0 byte `json:"rfid0"`
	// Rfid1
	Rfid1 byte `json:"rfid1"`
	// BloodNum
	BloodNum byte `json:"bloodNum"`
	// Blood0
	Blood0 byte `json:"blood0"`
	// Blood1
	Blood1 byte `json:"blood1"`
	// Blood2
	Blood2 byte `json:"blood2"`
	// IshasExShooter
	IshasExShooter byte `json:"ishasExShooter"`
	// SamllShooterDetectNum
	SamllShooterDetectNum byte `json:"samllShooterDetectNum"`
	// SamllshooterDetect0
	SamllshooterDetect0 byte `json:"samllshooterDetect0"`
	// SamllshooterDetect1
	SamllshooterDetect1 byte `json:"samllshooterDetect1"`
	// BigShooterDetectNum
	BigShooterDetectNum byte `json:"bigShooterDetectNum"`
	// BigshooterDetect
	BigshooterDetect byte `json:"bigshooterDetect"`
	// UwbNum
	UwbNum byte `json:"uwbNum"`
	// Uwb
	Uwb byte `json:"uwb"`
	// StrikeNum
	StrikeNum byte `json:"strikeNum"`
	// Strike0
	Strike0 byte `json:"strike0"`
	// Strike1
	Strike1 byte `json:"strike1"`
	// Strike2
	Strike2 byte `json:"strike2"`
	// Strike3
	Strike3 byte `json:"strike3"`
	// Strike4
	Strike4 byte `json:"strike4"`
	// Strike5
	Strike5 byte `json:"strike5"`
	// Strike6
	Strike6 byte `json:"strike6"`
	// Strike7
	Strike7 byte `json:"strike7"`
	// Strike8
	Strike8 byte `json:"strike8"`
	// Strike9
	Strike9 byte `json:"strike9"`
	// CameraNum
	CameraNum byte `json:"cameraNum"`
	// Camera
	Camera byte `json:"camera"`
	// CapNum
	CapNum byte `json:"capNum"`
	// Cap
	Cap byte `json:"cap"`
	// Spower
	Spower byte `json:"spower"`
	// Srfid0
	Srfid0 byte `json:"srfid0"`
	// Srfid1
	Srfid1 byte `json:"srfid1"`
	// Scamera
	Scamera byte `json:"scamera"`
	// Sblood0
	Sblood0 byte `json:"sblood0"`
	// Sblood1
	Sblood1 byte `json:"sblood1"`
	// SshooterDetect0
	SshooterDetect0 byte `json:"sshooterDetect0"`
	// SshooterDetect1
	SshooterDetect1 byte `json:"sshooterDetect1"`
	// SshooterDetect2
	SshooterDetect2 byte `json:"sshooterDetect2"`
	// Suwb
	Suwb byte `json:"suwb"`
	// Sarmor0
	Sarmor0 byte `json:"sarmor0"`
	// Sarmor1
	Sarmor1 byte `json:"sarmor1"`
	// Sarmor2
	Sarmor2 byte `json:"sarmor2"`
	// Sarmor3
	Sarmor3 byte `json:"sarmor3"`
	// Sarmor4
	Sarmor4 byte `json:"sarmor4"`
	// Sarmor5
	Sarmor5 byte `json:"sarmor5"`
	// Sarmor6
	Sarmor6 byte `json:"sarmor6"`
	// Sarmor7
	Sarmor7 byte `json:"sarmor7"`
	// Sarmor8
	Sarmor8 byte `json:"sarmor8"`
	// Sarmor9
	Sarmor9 byte `json:"sarmor9"`
	// Scap
	Scap byte `json:"scap"`
	// SmainController
	SmainController byte `json:"smainController"`
}

// S1BattleProto2022ClientsServerBaseSync 给客户端的数据（基础包同步（单次发送，非实时发送））ServerBaseSync
type S1BattleProto2022ClientsServerBaseSync struct {
	// RobotsSyncDatasLen 机器人的数量
	RobotsSyncDatasLen int32 `json:"robotsSyncDatasLen"`
	// RobotsBaseSyncData 每个机器人的数据
	RobotsBaseSyncData []S1BattleProto2022ClientRobotBaseDataSync `json:"robotsBaseSyncData"`
}

// S1BattleProto2022ClientsServerSync 同步给客户端的数据（100ms）ServerSync
type S1BattleProto2022ClientsServerSync struct {
	// PassTime 比赛Pass时间
	PassTime float32 `json:"passTime"`
	// LeftTime 剩余时间
	LeftTime float32 `json:"leftTime"`
	// CenterPointCoolTime 中心增益点冷却时间
	CenterPointCoolTime float32 `json:"centerPointCoolTime"`
	// TeamInfosLen 队伍信息
	TeamInfosLen int32 `json:"teamInfosLen"`
	// TeamInfos
	TeamInfos []S1BattleProto2022ClientTeamInfo `json:"teamInfos"`
	// RobotsSyncDatasLen 机器人的数量
	RobotsSyncDatasLen int32 `json:"robotsSyncDatasLen"`
	// RobotsSyncDatas 每个机器人的数据
	RobotsSyncDatas []S1BattleProto2022ClientRobotSyncData `json:"robotsSyncDatas"`
}

type S1BattleProto2022ConfigTabAck struct {
	// Magic
	Magic uint32 `json:"magic"`
	// RobotConfigVersion
	RobotConfigVersion byte `json:"robotConfigVersion"`
	// Color
	Color byte `json:"color"`
	// ModuleNum
	ModuleNum S1BattleProto2022ModuleNum `json:"moduleNum"`
	// HealthCalc
	HealthCalc S1BattleProto2022HealthCalc `json:"healthCalc"`
	// GameLimit
	GameLimit S1BattleProto2022GameLimit `json:"gameLimit"`
	// ArmorData
	ArmorData S1BattleProtot2022RobotArmorCfgData `json:"armorData"`
}

type S1BattleProto2022Energy2019AmorSelfCheck struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

type S1BattleProto2022Energy2019ArmorHitUpload struct {
	// Row
	Row byte `json:"row"`
	// ArmorId
	ArmorId byte `json:"armorId"`
	// HitType
	HitType uint16 `json:"hitType"`
}

type S1BattleProto2022Energy2019SetArmLight struct {
	// ArmorColor
	ArmorColor []uint32 `json:"armorColor"`
	// ArmColor
	ArmColor []uint32 `json:"armColor"`
	// Effect
	Effect []byte `json:"effect"`
	// ExtVar
	ExtVar []uint16 `json:"extVar"`
	// Row
	Row byte `json:"row"`
}

type S1BattleProto2022Energy2019StateReport struct {
	// State
	State byte `json:"state"`
	// Armors
	Armors []byte `json:"armors"`
	// Reserve
	Reserve byte `json:"reserve"`
	// RotateDeg
	RotateDeg uint16 `json:"rotateDeg"`
	// RotateSpeed
	RotateSpeed float32 `json:"rotateSpeed"`
}

type S1BattleProto2022Energy2020SetArmRotate struct {
	// CtrlMode
	CtrlMode byte `json:"ctrlMode"`
	// Direct
	Direct byte `json:"direct"`
}

// S1BattleProto2022EnergyReset 能量机关重置 energy_reset
type S1BattleProto2022EnergyReset struct {
	// Res
	Res byte `json:"res"`
}

// S1BattleProto2022EnergySetId 设置能量机关设置ID energy_set_id
type S1BattleProto2022EnergySetId struct {
	// Res
	Res byte `json:"res"`
}

// S1BattleProto2022EnergyStateChangeNotify 能量机关状态切换通知-RuneNotify
type S1BattleProto2022EnergyStateChangeNotify struct {
	// RuneId
	RuneId byte `json:"runeId"`
	// State
	State byte `json:"state"`
	// RingSum
	RingSum byte `json:"ringSum"`
	// AtkBuffVal 攻击力增益数值
	AtkBuffVal uint16 `json:"atkBuffVal"`
	// DefBuffVal 防御增益数值
	DefBuffVal byte `json:"defBuffVal"`
	// BeforeState
	BeforeState byte `json:"beforeState"`
}

// S1BattleProto2022EnergyStateSync 能量机关状态同步-RuneSync
type S1BattleProto2022EnergyStateSync struct {
	// RuneId
	RuneId byte `json:"runeId"`
	// Connect
	Connect byte `json:"connect"`
	// State
	State byte `json:"state"`
	// Round
	Round byte `json:"round"`
	// Time
	Time byte `json:"time"`
	// Error
	Error byte `json:"error"`
	// Armor0
	Armor0 byte `json:"armor0"`
	// Armor1
	Armor1 byte `json:"armor1"`
	// Armor2
	Armor2 byte `json:"armor2"`
	// Armor3
	Armor3 byte `json:"armor3"`
	// Armor4
	Armor4 byte `json:"armor4"`
	// Armor5
	Armor5 byte `json:"armor5"`
	// Motor
	Motor byte `json:"motor"`
	// AvailbleCountdown
	AvailbleCountdown byte `json:"availbleCountdown"`
}

// S1BattleProto2022EnerySetLogoLight 设置能量机关R的logo-rgb enery_set_logo_r_light
type S1BattleProto2022EnerySetLogoLight struct {
	// IconRColor
	IconRColor uint32 `json:"iconRcolor"`
	// Row
	Row byte `json:"row"`
}

// S1BattleProto2022EnvBaseShellControl 基地外壳控制协议，baseShellControl
type S1BattleProto2022EnvBaseShellControl struct {
	// BaseShellControl 0:自检重置   1.展开  2.闭合
	BaseShellControl byte `json:"baseShellControl"`
	// Rgb 颜色
	Rgb uint16 `json:"rgb"`
}

// S1BattleProto2022EnvDevicesDescripeAck 场地设备描述信息请求包-EnvDevicesDescripeAck
type S1BattleProto2022EnvDevicesDescripeAck struct {
	// Name
	Name S1BattleProto2022RobotDynamicUistr `json:"name"`
}

// S1BattleProto2022EnvDevicesDescripeReq 场地设备描述信息请求包-EnvDevicesDescripeReq
type S1BattleProto2022EnvDevicesDescripeReq struct {
	// Res
	Res byte `json:"res"`
}

type S1BattleProto2022EnvHeartBeatAck struct {
	// Time
	Time uint32 `json:"time"`
}

type S1BattleProto2022EnvHeartBeatReq struct {
	// Status
	Status byte `json:"status"`
}

// S1BattleProto2022ExchangeProClearOreRes 兑换站清矿结果-完成清矿动作后发送 v1重传
type S1BattleProto2022ExchangeProClearOreRes struct {
	// Res 1完成
	Res byte `json:"res"`
}

// S1BattleProto2022ExchangeProCtrCmd 机械臂运动控制-ExchangeProCtrCMD
type S1BattleProto2022ExchangeProCtrCmd struct {
	// Command
	Command byte `json:"command"`
	// Reserved1
	Reserved1 byte `json:"reserved1"`
	// Reserved2
	Reserved2 byte `json:"reserved2"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

// S1BattleProto2022ExchangeProCtrCmdack 机械臂运动控制Ack-ExchangeProCtrCMDAck
type S1BattleProto2022ExchangeProCtrCmdack struct {
	// ErrorCode
	ErrorCode byte `json:"errorCode"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

// S1BattleProto2022ExchangeProLidarInfo 机械臂运动控制-ExchangeProRealPosition
type S1BattleProto2022ExchangeProLidarInfo struct {
	// State 0无物体；1有物体
	State byte `json:"state"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

// S1BattleProto2022ExchangeProPosition 机械臂目标点位下发控制-ExchangeProPosition
type S1BattleProto2022ExchangeProPosition struct {
	// X
	X int32 `json:"x"`
	// Y
	Y int32 `json:"y"`
	// Z
	Z int32 `json:"z"`
	// Pitch
	Pitch int32 `json:"pitch"`
	// Roll
	Roll int32 `json:"roll"`
	// Yaw
	Yaw int32 `json:"yaw"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

// S1BattleProto2022ExchangeProRealPosition 机械臂运动控制-ExchangeProRealPosition
type S1BattleProto2022ExchangeProRealPosition struct {
	// X
	X int32 `json:"x"`
	// Y
	Y int32 `json:"y"`
	// Z
	Z int32 `json:"z"`
	// Pitch
	Pitch int32 `json:"pitch"`
	// Roll
	Roll int32 `json:"roll"`
	// Yaw
	Yaw int32 `json:"yaw"`
	// Seq 机械臂收到的服务器下发的最新位置包序号
	Seq uint32 `json:"seq"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

// S1BattleProto2022FlyCatStatus 飞猫状态-Safety_FlycatStatus
type S1BattleProto2022FlyCatStatus struct {
	// SysState 0正常 1警告 2错误
	SysState byte `json:"sysState"`
	// CtrlState 0服务器 1遥控器
	CtrlState byte `json:"ctrlState"`
	// WorkState
	WorkState byte `json:"workState"`
	// Battery
	Battery byte `json:"battery"`
	// SensorState 传感器状态，从低位到高位，1代表触发
	SensorState byte `json:"sensorState"`
}

type S1BattleProto2022GameLimit struct {
	// PowerThreshold
	PowerThreshold uint16 `json:"powerThreshold"`
	// PowerBuffer1
	PowerBuffer1 uint16 `json:"powerBuffer1"`
	// PowerBuffer2
	PowerBuffer2 uint16 `json:"powerBuffer2"`
	// PowerMaxhurt
	PowerMaxhurt uint16 `json:"powerMaxhurt"`
	// PowerHurtTabCount
	PowerHurtTabCount int32 `json:"powerHurtTabCount"`
	// PowerHurtTabs
	PowerHurtTabs []byte `json:"powerHurtTabs"`
	// ShooterLimitsCount
	ShooterLimitsCount int32 `json:"shooterLimitsCount"`
	// ShooterLimits
	ShooterLimits []S1BattleProto2022ShooterLimit `json:"shooterLimits"`
	// HeatLimitsCount
	HeatLimitsCount int32 `json:"heatLimitsCount"`
	// HeatLimits
	HeatLimits []S1BattleProto2022HeatLimit `json:"heatLimits"`
	// ShooterFreqLimitCount
	ShooterFreqLimitCount int32 `json:"shooterFreqLimitCount"`
	// FreqLimits
	FreqLimits []S1BattleProto2022ShooterFreqLimit `json:"freqLimits"`
}

// S1BattleProto2022GripperStateNotify 机械爪矿石信息-GripperStateNotify
type S1BattleProto2022GripperStateNotify struct {
	// IsConnect
	IsConnect byte `json:"isConnect"`
	// IsHasMineralsListLen
	IsHasMineralsListLen byte `json:"isHasMineralsListLen"`
	// IsHasMineralsList
	IsHasMineralsList []byte `json:"isHasMineralsList"`
	// ErrorsListLen
	ErrorsListLen byte `json:"errorsListLen"`
	// ErrorsList
	ErrorsList []byte `json:"errorsList"`
	// GrippersStatesListLen
	GrippersStatesListLen byte `json:"grippersStatesListLen"`
	// GripperStates
	GripperStates []byte `json:"gripperStates"`
}

type S1BattleProto2022HealthCalc struct {
	// HealthPointStart
	HealthPointStart uint16 `json:"healthPointStart"`
	// HealthPointFull
	HealthPointFull uint16 `json:"healthPointFull"`
	// BulletHurt
	BulletHurt uint16 `json:"bulletHurt"`
	// GolfHurt
	GolfHurt uint16 `json:"golfHurt"`
	// ImpactHurt
	ImpactHurt uint16 `json:"impactHurt"`
}

type S1BattleProto2022HeatLimit struct {
	// HeatSpdThreshold
	HeatSpdThreshold uint16 `json:"heatSpdThreshold"`
	// HeatSpdMaxHurt
	HeatSpdMaxHurt uint16 `json:"heatSpdMaxHurt"`
	// HeatHurtTabCount
	HeatHurtTabCount int32 `json:"heatHurtTabCount"`
	// HeatHurtTab
	HeatHurtTab []byte `json:"heatHurtTab"`
	// HeatCdFreq
	HeatCdFreq byte `json:"heatCdFreq"`
	// HeatCdStep
	HeatCdStep uint16 `json:"heatCdStep"`
}

// S1BattleProto2022IoctrCfgSet 设置DP的输出模式-IOCtrCfgSet
type S1BattleProto2022IoctrCfgSet struct {
	// IsDp2SpiMode DP2模式设置1为spi模式
	IsDp2SpiMode byte `json:"isDp2SpiMode"`
	// IsPwmDp2Io2 是否将DP2的3678口(bit2567)设置为PWM输出
	IsPwmDp2Io2 byte `json:"isPwmDp2Io2"`
	// IsPwmDp2Io5
	IsPwmDp2Io5 byte `json:"isPwmDp2Io5"`
	// IsPwmDp2Io6
	IsPwmDp2Io6 byte `json:"isPwmDp2Io6"`
	// IsPwmDp2Io7
	IsPwmDp2Io7 byte `json:"isPwmDp2Io7"`
	// Dp2PwmFrq 输出PWM的频率
	Dp2PwmFrq uint16 `json:"dp2PwmFrq"`
	// IsOutputDp1 是否设置DP1为输出
	IsOutputDp1 byte `json:"isOutputDp1"`
	// IsTriggerDp1Io0 DP1的IO口是否设置为触发式
	IsTriggerDp1Io0 byte `json:"isTriggerDp1Io0"`
	// IsTriggerDp1Io1
	IsTriggerDp1Io1 byte `json:"isTriggerDp1Io1"`
	// IsTriggerDp1Io2
	IsTriggerDp1Io2 byte `json:"isTriggerDp1Io2"`
	// IsTriggerDp1Io3
	IsTriggerDp1Io3 byte `json:"isTriggerDp1Io3"`
	// IsTriggerDp1Io4
	IsTriggerDp1Io4 byte `json:"isTriggerDp1Io4"`
	// IsTriggerDp1Io5
	IsTriggerDp1Io5 byte `json:"isTriggerDp1Io5"`
	// IsTriggerDp1Io6
	IsTriggerDp1Io6 byte `json:"isTriggerDp1Io6"`
	// IsTriggerDp1Io7
	IsTriggerDp1Io7 byte `json:"isTriggerDp1Io7"`
	// IsTriggerDp3Io0 DP3的IO口是否设置为触发式
	IsTriggerDp3Io0 byte `json:"isTriggerDp3Io0"`
	// IsTriggerDp3Io1
	IsTriggerDp3Io1 byte `json:"isTriggerDp3Io1"`
	// IsTriggerDp3Io2
	IsTriggerDp3Io2 byte `json:"isTriggerDp3Io2"`
	// IsTriggerDp3Io3
	IsTriggerDp3Io3 byte `json:"isTriggerDp3Io3"`
	// IsTriggerDp3Io4
	IsTriggerDp3Io4 byte `json:"isTriggerDp3Io4"`
	// IsTriggerDp3Io5
	IsTriggerDp3Io5 byte `json:"isTriggerDp3Io5"`
	// IsTriggerDp3Io6
	IsTriggerDp3Io6 byte `json:"isTriggerDp3Io6"`
	// IsTriggerDp3Io7
	IsTriggerDp3Io7 byte `json:"isTriggerDp3Io7"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrCfgSetAck IOCtrCfgSetAck
type S1BattleProto2022IoctrCfgSetAck struct {
	// ErrCode
	ErrCode byte `json:"errCode"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetActuator struct {
	// Actuator0
	Actuator0 byte `json:"actuator0"`
	// Actuator1
	Actuator1 byte `json:"actuator1"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetActuatorAck struct {
	// Actuator0
	Actuator0 byte `json:"actuator0"`
	// Actuator1
	Actuator1 byte `json:"actuator1"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetVal 设置IO控制板的输出值-IOCtrSetVal
type S1BattleProto2022IoctrSetVal struct {
	// Dp2OutputValsLen
	Dp2OutputValsLen int32 `json:"dp2OutputValsLen"`
	// Dp2OutputVal DP2输出值，PWM模式时，该值为占空比(万分比)；GPIO模式时，0低电平其余高电平
	Dp2OutputVal []uint16 `json:"dp2OutputVal"`
	// OutDp1Io0 设置DP1的8个IO口的值，只有当DP1_IsOutPut为1时设置为输出模式，下面的值设置才有效
	OutDp1Io0 byte `json:"outDp1Io0"`
	// OutDp1Io1
	OutDp1Io1 byte `json:"outDp1Io1"`
	// OutDp1Io2
	OutDp1Io2 byte `json:"outDp1Io2"`
	// OutDp1Io3
	OutDp1Io3 byte `json:"outDp1Io3"`
	// OutDp1Io4
	OutDp1Io4 byte `json:"outDp1Io4"`
	// OutDp1Io5
	OutDp1Io5 byte `json:"outDp1Io5"`
	// OutDp1Io6
	OutDp1Io6 byte `json:"outDp1Io6"`
	// OutDp1Io7
	OutDp1Io7 byte `json:"outDp1Io7"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetValAck IOCtrSetValAck
type S1BattleProto2022IoctrSetValAck struct {
	// ErrCode
	ErrCode byte `json:"errCode"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrState IO控制板上传的状态同步包，10Hz 只读-IOCtrState
type S1BattleProto2022IoctrState struct {
	// Dp1Gpio0 IO控制板的DP1口只有GPIO，该DP口Input/Output 复用
	Dp1Gpio0 byte `json:"dp1Gpio0"`
	// Dp1Gpio1
	Dp1Gpio1 byte `json:"dp1Gpio1"`
	// Dp1Gpio2
	Dp1Gpio2 byte `json:"dp1Gpio2"`
	// Dp1Gpio3
	Dp1Gpio3 byte `json:"dp1Gpio3"`
	// Dp1Gpio4
	Dp1Gpio4 byte `json:"dp1Gpio4"`
	// Dp1Gpio5
	Dp1Gpio5 byte `json:"dp1Gpio5"`
	// Dp1Gpio6
	Dp1Gpio6 byte `json:"dp1Gpio6"`
	// Dp1Gpio7
	Dp1Gpio7 byte `json:"dp1Gpio7"`
	// Dp3Adc1In0 DP3 模拟输入ADC，逻辑电平掩码，IO控制板会根据预设的阈值(500mv默认)将ADC转换成数字信号
	Dp3Adc1In0 byte `json:"dp3Adc1In0"`
	// Dp3Adc1In1
	Dp3Adc1In1 byte `json:"dp3Adc1In1"`
	// Dp3Adc1In2
	Dp3Adc1In2 byte `json:"dp3Adc1In2"`
	// Dp3Adc1In3
	Dp3Adc1In3 byte `json:"dp3Adc1In3"`
	// Dp3Adc1In4
	Dp3Adc1In4 byte `json:"dp3Adc1In4"`
	// Dp3Adc1In5
	Dp3Adc1In5 byte `json:"dp3Adc1In5"`
	// Dp3Adc1In6
	Dp3Adc1In6 byte `json:"dp3Adc1In6"`
	// Dp3Adc1In7
	Dp3Adc1In7 byte `json:"dp3Adc1In7"`
	// Dp3AdcValsListLen
	Dp3AdcValsListLen int32 `json:"dp3AdcValsListLen"`
	// Dp3AdcVals DP3 模拟输入ADC的值
	Dp3AdcVals []uint16 `json:"dp3AdcVals"`
	// IsPwmDp2Io0 DP2 为(PWM/GPIO共用)输出，8路输出中，只有3
	IsPwmDp2Io0 byte `json:"isPwmDp2Io0"`
	// IsPwmDp2Io1
	IsPwmDp2Io1 byte `json:"isPwmDp2Io1"`
	// IsPwmDp2Io2
	IsPwmDp2Io2 byte `json:"isPwmDp2Io2"`
	// IsPwmDp2Io3
	IsPwmDp2Io3 byte `json:"isPwmDp2Io3"`
	// IsPwmDp2Io4
	IsPwmDp2Io4 byte `json:"isPwmDp2Io4"`
	// IsPwmDp2Io5
	IsPwmDp2Io5 byte `json:"isPwmDp2Io5"`
	// IsPwmDp2Io6
	IsPwmDp2Io6 byte `json:"isPwmDp2Io6"`
	// IsPwmDp2Io7
	IsPwmDp2Io7 byte `json:"isPwmDp2Io7"`
	// Dp2OutValsListLen
	Dp2OutValsListLen int32 `json:"dp2OutValsListLen"`
	// Dp2OutVars dp2输出的值，该路被设置为PWM是，代表占空比(万分值10000代表100%)
	Dp2OutVars []uint16 `json:"dp2OutVars"`
	// Dp2PwmFrq
	Dp2PwmFrq uint16 `json:"dp2PwmFrq"`
	// Error
	Error byte `json:"error"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrTriggerVal IO控制板-触发式协议，用于触发时间小于100ms的设备(光电传感器)。IOCtrTriggerVal
type S1BattleProto2022IoctrTriggerVal struct {
	// Dp1InputValBefore 用byte代表8个IO，每一位代表一个IO口的val。通过before和after比较得出变化的IO口
	Dp1InputValBefore byte `json:"dp1InputValBefore"`
	// Dp3InputValBefore
	Dp3InputValBefore byte `json:"dp3InputValBefore"`
	// Dp1InputValAfter
	Dp1InputValAfter byte `json:"dp1InputValAfter"`
	// Dp3InputValAfter
	Dp3InputValAfter byte `json:"dp3InputValAfter"`
	// SysTickMs
	SysTickMs uint32 `json:"sysTickMs"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022MapClickInfoNotify 客户端小地图标点-MapClickInfoNotify
type S1BattleProto2022MapClickInfoNotify struct {
	// TeamId 发送消息的队伍ID
	TeamId byte `json:"teamId"`
	// IsSendAll 1.发送给所有客户端  0：发送给指定客户端
	IsSendAll byte `json:"isSendAll"`
	// RobotidListLen 发送给制定客户端的机器人ID
	RobotidListLen byte `json:"robotidListLen"`
	// RobotidList
	RobotidList []byte `json:"robotidList"`
	// Mode 标记类型 1：攻击  2:防御  3：警戒  4.自定义
	Mode byte `json:"mode"`
	// EnemyId 标定的敌人ID
	EnemyId byte `json:"enemyId"`
	// Ascii 自定义图标的ASCII 码
	Ascii byte `json:"ascii"`
	// Type 标记模式 1：地图  2:敌方机器人
	Type byte `json:"type"`
	// ScreenX 屏幕坐标x
	ScreenX uint16 `json:"screenX"`
	// ScreenY
	ScreenY uint16 `json:"screenY"`
	// MapX 地图坐标x
	MapX float32 `json:"mapX"`
	// MapY
	MapY float32 `json:"mapY"`
}

type S1BattleProto2022ModuleNum struct {
	// SmallStrikeNum
	SmallStrikeNum byte `json:"smallStrikeNum"`
	// BigStrikeNum
	BigStrikeNum byte `json:"bigStrikeNum"`
	// LightBarNum
	LightBarNum byte `json:"lightBarNum"`
	// PowerNum
	PowerNum byte `json:"powerNum"`
	// SmallShootNum
	SmallShootNum byte `json:"smallShootNum"`
	// BigShootNum
	BigShootNum byte `json:"bigShootNum"`
	// CameraNum
	CameraNum byte `json:"cameraNum"`
	// RfidNum
	RfidNum byte `json:"rfidNum"`
	// UwbNum
	UwbNum byte `json:"uwbNum"`
	// SmallArmorImp
	SmallArmorImp byte `json:"smallArmorImp"`
	// BigArmorImp
	BigArmorImp byte `json:"bigArmorImp"`
	// LightBrdImp
	LightBrdImp byte `json:"lightBrdImp"`
	// PowerBrdImp
	PowerBrdImp byte `json:"powerBrdImp"`
	// SmallShotImp
	SmallShotImp byte `json:"smallShotImp"`
	// BigShotImp
	BigShotImp byte `json:"bigShotImp"`
	// CameraImp
	CameraImp byte `json:"cameraImp"`
	// RfidImp
	RfidImp byte `json:"rfidImp"`
	// UwbImp
	UwbImp byte `json:"uwbImp"`
	// CapNum
	CapNum byte `json:"capNum"`
	// CapImp
	CapImp byte `json:"capImp"`
	// ExtTypeAnum
	ExtTypeAnum byte `json:"extTypeAnum"`
	// ExtTypeAimp
	ExtTypeAimp byte `json:"extTypeAimp"`
	// ExtTypeBnum
	ExtTypeBnum byte `json:"extTypeBnum"`
	// ExtTypeBimp
	ExtTypeBimp byte `json:"extTypeBimp"`
	// ExtTypeCnum
	ExtTypeCnum byte `json:"extTypeCnum"`
	// ExtTypeCimp
	ExtTypeCimp byte `json:"extTypeCimp"`
	// Res
	Res byte `json:"res"`
}

type S1BattleProto2022PowerSwitchState struct {
	// Chassis
	Chassis byte `json:"chassis"`
	// Gimbal
	Gimbal byte `json:"gimbal"`
	// Shooter
	Shooter byte `json:"shooter"`
}

// S1BattleProto2022QueryRobotInfo 查询版本-QueryRobotInfo
type S1BattleProto2022QueryRobotInfo struct {
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

// S1BattleProto2022QueryRobotInfoResult 查询版本回复-QueryRobotInfoResult
type S1BattleProto2022QueryRobotInfoResult struct {
	// LoaderVersion
	LoaderVersion uint32 `json:"loaderVersion"`
	// AppVersion
	AppVersion uint32 `json:"appVersion"`
	// DeviceIdListLen
	DeviceIdListLen int32 `json:"deviceIdListLen"`
	// DeviceIdList
	DeviceIdList []int32 `json:"deviceIdList"`
	// Reserved
	Reserved uint32 `json:"reserved"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
}

type S1BattleProto2022ReLoginDataSync struct {
	// RoleS0Id
	RoleS0Id uint64 `json:"roleS0Id"`
	// BuffCount
	BuffCount byte `json:"buffCount"`
	// BuffUids
	BuffUids []uint64 `json:"buffUids"`
	// BuffTids
	BuffTids []uint32 `json:"buffTids"`
	// BuffLevels
	BuffLevels []uint32 `json:"buffLevels"`
	// BuffVisibles
	BuffVisibles []byte `json:"buffVisibles"`
	// BuffMaxTime
	BuffMaxTime []float32 `json:"buffMaxTime"`
	// BuffLeftTime
	BuffLeftTime []float32 `json:"buffLeftTime"`
}

// S1BattleProto2022RobotBaseDevCtlCmdAck 基地控制命令的ACK-BaseRobotDevCtlCmdAck
type S1BattleProto2022RobotBaseDevCtlCmdAck struct {
	// Placeholder
	Placeholder byte `json:"placeholder"`
}

// S1BattleProto2022RobotCheckinTimestamp 检录时间戳-flash_timestamp_t
type S1BattleProto2022RobotCheckinTimestamp struct {
	// Magic
	Magic uint32 `json:"magic"`
	// DataVer
	DataVer byte `json:"dataVer"`
	// Timestamp
	Timestamp uint32 `json:"timestamp"`
}

type S1BattleProto2022RobotDataSync struct {
	// CurHp 当前血量
	CurHp int32 `json:"curHp"`
	// MaxHp 总血量
	MaxHp int32 `json:"maxHp"`
	// Level 当前等级
	Level byte `json:"level"`
	// LightEffectMask 灯效掩码
	LightEffectMask uint32 `json:"lightEffectMask"`
	// ShooterHeatsCount 发射机构热量
	ShooterHeatsCount int32 `json:"shooterHeatsCount"`
	// ShooterHeats
	ShooterHeats []float32 `json:"shooterHeats"`
	// ShooterHeatThresCount 发射机构热量上限
	ShooterHeatThresCount int32 `json:"shooterHeatThresCount"`
	// ShooterHeatThres
	ShooterHeatThres []float32 `json:"shooterHeatThres"`
	// CurHeatCoolCount 发射机构热量
	CurHeatCoolCount int32 `json:"curHeatCoolCount"`
	// CurHeatCool
	CurHeatCool []float32 `json:"curHeatCool"`
	// IsGuardAlive 哨兵是否存活
	IsGuardAlive uint16 `json:"isGuardAlive"`
	// InventedShieldValue 基地的护盾值
	InventedShieldValue uint16 `json:"inventedShieldValue"`
	// FixPower 最大功率上限
	FixPower uint16 `json:"fixPower"`
	// FixSamllShootSpeed1 0号小枪口最大小枪口射速上限
	FixSamllShootSpeed1 uint16 `json:"fixSamllShootSpeed1"`
	// FixBigShootSpeed 最大大枪口射速上限
	FixBigShootSpeed uint16 `json:"fixBigShootSpeed"`
	// FixSamllShootSpeed2 1号小枪最大小枪口射速上限
	FixSamllShootSpeed2 uint16 `json:"fixSamllShootSpeed2"`
}

// S1BattleProto2022RobotDynamicUistr 场地设备名字由服务器配置并下发给主控-RobotDynamicUIStr
type S1BattleProto2022RobotDynamicUistr struct {
	// EnvDeviceXnameEn
	EnvDeviceXnameEn string `json:"envDeviceXnameEn"`
	// EnvDeviceXnameZh
	EnvDeviceXnameZh string `json:"envDeviceXnameZh"`
	// EnvDeviceYnameEn
	EnvDeviceYnameEn string `json:"envDeviceYnameEn"`
	// EnvDeviceYnameZh
	EnvDeviceYnameZh string `json:"envDeviceYnameZh"`
	// EnvDeviceZnameEn
	EnvDeviceZnameEn string `json:"envDeviceZnameEn"`
	// EnvDeviceZnameZh
	EnvDeviceZnameZh string `json:"envDeviceZnameZh"`
	// ExtModuleAnameEn
	ExtModuleAnameEn string `json:"extModuleAnameEn"`
	// ExtModuleAnameZh
	ExtModuleAnameZh string `json:"extModuleAnameZh"`
	// ExtModuleBnameEn
	ExtModuleBnameEn string `json:"extModuleBnameEn"`
	// ExtModuleBnameZh
	ExtModuleBnameZh string `json:"extModuleBnameZh"`
	// ExtModuleCnameEn
	ExtModuleCnameEn string `json:"extModuleCnameEn"`
	// ExtModuleCnameZh
	ExtModuleCnameZh string `json:"extModuleCnameZh"`
}

// S1BattleProto2022RobotGetCheckCapAck 电容检测-检录电容的ack-GetCheckCapAck
type S1BattleProto2022RobotGetCheckCapAck struct {
	// CheckedCap
	CheckedCap float32 `json:"checkedCap"`
	// VoltageFirst
	VoltageFirst float32 `json:"voltageFirst"`
	// VoltageFinal
	VoltageFinal float32 `json:"voltageFinal"`
	// MeasureTime
	MeasureTime float32 `json:"measureTime"`
	// OutputQEnergy
	OutputQEnergy float32 `json:"outputQenergy"`
	// RecharingQEnergy
	RecharingQEnergy float32 `json:"recharingQenergy"`
	// Date
	Date uint32 `json:"date"`
	// Time
	Time uint32 `json:"time"`
}

// S1BattleProto2022RobotGetCheckCapReq 电容检测-获取检录电容值req-GetCheckCapReq
type S1BattleProto2022RobotGetCheckCapReq struct {
	// Res
	Res byte `json:"res"`
}

type S1BattleProto2022RobotHited struct {
	// OnHitType 子弹类型
	OnHitType byte `json:"onHitType"`
	// AttackSpeed 子弹射速
	AttackSpeed uint16 `json:"attackSpeed"`
	// ArmorNumber 装甲ID
	ArmorNumber byte `json:"armorNumber"`
	// Press 子弹压力
	Press float32 `json:"press"`
	// TimeStamp 时间戳
	TimeStamp uint64 `json:"timeStamp"`
}

// S1BattleProto2022RobotIrcheckReq 机器人红外检测-RobotIRCheckReq
type S1BattleProto2022RobotIrcheckReq struct {
	// Check
	Check byte `json:"check"`
}

type S1BattleProto2022RobotInitCfgTab struct {
	// Ext17MmSpeed
	Ext17MmSpeed byte `json:"ext17MmSpeed"`
	// ResCount
	ResCount int32 `json:"resCount"`
	// Res
	Res []byte `json:"res"`
}

// S1BattleProto2022RobotMeasureStartReq 电容检测-开始检测req-MeasureStartReq
type S1BattleProto2022RobotMeasureStartReq struct {
	// MaxPrriod
	MaxPrriod uint32 `json:"maxPrriod"`
	// PushPrriod
	PushPrriod uint32 `json:"pushPrriod"`
}

// S1BattleProto2022RobotMeasureStartRsp 电容检测-开始检测rsp-MeasureStartRsp
type S1BattleProto2022RobotMeasureStartRsp struct {
	// Type
	Type byte `json:"type"`
	// Id
	Id byte `json:"id"`
	// Status
	Status byte `json:"status"`
}

// S1BattleProto2022RobotMeasureStopReq 电容检测-结束检测req-MeasureStopReq
type S1BattleProto2022RobotMeasureStopReq struct {
	// Res
	Res byte `json:"res"`
}

// S1BattleProto2022RobotMeasureStopRsp 电容检测-结束检测rsp-MeasureStopRsp
type S1BattleProto2022RobotMeasureStopRsp struct {
	// Res
	Res byte `json:"res"`
}

// S1BattleProto2022RobotMoudleLife 查询装甲寿命-RobotMoudleLife
type S1BattleProto2022RobotMoudleLife struct {
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
	// AckModuleId
	AckModuleId byte `json:"ackModuleId"`
	// AckModuleType
	AckModuleType byte `json:"ackModuleType"`
	// Status
	Status string `json:"status"`
}

type S1BattleProto2022RobotNewHeartBeatAck struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

// S1BattleProto2022RobotNewHeartBeatReq new heart beat
type S1BattleProto2022RobotNewHeartBeatReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

// S1BattleProto2022RobotNtpServerIpAck 回复NTP服务器地址-NtpServerIpAck
type S1BattleProto2022RobotNtpServerIpAck struct {
	// Ip
	Ip uint32 `json:"ip"`
}

// S1BattleProto2022RobotNtpServerIpReq NTP服务器IP地址-NtpServerIpReq
type S1BattleProto2022RobotNtpServerIpReq struct {
	// Res
	Res byte `json:"res"`
}

// S1BattleProto2022RobotPushCapinfo 电容检测-检测结果-PushCAPInfo
type S1BattleProto2022RobotPushCapinfo struct {
	// VoltageFirst
	VoltageFirst float32 `json:"voltageFirst"`
	// VoltageFinal
	VoltageFinal float32 `json:"voltageFinal"`
	// MeasureTime
	MeasureTime float32 `json:"measureTime"`
	// OutputQEnergy
	OutputQEnergy float32 `json:"outputQenergy"`
	// RecharingQEnergy
	RecharingQEnergy float32 `json:"recharingQenergy"`
	// MeasureCap
	MeasureCap float32 `json:"measureCap"`
	// CheckCap
	CheckCap float32 `json:"checkCap"`
}

// S1BattleProto2022RobotPushCaprtinfo 电容检测-实时推送-PushCAPRTInfo
type S1BattleProto2022RobotPushCaprtinfo struct {
	// Voltage
	Voltage float32 `json:"voltage"`
	// OutputCurrent
	OutputCurrent float32 `json:"outputCurrent"`
	// RechargingCurrent
	RechargingCurrent float32 `json:"rechargingCurrent"`
	// VoltageFirst
	VoltageFirst float32 `json:"voltageFirst"`
	// MeasureTime
	MeasureTime float32 `json:"measureTime"`
	// MeasureCap
	MeasureCap float32 `json:"measureCap"`
	// RecharingQEnergy
	RecharingQEnergy float32 `json:"recharingQenergy"`
	// OutputQEnergy
	OutputQEnergy float32 `json:"outputQenergy"`
}

// S1BattleProto2022RobotPushSensorInfo 电容检测-传感器数值-PushSensorInfo
type S1BattleProto2022RobotPushSensorInfo struct {
	// Voltage
	Voltage float32 `json:"voltage"`
	// OutputCurrent
	OutputCurrent float32 `json:"outputCurrent"`
	// RechargingCurrent
	RechargingCurrent float32 `json:"rechargingCurrent"`
	// Status
	Status byte `json:"status"`
}

// S1BattleProto2022RobotResurgenceNotify 客户端下发信息 135-RobotResurgenceNotify
type S1BattleProto2022RobotResurgenceNotify struct {
	// RobotId
	RobotId byte `json:"robotId"`
}

type S1BattleProto2022RobotStatus struct {
	// RobotStatusData
	RobotStatusData S1BattleProto2022RobotStatusData `json:"robotStatusData"`
}

type S1BattleProto2022RobotStatusData struct {
	// Hp
	Hp uint16 `json:"hp"`
	// Voltage
	Voltage float32 `json:"voltage"`
	// Current
	Current float32 `json:"current"`
	// ChassisPower
	ChassisPower float32 `json:"chassisPower"`
	// PowerBuffer
	PowerBuffer float32 `json:"powerBuffer"`
	// GimbalVoltage
	GimbalVoltage float32 `json:"gimbalVoltage"`
	// GimbalPower
	GimbalPower float32 `json:"gimbalPower"`
	// ShooterVoltage
	ShooterVoltage float32 `json:"shooterVoltage"`
	// ShooterPower
	ShooterPower float32 `json:"shooterPower"`
	// Rssi
	Rssi byte `json:"rssi"`
	// HwId
	HwId uint32 `json:"hwId"`
	// ShooterPitch
	ShooterPitch int16 `json:"shooterPitch"`
	// ShooterRoll
	ShooterRoll int16 `json:"shooterRoll"`
	// ShooterYaw
	ShooterYaw int16 `json:"shooterYaw"`
	// Status
	Status S1BattleProto2022RobotSystemStatus `json:"status"`
	// Uwb
	Uwb S1BattleProto2022UwbData `json:"uwb"`
	// PowerState
	PowerState S1BattleProto2022PowerSwitchState `json:"powerState"`
	// TimeStamp
	TimeStamp uint64 `json:"timeStamp"`
}

type S1BattleProto2022RobotSystemStatus struct {
	// Power
	Power byte `json:"power"`
	// Rfid0
	Rfid0 byte `json:"rfid0"`
	// Rfid1
	Rfid1 byte `json:"rfid1"`
	// Camera
	Camera byte `json:"camera"`
	// Blood0
	Blood0 byte `json:"blood0"`
	// Blood1
	Blood1 byte `json:"blood1"`
	// ShooterDetect0
	ShooterDetect0 byte `json:"shooterDetect0"`
	// ShooterDetect1
	ShooterDetect1 byte `json:"shooterDetect1"`
	// ShooterDetect2
	ShooterDetect2 byte `json:"shooterDetect2"`
	// Uwb
	Uwb byte `json:"uwb"`
	// Armor0
	Armor0 byte `json:"armor0"`
	// Armor1
	Armor1 byte `json:"armor1"`
	// Armor2
	Armor2 byte `json:"armor2"`
	// Armor3
	Armor3 byte `json:"armor3"`
	// Armor4
	Armor4 byte `json:"armor4"`
	// Armor5
	Armor5 byte `json:"armor5"`
	// Armor6
	Armor6 byte `json:"armor6"`
	// Armor7
	Armor7 byte `json:"armor7"`
	// Armor8
	Armor8 byte `json:"armor8"`
	// Armor9
	Armor9 byte `json:"armor9"`
	// Cap
	Cap byte `json:"cap"`
	// ExtA0
	ExtA0 byte `json:"extA0"`
	// ExtA1
	ExtA1 byte `json:"extA1"`
	// ExtA2
	ExtA2 byte `json:"extA2"`
	// ExtA3
	ExtA3 byte `json:"extA3"`
	// ExtA4
	ExtA4 byte `json:"extA4"`
	// ExtA5
	ExtA5 byte `json:"extA5"`
	// ExtA6
	ExtA6 byte `json:"extA6"`
	// ExtB0
	ExtB0 byte `json:"extB0"`
	// ExtB1
	ExtB1 byte `json:"extB1"`
	// ExtB2
	ExtB2 byte `json:"extB2"`
	// ExtC0
	ExtC0 byte `json:"extC0"`
	// ExtC1
	ExtC1 byte `json:"extC1"`
	// ExtC2
	ExtC2 byte `json:"extC2"`
	// ReservedCount
	ReservedCount int32 `json:"reservedCount"`
	// ReservedList
	ReservedList []byte `json:"reservedList"`
}

type S1BattleProto2022RobotVtmsettingFull struct {
	// ListLen
	ListLen byte `json:"listLen"`
	// Tds1
	Tds1 []byte `json:"tds1"`
	// Tds2
	Tds2 []byte `json:"tds2"`
	// Tds3
	Tds3 []byte `json:"tds3"`
}

type S1BattleProto2022RobotWeaponFire struct {
	// ShooterId 枪口ID  小枪口：0 2 大枪口：1
	ShooterId byte `json:"shooterId"`
	// BulletType 子弹类型
	BulletType byte `json:"bulletType"`
	// Speed 射速
	Speed float32 `json:"speed"`
	// Frequency 射频
	Frequency float32 `json:"frequency"`
	// Yaw 射击yaw角度
	Yaw float32 `json:"yaw"`
	// Pitch 射击pitch角度
	Pitch float32 `json:"pitch"`
	// RollAngle 射击roll角度
	RollAngle float32 `json:"rollAngle"`
	// TimeStamp 时间戳
	TimeStamp uint64 `json:"timeStamp"`
}

type S1BattleProto2022RobotsDataSync struct {
	// Progress 当前比赛状态
	Progress byte `json:"progress"`
	// TimeRemain 剩余时间，从比赛开始倒计时
	TimeRemain uint16 `json:"timeRemain"`
	// Utc UTC时间
	Utc uint32 `json:"utc"`
	// RobotData 机器人状态包
	RobotData S1BattleProto2022RobotDataSync `json:"robotData"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022S2CPowerStateNotify 客户端协议-S2C_设备上电状态_Notify
type S1BattleProto2022S2CPowerStateNotify struct {
	// ClientId
	ClientId byte `json:"clientId"`
	// ChassisPowerState
	ChassisPowerState int32 `json:"chassisPowerState"`
	// GimbalPowerState
	GimbalPowerState int32 `json:"gimbalPowerState"`
	// ShooterPowerState
	ShooterPowerState int32 `json:"shooterPowerState"`
}

// S1BattleProto2022SendDartRobotStatus 飞镖发射架的状态给飞镖发射站硬件-FaSheJiaStatus
type S1BattleProto2022SendDartRobotStatus struct {
	// Connectstatus
	Connectstatus byte `json:"connectstatus"`
}

type S1BattleProto2022SetFlyCatLight struct {
	// Lightcode
	Lightcode uint32 `json:"lightcode"`
}

type S1BattleProto2022SetFlyCatLightAck struct {
	// RecvedLightcode
	RecvedLightcode uint32 `json:"recvedLightcode"`
}

// S1BattleProto2022SetFlyCatStatus 设置飞猫状态-Safety_SetFlycatStatus
type S1BattleProto2022SetFlyCatStatus struct {
	// WorkState
	WorkState byte `json:"workState"`
}

// S1BattleProto2022SetFlyCatStatusAck 设置飞猫状态的Ack-Safety_SetFlycatStatus
type S1BattleProto2022SetFlyCatStatusAck struct {
	// RecvedWorkState
	RecvedWorkState byte `json:"recvedWorkState"`
}

// S1BattleProto2022SetSupplyStateReq 设置补给站-SetSupplyStateReq
type S1BattleProto2022SetSupplyStateReq struct {
	// ReportTime ms 0为不自动报告
	ReportTime uint16 `json:"reportTime"`
	// ReportReq 此协议是否回复Ack
	ReportReq byte `json:"reportReq"`
}

type S1BattleProto2022ShooterFreqLimit struct {
	// ShooterFreqThreshold
	ShooterFreqThreshold float32 `json:"shooterFreqThreshold"`
	// ShooterFreqMaxHurt
	ShooterFreqMaxHurt uint16 `json:"shooterFreqMaxHurt"`
	// ShooterFreqHurtTabCount
	ShooterFreqHurtTabCount int32 `json:"shooterFreqHurtTabCount"`
	// ShooterFreqHurtTab
	ShooterFreqHurtTab []byte `json:"shooterFreqHurtTab"`
}

type S1BattleProto2022ShooterLimit struct {
	// ShooterSpdThreshold
	ShooterSpdThreshold float32 `json:"shooterSpdThreshold"`
	// ShooterSpdMaxHurt
	ShooterSpdMaxHurt uint16 `json:"shooterSpdMaxHurt"`
	// ShooterSqdHurtTabCount
	ShooterSqdHurtTabCount int32 `json:"shooterSqdHurtTabCount"`
	// ShooterSqdHurtTab
	ShooterSqdHurtTab []byte `json:"shooterSqdHurtTab"`
}

type S1BattleProto2022SiloCtrLedData struct {
	// R
	R byte `json:"r"`
	// G
	G byte `json:"g"`
	// B
	B byte `json:"b"`
	// A
	A byte `json:"a"`
}

type S1BattleProto2022SiloDevStatus struct {
	// Online 是否在线  1：在线  2.离线
	Online byte `json:"online"`
	// DoorStatus 门的状态  0:打开  1：关闭 2:运行中
	DoorStatus byte `json:"doorStatus"`
	// FloorStatus 地板状态  0:被拉出或者划出   1.推回状态
	FloorStatus byte `json:"floorStatus"`
	// Speed 飞镖的速度 * 10   单位m/s
	Speed uint16 `json:"speed"`
	// Angle 飞镖发射的角度
	Angle uint16 `json:"angle"`
	// Target 飞镖发射的瞄准目标
	Target byte `json:"target"`
	// CountdownTime 发射井关闭倒计时
	CountdownTime float32 `json:"countdownTime"`
	// ShootCheckCountdownTime 检测导弹窗口了倒计时
	ShootCheckCountdownTime float32 `json:"shootCheckCountdownTime"`
	// OpenedTimeCount 每局比赛只能打开两次
	OpenedTimeCount int32 `json:"openedTimeCount"`
	// IsDetectionWindows 是否在导弹检测窗口期，导弹发出去5s后，不计算伤害
	IsDetectionWindows byte `json:"isDetectionWindows"`
	// MissileCount 发射的飞镖数量
	MissileCount int32 `json:"missileCount"`
	// IsCanOpen 发射井是否可以开启
	IsCanOpen byte `json:"isCanOpen"`
	// SiloCooldown 发射井冷却倒计时
	SiloCooldown float32 `json:"siloCooldown"`
	// SiloErrorCode 发射井错误码 0:OK 1:error
	SiloErrorCode byte `json:"siloErrorCode"`
	// MissileHitCount 飞镖命中数
	MissileHitCount byte `json:"missileHitCount"`
	// ErrorGateSensorUp 闸门上部传感器中某一个损坏
	ErrorGateSensorUp byte `json:"errorGateSensorUp"`
	// ErrorGateSensorDown 闸门下部传感器中某一个损坏
	ErrorGateSensorDown byte `json:"errorGateSensorDown"`
	// ErrorBaseBoardSensor 滑台末端传感器中某一个损坏
	ErrorBaseBoardSensor byte `json:"errorBaseBoardSensor"`
	// ErrorGateSensorBoth 闸门上下两端传感器同时触发
	ErrorGateSensorBoth byte `json:"errorGateSensorBoth"`
	// ErrorMotorBroken 电机损坏，具体为长时间接收不到电机的反馈信号
	ErrorMotorBroken byte `json:"errorMotorBroken"`
	// ErrorBrakeConstNo 制动器无法制动，具体为该刹车的时候无法刹车
	ErrorBrakeConstNo byte `json:"errorBrakeConstNo"`
	// ErrorBrakeConstYes 制动器一直制动，具体为电机因制动器而堵转、无法运动
	ErrorBrakeConstYes byte `json:"errorBrakeConstYes"`
	// ErrorFacingObstacle 闸门上下运动过程中，因电机输出力矩过大，判定为遇到障碍物，可能压到了学生设备或人体
	ErrorFacingObstacle byte `json:"errorFacingObstacle"`
	// ErrorEmergencyStop 急停开关按下
	ErrorEmergencyStop byte `json:"errorEmergencyStop"`
	// ErrorMotorOverHeat 电机高温警告 > 50 0正常 1异常
	ErrorMotorOverHeat byte `json:"errorMotorOverHeat"`
	// ErrorCloseOverTime 闸门关闭超时
	ErrorCloseOverTime byte `json:"errorCloseOverTime"`
	// ErrorOpenOverTime 闸门开启超时
	ErrorOpenOverTime byte `json:"errorOpenOverTime"`
}

type S1BattleProto2022SiloDevStatusSyncData struct {
	// SiloDevStatusSyncDataListLen
	SiloDevStatusSyncDataListLen int32 `json:"siloDevStatusSyncDataListLen"`
	// SiloDevStatusDataList 连接状态
	SiloDevStatusDataList []S1BattleProto2022SiloDevStatus `json:"siloDevStatusDataList"`
}

// S1BattleProto2022StuAreialEnergy 空中机器人能量状态0x205 AirPlane 10Hz-StuAreialEnergy
type S1BattleProto2022StuAreialEnergy struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// RemainTime 剩余可攻击时间30s递减至0
	RemainTime byte `json:"remainTime"`
}

// S1BattleProto2022StuClientRecvInfo 客户端接收信息 0x0305-ClientRecvInfo
type S1BattleProto2022StuClientRecvInfo struct {
	// TargetRobotId
	TargetRobotId uint16 `json:"targetRobotId"`
	// TargetPosX
	TargetPosX float32 `json:"targetPosX"`
	// TargetPosY
	TargetPosY float32 `json:"targetPosY"`
}

// S1BattleProto2022StuClientSendCmd 客户端下发信息 0x0303-ClientSendCMD
type S1BattleProto2022StuClientSendCmd struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// TargetPosX
	TargetPosX float32 `json:"targetPosX"`
	// TargetPosY
	TargetPosY float32 `json:"targetPosY"`
	// TargetPosZ
	TargetPosZ float32 `json:"targetPosZ"`
	// CmdKeyboard
	CmdKeyboard byte `json:"cmdKeyboard"`
	// TargetRobotId
	TargetRobotId uint16 `json:"targetRobotId"`
}

// S1BattleProto2022StuCommunication 车间通信协议0x301-Communication
type S1BattleProto2022StuCommunication struct {
	// DataId
	DataId uint16 `json:"dataId"`
	// SenderId
	SenderId uint16 `json:"senderId"`
	// ReceiverId
	ReceiverId uint16 `json:"receiverId"`
}

// S1BattleProto2022StuCommunicationByteData 车间通信协议-CommunicationByteData
type S1BattleProto2022StuCommunicationByteData struct {
	// Cmd 0x301车间通信r->r，0x305雷达信息r->client
	Cmd uint16 `json:"cmd"`
	// ByteDataListLen
	ByteDataListLen byte `json:"byteDataListLen"`
	// ByteDataList
	ByteDataList []byte `json:"byteDataList"`
	// S0HeaderBodyLen 机器人上发的协议不一定是128长，记录下来S0包头里带的包长度
	S0HeaderBodyLen uint16 `json:"s0HeaderBodyLen"`
	// Dataid 解析出的学生信息
	Dataid uint16 `json:"dataid"`
	// Senderid
	Senderid uint16 `json:"senderid"`
	// Receiverid
	Receiverid uint16 `json:"receiverid"`
}

// S1BattleProto2022StuCustomControlData 自定义控制器交互数据 0x0302-StuCustomControlData
type S1BattleProto2022StuCustomControlData struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// ListLen
	ListLen byte `json:"listLen"`
	// Data
	Data []byte `json:"data"`
}

// S1BattleProto2022StuEnvStatus 比赛场地数据0x0101 SelfTeam Trigger-StuEnvStatus
type S1BattleProto2022StuEnvStatus struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// Num1AddBloodPointStatus
	Num1AddBloodPointStatus uint32 `json:"num1AddBloodPointStatus"`
	// Num2AddBloodPointStatus
	Num2AddBloodPointStatus uint32 `json:"num2AddBloodPointStatus"`
	// Num3AddBloodPointStatus
	Num3AddBloodPointStatus uint32 `json:"num3AddBloodPointStatus"`
	// RuneRfidStatus
	RuneRfidStatus uint32 `json:"runeRfidStatus"`
	// SmallRuneStatus
	SmallRuneStatus uint32 `json:"smallRuneStatus"`
	// BigRuneStatus
	BigRuneStatus uint32 `json:"bigRuneStatus"`
	// RingUpland 己方环形高地R2占领状态，1已占领
	RingUpland uint32 `json:"ringUpland"`
	// TrapezoidR3 己方梯形高地R3占领状态，1已占领
	TrapezoidR3 uint32 `json:"trapezoidR3"`
	// TrapezoidR4 己方梯形高地R4占领状态，1已占领
	TrapezoidR4 uint32 `json:"trapezoidR4"`
	// BaseShieldStatus 己方基地护盾状态，1有护盾血量 0无虚拟护盾血量
	BaseShieldStatus uint32 `json:"baseShieldStatus"`
	// OutpostStatus 己方前哨站状态 1存活 0被击毁
	OutpostStatus uint32 `json:"outpostStatus"`
}

// S1BattleProto2022StuGameResult 比赛结果数据0x0002 AllRobot Trigger-StuGameResult
type S1BattleProto2022StuGameResult struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// Winner 0tie 1red 2blue
	Winner byte `json:"winner"`
}

// S1BattleProto2022StuGameStatus 比赛状态信息0x0001 AllRobot 1Hz-StuGameStatus
type S1BattleProto2022StuGameStatus struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// GameMode 比赛类型
	GameMode byte `json:"gameMode"`
	// GameStatus 比赛阶段
	GameStatus byte `json:"gameStatus"`
	// RemainTime 当前阶段剩余时间
	RemainTime uint16 `json:"remainTime"`
	// SyncTimeStamp 机器人接收到该指令的精确unix时间(ms)
	SyncTimeStamp uint64 `json:"syncTimeStamp"`
}

// S1BattleProto2022StuIcrabuffDebuffZoneStatus 人工智能挑战赛加成与惩罚区状态0x00005 AllRobot 1Hz-StuICRA_Buff_Debuff_Zone_Status
type S1BattleProto2022StuIcrabuffDebuffZoneStatus struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// F1Status F1区域激活状态
	F1Status byte `json:"f1Status"`
	// F1StatusInfo F1区域状态信息
	F1StatusInfo byte `json:"f1StatusInfo"`
	// F2Status F2区域激活状态
	F2Status byte `json:"f2Status"`
	// F2StatusInfo F2区域状态信息
	F2StatusInfo byte `json:"f2StatusInfo"`
	// F3Status
	F3Status byte `json:"f3Status"`
	// F3StatusInfo
	F3StatusInfo byte `json:"f3StatusInfo"`
	// F4Status
	F4Status byte `json:"f4Status"`
	// F4StatusInfo
	F4StatusInfo byte `json:"f4StatusInfo"`
	// F5Status
	F5Status byte `json:"f5Status"`
	// F5StatusInfo
	F5StatusInfo byte `json:"f5StatusInfo"`
	// F6Status
	F6Status byte `json:"f6Status"`
	// F6StatusInfo
	F6StatusInfo byte `json:"f6StatusInfo"`
}

// S1BattleProto2022StuLeftBullet 飞机和哨兵剩余子弹0x0208 空中or哨兵or英雄 10Hz-StuLeftBullet
type S1BattleProto2022StuLeftBullet struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// SmallAvaliableBulletsNum
	SmallAvaliableBulletsNum uint16 `json:"smallAvaliableBulletsNum"`
	// BigAvaliableBulletsNum
	BigAvaliableBulletsNum uint16 `json:"bigAvaliableBulletsNum"`
	// LeftCoinsNum 剩余金币数量
	LeftCoinsNum uint16 `json:"leftCoinsNum"`
}

// S1BattleProto2022StuMissileRemainingTime 飞镖发射口倒计时0x0105 SelfTeam 1Hz-StuMissileRemainingTime
type S1BattleProto2022StuMissileRemainingTime struct {
	// Cmdid
	Cmdid uint16 `json:"cmdid"`
	// Time 15s倒计时
	Time byte `json:"time"`
}

// S1BattleProto2022StuRfidstatus RFID状态0x0209 SingleRobot 1Hz-StuRFIDStatus
type S1BattleProto2022StuRfidstatus struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// BaseArea
	BaseArea uint32 `json:"baseArea"`
	// UplandArea
	UplandArea uint32 `json:"uplandArea"`
	// RuneArea
	RuneArea uint32 `json:"runeArea"`
	// FlyArea
	FlyArea uint32 `json:"flyArea"`
	// OutpostArea
	OutpostArea uint32 `json:"outpostArea"`
	// IslandArea
	IslandArea uint32 `json:"islandArea"`
	// HomeArea
	HomeArea uint32 `json:"homeArea"`
	// SapperRfid
	SapperRfid uint32 `json:"sapperRfid"`
}

// S1BattleProto2022StuRobotBuff 机器人Buff包 0x0204 单个机器人 1Hz-StuRobotBuff
type S1BattleProto2022StuRobotBuff struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// HealBuff 加血BUFF
	HealBuff byte `json:"healBuff"`
	// CoolBuff 冷却Buff
	CoolBuff byte `json:"coolBuff"`
	// DefenceBuff 防御Buff
	DefenceBuff byte `json:"defenceBuff"`
	// AttackBuff 攻击Buff
	AttackBuff byte `json:"attackBuff"`
}

// S1BattleProto2022StuRobotCurrentHp 机器人血量数据0x0003 AllRobot 1Hz-StuRobotCurrentHP
type S1BattleProto2022StuRobotCurrentHp struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// RedHero
	RedHero uint16 `json:"redHero"`
	// RedSapper
	RedSapper uint16 `json:"redSapper"`
	// RedInfantry1
	RedInfantry1 uint16 `json:"redInfantry1"`
	// RedInfantry2
	RedInfantry2 uint16 `json:"redInfantry2"`
	// RedInfantry3
	RedInfantry3 uint16 `json:"redInfantry3"`
	// RedGuard
	RedGuard uint16 `json:"redGuard"`
	// RedOutpost
	RedOutpost uint16 `json:"redOutpost"`
	// RedBase
	RedBase uint16 `json:"redBase"`
	// BlueHero
	BlueHero uint16 `json:"blueHero"`
	// BlueSapper
	BlueSapper uint16 `json:"blueSapper"`
	// BlueInfantry1
	BlueInfantry1 uint16 `json:"blueInfantry1"`
	// BlueInfantry2
	BlueInfantry2 uint16 `json:"blueInfantry2"`
	// BlueInfantry3
	BlueInfantry3 uint16 `json:"blueInfantry3"`
	// BlueGuard
	BlueGuard uint16 `json:"blueGuard"`
	// BlueOutpost
	BlueOutpost uint16 `json:"blueOutpost"`
	// BlueBase
	BlueBase uint16 `json:"blueBase"`
}

// S1BattleProto2022StuRobotHurt 伤害状态0x206 SingleRobot Trigger-StuRobotHurt
type S1BattleProto2022StuRobotHurt struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// Armor
	Armor byte `json:"armor"`
	// Hurttype 伤害类型 0x0打击 0x1模块离线 0x02超射速 0x03超热量 0x04超功率 0x05撞击
	Hurttype byte `json:"hurttype"`
}

// S1BattleProto2022StuSiloInfo 飞镖机器人客户端指令数据0x020A 飞镖站 10Hz-StuSiloInfo
type S1BattleProto2022StuSiloInfo struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// LaunchOpeningStatus 当前飞镖发射口状态 1关闭 2运行中 0开启
	LaunchOpeningStatus byte `json:"launchOpeningStatus"`
	// AttackTarget 打击目标 0前哨站 1基地
	AttackTarget byte `json:"attackTarget"`
	// TargetChangeTime 切换打击目标时的比赛剩余时间(s) 默认0
	TargetChangeTime uint16 `json:"targetChangeTime"`
	// OperateLaunchCmdTime 最后一次发起发射飞镖的比赛剩余时间(s) 默认0
	OperateLaunchCmdTime uint16 `json:"operateLaunchCmdTime"`
}

// S1BattleProto2022StuSupplyStatus 补给站状态包0x0102 SelfTeam Trigger-StuSupplyStatus
type S1BattleProto2022StuSupplyStatus struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// SupplyEnvId 1号补给口 2号补给口
	SupplyEnvId byte `json:"supplyEnvId"`
	// SupplyRobotId 0无 1\101英雄 2\102工程 3 4 5\103 104 105步兵
	SupplyRobotId byte `json:"supplyRobotId"`
	// SupplyEnvStatus 0关闭 1准备中 2子弹下落
	SupplyEnvStatus byte `json:"supplyEnvStatus"`
	// BulletsNum 子弹数量 50 100 150 200
	BulletsNum byte `json:"bulletsNum"`
}

// S1BattleProto2022StuWarning 裁判警告命令0x0104 SelfTeam Trigger-StuWarning
type S1BattleProto2022StuWarning struct {
	// Cmd
	Cmd uint16 `json:"cmd"`
	// WarningLevel 1黄牌 2红牌 3判负
	WarningLevel byte `json:"warningLevel"`
	// WarningRobot 判负时：0    Other：RobotId
	WarningRobot byte `json:"warningRobot"`
}

// S1BattleProto2022SupplyClearScaleAck 补给站清扫子弹称回复-SupplyClearScaleAck
type S1BattleProto2022SupplyClearScaleAck struct {
	// Placeholder
	Placeholder byte `json:"placeholder"`
}

// S1BattleProto2022SupplyClearScaleReq 补给站清扫子弹称请求-SupplyClearScaleReq
type S1BattleProto2022SupplyClearScaleReq struct {
	// Placeholder
	Placeholder byte `json:"placeholder"`
}

// S1BattleProto2022SupplyExportAck 释放子弹回复-SupplyExportAck
type S1BattleProto2022SupplyExportAck struct {
	// LastState 上一次状态
	LastState byte `json:"lastState"`
	// NowState 这一次状态
	NowState byte `json:"nowState"`
}

// S1BattleProto2022SupplyExportReq 打开供弹口请求-SupplyExportReq
type S1BattleProto2022SupplyExportReq struct {
	// ExportEnable 1:开启 0:关闭
	ExportEnable byte `json:"exportEnable"`
}

// S1BattleProto2022SupplyFreedAck 释放子弹回复-SupplyFreedAck
type S1BattleProto2022SupplyFreedAck struct {
	// IsAccept 是否接受了请求
	IsAccept byte `json:"isAccept"`
}

// S1BattleProto2022SupplyFreedReq 释放子弹请求-SupplyReloadReq
type S1BattleProto2022SupplyFreedReq struct {
	// BoxFreedNum 释放子弹盒数量
	BoxFreedNum byte `json:"boxFreedNum"`
}

// S1BattleProto2022SupplyReloadAck 装填盒子回复-SupplyReloadAck
type S1BattleProto2022SupplyReloadAck struct {
	// IsAccept 是否接受了请求
	IsAccept byte `json:"isAccept"`
	// Action 开始装填 or 结束装填
	Action byte `json:"action"`
}

// S1BattleProto2022SupplyReloadReq 装填盒子请求-SupplyReloadReq
type S1BattleProto2022SupplyReloadReq struct {
	// ReloadEnable 装填开关 1：开启装填 0：关闭装填
	ReloadEnable byte `json:"reloadEnable"`
}

// S1BattleProto2022SupplyReport 补给站状态包 默认100ms-SupplyReport
type S1BattleProto2022SupplyReport struct {
	// BulletGear0 拨弹轮0号故障
	BulletGear0 byte `json:"bulletGear0"`
	// BulletGear1  拨弹轮1号故障
	BulletGear1 byte `json:"bulletGear1"`
	// Loadselector0 拨弹轮通道选择器0号故障
	Loadselector0 byte `json:"loadselector0"`
	// Box0 供弹盒0号故障
	Box0 byte `json:"box0"`
	// Box1 供弹盒1号故障
	Box1 byte `json:"box1"`
	// Box2 供弹盒2号故障
	Box2 byte `json:"box2"`
	// Relase 挡板故障
	Relase byte `json:"relase"`
	// Scales 子弹称故障
	Scales byte `json:"scales"`
	// State supply_state
	State byte `json:"state"`
	// BoxReadyNum 供弹盒准备数量，可立即进行释放的数量
	BoxReadyNum byte `json:"boxReadyNum"`
	// BoxFreedNum 供弹盒已经被释放的数量
	BoxFreedNum byte `json:"boxFreedNum"`
}

// S1BattleProto2022SupplyResetReq 重置补给站（用于清空错误状态）-SupplyResetReq
type S1BattleProto2022SupplyResetReq struct {
	// Placeholder
	Placeholder byte `json:"placeholder"`
}

// S1BattleProto2022SupplyRestartReq 重启补给站（用于清空子弹）-SupplyRestartReq
type S1BattleProto2022SupplyRestartReq struct {
	// Placeholder
	Placeholder byte `json:"placeholder"`
}

// S1BattleProto2022TestTimeDelayDownComplete Server到主控 - 结束包
type S1BattleProto2022TestTimeDelayDownComplete struct {
	// Count
	Count uint32 `json:"count"`
}

// S1BattleProto2022TestTimeDelayDownData 数据包延时测试协议(Server发起)
type S1BattleProto2022TestTimeDelayDownData struct {
	// Seq
	Seq uint32 `json:"seq"`
	// CurTime
	CurTime int64 `json:"curTime"`
	// WifiTime
	WifiTime uint32 `json:"wifiTime"`
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Datas
	Datas []byte `json:"datas"`
}

// S1BattleProto2022TestTimeDelayDownRecord Server到主控 - 主控回报表
type S1BattleProto2022TestTimeDelayDownRecord struct {
	// GlobalDelay
	GlobalDelay uint32 `json:"globalDelay"`
	// UartDelay
	UartDelay uint32 `json:"uartDelay"`
	// UartLossCnt
	UartLossCnt byte `json:"uartLossCnt"`
	// WifiLossCnt
	WifiLossCnt byte `json:"wifiLossCnt"`
	// WifiLossTabCount
	WifiLossTabCount int32 `json:"wifiLossTabCount"`
	// WifiLossTabS
	WifiLossTabS []uint16 `json:"wifiLossTabS"`
}

type S1BattleProto2022UwbData struct {
	// X
	X float32 `json:"x"`
	// Y
	Y float32 `json:"y"`
	// Z
	Z float32 `json:"z"`
	// Compass
	Compass float32 `json:"compass"`
	// AnchorMask
	AnchorMask byte `json:"anchorMask"`
	// WifiEn
	WifiEn byte `json:"wifiEn"`
	// Rsv0
	Rsv0 byte `json:"rsv0"`
	// Rsv1
	Rsv1 byte `json:"rsv1"`
}

type S1BattleProto2022VtmSpeedSync struct {
	// Clientid
	Clientid int32 `json:"clientid"`
	// CurrentFreq 当前的图传的频道
	CurrentFreq int32 `json:"currentFreq"`
	// TxConnect 当前与发射端的连接状态
	TxConnect int32 `json:"txConnect"`
	// CurrentSpeedRate 当前的速率
	CurrentSpeedRate int32 `json:"currentSpeedRate"`
	// CurrentCountryCode 当前的国家码
	CurrentCountryCode int32 `json:"currentCountryCode"`
}

// S1BattleProto2022IoctrCfgReq -IOCtrCfgReq
type S1BattleProto2022IoctrCfgReq struct {
	// Res 保留字段
	Res byte `json:"res"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrRmMotorsStatus 电机回传数据
type S1BattleProto2022IoctrRmMotorsStatus struct {
	// RealPositionListLen
	RealPositionListLen int32 `json:"realPositionListLen"`
	// RealPositions 转子机械角度
	RealPositions []int64 `json:"realPositions"`
	// RealSpeedListLen
	RealSpeedListLen int32 `json:"realSpeedListLen"`
	// RealSpeeds 转子转速
	RealSpeeds []int16 `json:"realSpeeds"`
	// RealTemperatureListLen
	RealTemperatureListLen int32 `json:"realTemperatureListLen"`
	// RealTemperatures 电机温度
	RealTemperatures []byte `json:"realTemperatures"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetLightsRgbType1 type1：设置灯条颜色，同时设置多个灯为一种颜色-IOCtrSetLightsRgbType1
type S1BattleProto2022IoctrSetLightsRgbType1 struct {
	// StartLedIndex 控制起始LED灯珠的序号，0开始
	StartLedIndex uint16 `json:"startLedIndex"`
	// LedNum 控制的led灯珠数量，1开始
	LedNum uint16 `json:"ledNum"`
	// Color
	Color S1BattleProto2022LedColor `json:"color"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetLightsRgbType1Ack IOCtrSetLightsRgbType1Ack
type S1BattleProto2022IoctrSetLightsRgbType1Ack struct {
	// ErrCode 0正常
	ErrCode byte `json:"errCode"`
	// SeqAck
	SeqAck uint32 `json:"seqAck"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetLightsRgbType2 type2：设置灯条颜色，分别设置多个灯为不同颜色-IOCtrSetLightsRgbType2
type S1BattleProto2022IoctrSetLightsRgbType2 struct {
	// StartLedIndex 控制起始LED灯珠的序号，0开始
	StartLedIndex uint16 `json:"startLedIndex"`
	// LedNum 控制的led灯珠数量，1开始
	LedNum uint16 `json:"ledNum"`
	// ListLen
	ListLen int32 `json:"listLen"`
	// LedsColors
	LedsColors []S1BattleProto2022LedColor `json:"ledsColors"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetLightsRgbType2Ack struct {
	// ErrCode
	ErrCode byte `json:"errCode"`
	// SeqAck
	SeqAck uint32 `json:"seqAck"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetLightsRgbType3 type3 设置灯条特殊模式，0常亮，1呼吸-IOCtrSetLightsRgbType3
type S1BattleProto2022IoctrSetLightsRgbType3 struct {
	// LightEffect 0常亮，1呼吸
	LightEffect byte `json:"lightEffect"`
	// TimeSpan ms
	TimeSpan uint16 `json:"timeSpan"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetLightsRgbType3Ack IOCtrSetLightsRgbType3Ack
type S1BattleProto2022IoctrSetLightsRgbType3Ack struct {
	// ErrCode 0正常
	ErrCode byte `json:"errCode"`
	// SeqAck
	SeqAck uint32 `json:"seqAck"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetRmMotorsCfg 初始化电机控制参数，0-3控制can_id 0x1ff电机， 4-7控制can_id 0x200和0x2ff电机-IOCtrSetRmMotorsCfg
type S1BattleProto2022IoctrSetRmMotorsCfg struct {
	// SpeedPidListLen
	SpeedPidListLen int32 `json:"speedPidListLen"`
	// SpeedPids 速度环的pid
	SpeedPids []S1BattleProto2022SpeedModePid `json:"speedPids"`
	// PositionPidListLen
	PositionPidListLen int32 `json:"positionPidListLen"`
	// PositionPids 位置环pid
	PositionPids []S1BattleProto2022PositionModePid `json:"positionPids"`
	// TransRatioListLen
	TransRatioListLen int32 `json:"transRatioListLen"`
	// TransRatios 实际传动比，正为同向，负为反向
	TransRatios []float32 `json:"transRatios"`
	// IsStatusReturnListLen
	IsStatusReturnListLen int32 `json:"isStatusReturnListLen"`
	// IsStatusReturn 是否回传电机数据
	IsStatusReturn []byte `json:"isStatusReturn"`
	// Seq
	Seq uint32 `json:"seq"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022IoctrSetRmMotorsCfgAck struct {
	// ErrCode
	ErrCode byte `json:"errCode"`
	// SeqAck
	SeqAck uint32 `json:"seqAck"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

// S1BattleProto2022IoctrSetRmMotorsMoveVal 电机运行参数，根据控制模式，该期望值为速度或位置-IOCtrSetRmMotorsMoveVal
type S1BattleProto2022IoctrSetRmMotorsMoveVal struct {
	// LoopModeListLen
	LoopModeListLen int32 `json:"loopModeListLen"`
	// LoopMode 控制模式，0：默认无输出1:实际速度闭环 2:转子速度闭环 3:转子位置闭环 4:转子增量位置闭环 5:实际位置闭环 6:实际轴角度闭环 7:紧急停止
	LoopMode []byte `json:"loopMode"`
	// Len
	Len int32 `json:"len"`
	// ExpectVal
	ExpectVal []int64 `json:"expectVal"`
	// ModuleId
	ModuleId byte `json:"moduleId"`
	// ModuleType
	ModuleType byte `json:"moduleType"`
}

type S1BattleProto2022LedColor struct {
	// R
	R byte `json:"r"`
	// G
	G byte `json:"g"`
	// B
	B byte `json:"b"`
}

// S1BattleProto2022PositionModePid 位置环PID参数结构体-position_mode_pid_t
type S1BattleProto2022PositionModePid struct {
	// Kp 比例
	Kp float32 `json:"kp"`
	// Ki 积分
	Ki float32 `json:"ki"`
	// Kd 微分
	Kd float32 `json:"kd"`
	// NoResponse 不响应区间
	NoResponse byte `json:"noResponse"`
	// IntergralLimit 积分限幅
	IntergralLimit uint16 `json:"intergralLimit"`
	// MaxSpeed 最大速度
	MaxSpeed uint16 `json:"maxSpeed"`
}

type S1BattleProto2022SiloCtrlDoor struct {
	// Cmd
	Cmd byte `json:"cmd"`
}

type S1BattleProto2022SiloStatus struct {
	// DoorStatus
	DoorStatus byte `json:"doorStatus"`
	// FloorStatus
	FloorStatus byte `json:"floorStatus"`
	// Errorcode
	Errorcode byte `json:"errorcode"`
}

// S1BattleProto2022SpeedModePid 速度环PID参数结构体-speed_mode_pid_t
type S1BattleProto2022SpeedModePid struct {
	// Kp 比例
	Kp float32 `json:"kp"`
	// Ki 积分
	Ki float32 `json:"ki"`
	// Kd 微分
	Kd float32 `json:"kd"`
	// MaxOutput 最大输出值
	MaxOutput uint16 `json:"maxOutput"`
}

type S1BattleProto2023ClientChangeTokenCmd struct {
	// Len
	Len uint16 `json:"len"`
	// NewToken
	NewToken string `json:"newToken"`
}

type S1BattleProto2023ClientGameSystemInfoNotify struct {
	// TokenLen
	TokenLen uint16 `json:"tokenLen"`
	// CurMatchToken
	CurMatchToken string `json:"curMatchToken"`
}

// S1BattleProto2023ClientGuardInPatrolInfoSync 哨兵巡逻区信息同步
type S1BattleProto2023ClientGuardInPatrolInfoSync struct {
	// RobotId 机器人ID
	RobotId byte `json:"robotId"`
	// IsCurrentInPatrolArea 当前是否在巡逻区
	IsCurrentInPatrolArea byte `json:"isCurrentInPatrolArea"`
	// IsLeavePatrolAreaTimeout 是否离开巡逻区超时
	IsLeavePatrolAreaTimeout byte `json:"isLeavePatrolAreaTimeout"`
	// LeavePatrolAreaLeftTime 离开巡逻区倒计时
	LeavePatrolAreaLeftTime float32 `json:"leavePatrolAreaLeftTime"`
	// IsCountDownActive 倒计时是否生效
	IsCountDownActive byte `json:"isCountDownActive"`
}

type S1BattleProto2023ClientModuleVersionData struct {
	// MoudleId 模块id
	MoudleId byte `json:"moudleId"`
	// MoudleType 模块类型
	MoudleType byte `json:"moudleType"`
	// MoudleSubType 模块子类型
	MoudleSubType byte `json:"moudleSubType"`
	// NewMoudleVersion 模块最新版本号
	NewMoudleVersion string `json:"newMoudleVersion"`
	// CurMoudleVersion 模块当前版本号
	CurMoudleVersion string `json:"curMoudleVersion"`
	// VersionState 模块当前状态 0：查询无果 1：已是最新 2：低于最新 3: 高于最新 4:服务器解析出错
	VersionState byte `json:"versionState"`
	// MoudleIsImportant 是否重要模块
	MoudleIsImportant byte `json:"moudleIsImportant"`
}

type S1BattleProto2023ClientModuleVersionState struct {
	// RobotId
	RobotId byte `json:"robotId"`
	// ModuleNum
	ModuleNum byte `json:"moduleNum"`
	// ModuleNumMax
	ModuleNumMax byte `json:"moduleNumMax"`
	// ModuleVersion
	ModuleVersion []S1BattleProto2023ClientModuleVersionData `json:"moduleVersion"`
}

type S1BattleProto2023ClientPenaltyInfo struct {
	// RobotId 被罚机器人id
	RobotId byte `json:"robotId"`
	// PenaltyType 判罚类型
	PenaltyType byte `json:"penaltyType"`
	// PenaltyReason 被罚原因
	PenaltyReason string `json:"penaltyReason"`
}

type S1BattleProto2023ClientPenaltyTableInfos struct {
	// UploadType 判罚上传类型 0 自动，1手动
	UploadType byte `json:"uploadType"`
	// InfosLen
	InfosLen byte `json:"infosLen"`
	// Infos 判罚信息表
	Infos []S1BattleProto2023ClientPenaltyInfo `json:"infos"`
}

type S1BattleProto2023ClientPenaltyTableInfosAck struct {
	// State 判罚信息表接收状态
	State byte `json:"state"`
	// UploadType 判罚上传类型 0 自动，1手动
	UploadType byte `json:"uploadType"`
}

// S1BattleProto2023ClientRobotBattleInfo 单个机器人战斗信息-RobotBattleInfo
type S1BattleProto2023ClientRobotBattleInfo struct {
	// Robots0Id 机器人s0id号
	Robots0Id byte `json:"robots0Id"`
	// BattleState 是否脱战
	BattleState byte `json:"battleState"`
	// CanRemoteHeal 是否能远程补血
	CanRemoteHeal byte `json:"canRemoteHeal"`
	// CanRemoteExchangeSmallBullet 是否能远程兑换小弹丸
	CanRemoteExchangeSmallBullet byte `json:"canRemoteExchangeSmallBullet"`
	// CanRemoteExchangeBigBullet 是否能远程兑换大弹丸
	CanRemoteExchangeBigBullet byte `json:"canRemoteExchangeBigBullet"`
	// RemaindBuyRevivalCount 剩余买活次数
	RemaindBuyRevivalCount byte `json:"remaindBuyRevivalCount"`
	// BuyRevivalPrice 当前买活所需花费
	BuyRevivalPrice int32 `json:"buyRevivalPrice"`
	// BattleStateCountDown 脱战倒计时
	BattleStateCountDown float32 `json:"battleStateCountDown"`
	// IsOnHpPoint 是否在补血点IsRobotOnHppoint1Area
	IsOnHpPoint byte `json:"isOnHpPoint"`
	// ChassisPowerCtrl 底盘控电方式 默认0，0:规则 1:裁判控制
	ChassisPowerCtrl byte `json:"chassisPowerCtrl"`
	// GimbalPowerCtrl 云台控电方式 默认0，0:规则 1:裁判控制
	GimbalPowerCtrl byte `json:"gimbalPowerCtrl"`
	// ShooterPowerCtrl 枪口控电方式 默认0，0:规则 1:裁判控制
	ShooterPowerCtrl byte `json:"shooterPowerCtrl"`
}

// S1BattleProto2023ClientServerRobotsBattleInfoSync 机器人战斗相关信息同步客户端-RobotsBattleInfoSync
type S1BattleProto2023ClientServerRobotsBattleInfoSync struct {
	// RobotsNum
	RobotsNum byte `json:"robotsNum"`
	// RobotsBattleInfoSyncDatas 每个机器人的战斗信息
	RobotsBattleInfoSyncDatas []S1BattleProto2023ClientRobotBattleInfo `json:"robotsBattleInfoSyncDatas"`
}

// S1BattleProto2023ClientTeamSupplyInfoSync 队伍补给信息同步-TeamSupplyInfoSync
type S1BattleProto2023ClientTeamSupplyInfoSync struct {
	// Teamid 队伍ID
	Teamid byte `json:"teamid"`
	// ExchangeSmallBulletPackageCount 远程兑换小弹丸兑换剩余次数
	ExchangeSmallBulletPackageCount byte `json:"exchangeSmallBulletPackageCount"`
	// ExchangeBigBulletPackageCount 远程兑换大弹丸兑换剩余次数
	ExchangeBigBulletPackageCount byte `json:"exchangeBigBulletPackageCount"`
	// ExchangeRemoteHealCount 远程兑换回血剩余次数
	ExchangeRemoteHealCount byte `json:"exchangeRemoteHealCount"`
	// SentryRemainRevivalCount 哨兵复活剩余次数
	SentryRemainRevivalCount byte `json:"sentryRemainRevivalCount"`
	// SmallBulletPackageUnitPrice 远程兑换小弹丸单枚价格
	SmallBulletPackageUnitPrice float32 `json:"smallBulletPackageUnitPrice"`
	// BigBulletPackageUnitPrice 远程兑换大弹丸单枚价格
	BigBulletPackageUnitPrice float32 `json:"bigBulletPackageUnitPrice"`
	// BigBulletUnitPrice 补血点兑换大弹丸单枚价格
	BigBulletUnitPrice float32 `json:"bigBulletUnitPrice"`
	// SmallBulletUnitPrice 补血点兑换小弹丸单枚价格
	SmallBulletUnitPrice float32 `json:"smallBulletUnitPrice"`
	// CurrentRemoteHealPrice 当前远程补血价格
	CurrentRemoteHealPrice float32 `json:"currentRemoteHealPrice"`
	// SentryControlPrice 云台手控制哨兵基础花费
	SentryControlPrice float32 `json:"sentryControlPrice"`
}

// S1BattleProto2023ExchangeProErrorInfoNotify 兑换站报错内容同步客户端ExchangeProErrorInfoNotify
type S1BattleProto2023ExchangeProErrorInfoNotify struct {
	// TeamId 队伍ID
	TeamId byte `json:"teamId"`
	// IsOnline
	IsOnline byte `json:"isOnline"`
	// ArmErrorCode 机械臂报错信息，ServerUi显示 Err+数字
	ArmErrorCode uint16 `json:"armErrorCode"`
	// RfidErrorCode 0离线 1在线 2非矿石卡 3矿石已被兑换
	RfidErrorCode byte `json:"rfidErrorCode"`
	// RunningState 道具当前正在执行的操作，0代表无；长时间（大于10s）处于某个状态需要报警
	RunningState byte `json:"runningState"`
	// Ir1State 到位红外1
	Ir1State byte `json:"ir1State"`
	// Ir2State 到位红外2
	Ir2State byte `json:"ir2State"`
	// Ir3State 有矿红外
	Ir3State byte `json:"ir3State"`
	// IrErrorCode 1：到位红外触发，但有矿红外未触发
	IrErrorCode byte `json:"irErrorCode"`
	// TempAlertVal 温度阈值
	TempAlertVal int16 `json:"tempAlertVal"`
	// TempRmMotorRoll Roll轴电机温度
	TempRmMotorRoll int16 `json:"tempRmMotorRoll"`
	// TempRmMotorPitch Pitch轴电机温度
	TempRmMotorPitch int16 `json:"tempRmMotorPitch"`
	// TempRmMotorPush 推矿电机温度
	TempRmMotorPush int16 `json:"tempRmMotorPush"`
	// MotorRollState
	MotorRollState byte `json:"motorRollState"`
	// MotorPitchState
	MotorPitchState byte `json:"motorPitchState"`
	// MotorPushState 1：可放入矿石状态下，推矿电机未复位
	MotorPushState byte `json:"motorPushState"`
	// IsSoftwareEmergencyState 0无紧急停止；1软件紧急停止
	IsSoftwareEmergencyState byte `json:"isSoftwareEmergencyState"`
	// IsHardwareEmergencyState 0无紧急停止；1硬件紧急停止
	IsHardwareEmergencyState byte `json:"isHardwareEmergencyState"`
}

type S1BattleProtoAddBulletNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// BulletIndex
	BulletIndex byte `json:"bulletIndex"`
}

type S1BattleProtoArmorHitAck struct {
	// Result
	Result int32 `json:"result"`
}

type S1BattleProtoArmorHitReq struct {
	// Index
	Index int32 `json:"index"`
	// AttackType
	AttackType int32 `json:"attackType"`
	// AttackerUid
	AttackerUid uint64 `json:"attackerUid"`
	// AccValue
	AccValue int32 `json:"accValue"`
	// MicValue
	MicValue int32 `json:"micValue"`
}

type S1BattleProtoAttrsNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Tid
	Tid uint32 `json:"tid"`
	// Teamid
	Teamid byte `json:"teamid"`
	// SeatIndex
	SeatIndex byte `json:"seatIndex"`
	// Hp
	Hp int32 `json:"hp"`
	// HpMax
	HpMax int32 `json:"hpMax"`
	// Armor
	Armor int32 `json:"armor"`
	// ArmorMax
	ArmorMax int32 `json:"armorMax"`
	// Durability
	Durability int32 `json:"durability"`
	// DurabilityMax
	DurabilityMax int32 `json:"durabilityMax"`
	// Bullet
	Bullet int32 `json:"bullet"`
	// BulletMax
	BulletMax int32 `json:"bulletMax"`
	// Attack1
	Attack1 int32 `json:"attack1"`
	// Attack2
	Attack2 float32 `json:"attack2"`
}

type S1BattleProtoAutoStateNotify struct {
	// State
	State int32 `json:"state"`
}

type S1BattleProtoBeAttackNotify struct {
	// Index
	Index int32 `json:"index"`
	// AttackType
	AttackType int32 `json:"attackType"`
	// AttackerUid
	AttackerUid uint64 `json:"attackerUid"`
	// BeAttackerUid
	BeAttackerUid uint64 `json:"beAttackerUid"`
	// Damage
	Damage int32 `json:"damage"`
}

type S1BattleProtoBuffAddNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// BuffUid
	BuffUid uint64 `json:"buffUid"`
	// BuffTid
	BuffTid uint32 `json:"buffTid"`
	// BuffLevel
	BuffLevel uint32 `json:"buffLevel"`
	// BuffVisible
	BuffVisible byte `json:"buffVisible"`
	// BuffMaxTime
	BuffMaxTime float32 `json:"buffMaxTime"`
	// BuffLeftTime
	BuffLeftTime float32 `json:"buffLeftTime"`
	// MsgParams
	MsgParams string `json:"msgParams"`
}

type S1BattleProtoBuffDelNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// BuffUid
	BuffUid uint64 `json:"buffUid"`
}

type S1BattleProtoBuffModifyNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// BuffUid
	BuffUid uint64 `json:"buffUid"`
	// BuffTid
	BuffTid uint32 `json:"buffTid"`
	// BuffLevel
	BuffLevel uint32 `json:"buffLevel"`
	// BuffVisible
	BuffVisible byte `json:"buffVisible"`
	// BuffMaxTime
	BuffMaxTime float32 `json:"buffMaxTime"`
	// BuffLeftTime
	BuffLeftTime float32 `json:"buffLeftTime"`
	// UpdateProgress
	UpdateProgress byte `json:"updateProgress"`
	// MsgParams
	MsgParams string `json:"msgParams"`
}

type S1BattleProtoClientSyncNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// ConnClientState
	ConnClientState int32 `json:"connClientState"`
	// ConnRobotState
	ConnRobotState int32 `json:"connRobotState"`
	// Battery
	Battery int32 `json:"battery"`
	// SignalQuality
	SignalQuality int32 `json:"signalQuality"`
	// BattleMode
	BattleMode byte `json:"battleMode"`
}

type S1BattleProtoClientSyncReq struct {
	// Uid
	Uid uint64 `json:"uid"`
	// ConnClientState
	ConnClientState int32 `json:"connClientState"`
	// ConnRobotState
	ConnRobotState int32 `json:"connRobotState"`
	// Battery
	Battery int32 `json:"battery"`
	// SignalQuality
	SignalQuality int32 `json:"signalQuality"`
	// BattleMode
	BattleMode byte `json:"battleMode"`
}

type S1BattleProtoDeviceBulletNotify struct {
	// BulletBottleCount
	BulletBottleCount int32 `json:"bulletBottleCount"`
	// BulletBottleState
	BulletBottleState []byte `json:"bulletBottleState"`
	// PlayersUid
	PlayersUid []uint64 `json:"playersUid"`
}

type S1BattleProtoDeviceModuleNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// ProductType
	ProductType uint16 `json:"productType"`
	// AllDeviceCount
	AllDeviceCount uint16 `json:"allDeviceCount"`
	// AllDevice
	AllDevice []uint64 `json:"allDevice"`
	// AllDeviceConnection
	AllDeviceConnection []byte `json:"allDeviceConnection"`
	// AllDevicePrority
	AllDevicePrority []byte `json:"allDevicePrority"`
}

type S1BattleProtoDeviceModuleReq struct {
	// Uid
	Uid uint64 `json:"uid"`
	// ProductType
	ProductType uint16 `json:"productType"`
	// ConnectedDeviceCount
	ConnectedDeviceCount uint16 `json:"connectedDeviceCount"`
	// ConnectedDevice
	ConnectedDevice []uint64 `json:"connectedDevice"`
}

type S1BattleProtoFullSceneDataReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

type S1BattleProtoGameSettlementNotify struct {
	// TeamidWin
	TeamidWin int32 `json:"teamidWin"`
	// PlayersCount
	PlayersCount int32 `json:"playersCount"`
	// PlayersUid
	PlayersUid []uint64 `json:"playersUid"`
	// PlayersHp
	PlayersHp []int32 `json:"playersHp"`
	// PlayersHpMax
	PlayersHpMax []int32 `json:"playersHpMax"`
	// PlayersBekilled
	PlayersBekilled []int32 `json:"playersBekilled"`
	// Winreason
	Winreason int32 `json:"winreason"`
	// TeamCount
	TeamCount int32 `json:"teamCount"`
	// TeamHurtSum
	TeamHurtSum []int32 `json:"teamHurtSum"`
	// TeamAtkCount
	TeamAtkCount []int32 `json:"teamAtkCount"`
	// TeamPojiaCount
	TeamPojiaCount []int32 `json:"teamPojiaCount"`
	// DurationTime
	DurationTime int32 `json:"durationTime"`
}

type S1BattleProtoGunBulletNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// GunBulletMax
	GunBulletMax int16 `json:"gunBulletMax"`
	// GunBullet
	GunBullet int16 `json:"gunBullet"`
}

type S1BattleProtoGunFireAck struct {
	// Result
	Result int32 `json:"result"`
}

type S1BattleProtoGunFireNotify struct {
	// AttackerUid
	AttackerUid uint64 `json:"attackerUid"`
	// GunType
	GunType byte `json:"gunType"`
	// GunSpeed
	GunSpeed float32 `json:"gunSpeed"`
}

type S1BattleProtoGunFireReq struct {
	// AttackerUid
	AttackerUid uint64 `json:"attackerUid"`
	// GunType
	GunType byte `json:"gunType"`
	// GunSpeed
	GunSpeed float32 `json:"gunSpeed"`
}

type S1BattleProtoGunHeatNotify struct {
	// GunHeatMax
	GunHeatMax float32 `json:"gunHeatMax"`
	// GunHeat
	GunHeat float32 `json:"gunHeat"`
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
}

type S1BattleProtoIsCanUseAtknotify struct {
	// Teamcount
	Teamcount int32 `json:"teamcount"`
	// Teamid
	Teamid []int32 `json:"teamid"`
	// State
	State []int32 `json:"state"`
}

type S1BattleProtoIsCanUsePjnotify struct {
	// Teamcount
	Teamcount int32 `json:"teamcount"`
	// Teamid
	Teamid []int32 `json:"teamid"`
	// State
	State []int32 `json:"state"`
}

type S1BattleProtoMarkerDetectAck struct {
	// Result
	Result byte `json:"result"`
}

type S1BattleProtoMarkerDetectReq struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Act
	Act byte `json:"act"`
	// MarkerId
	MarkerId byte `json:"markerId"`
	// MarkerStr
	MarkerStr string `json:"markerStr"`
}

type S1BattleProtoModuleOffLineInfoNotify struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

type S1BattleProtoPlaceholderNotify struct {
	// EventName
	EventName string `json:"eventName"`
	// State
	State byte `json:"state"`
}

type S1BattleProtoPlayerCommandAutoControlNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandGunFireNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandLockScreenNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandMoveNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandSetDisconnectNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandUserCustomActionNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerCommandUserInputNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1BattleProtoPlayerDeadNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// KillerUid
	KillerUid uint64 `json:"killerUid"`
	// KillHonor
	KillHonor int32 `json:"killHonor"`
	// DeadType
	DeadType int32 `json:"deadType"`
	// Flag0
	Flag0 int32 `json:"flag0"`
	// Flag1
	Flag1 int32 `json:"flag1"`
	// Flag2
	Flag2 int32 `json:"flag2"`
}

type S1BattleProtoPlayerInitOkNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
}

type S1BattleProtoPlayerReliveNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
}

type S1BattleProtoPlayerResetNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
}

type S1BattleProtoPoJiaNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// State
	State byte `json:"state"`
	// Timeleft
	Timeleft float32 `json:"timeleft"`
	// Timemax
	Timemax float32 `json:"timemax"`
}

type S1BattleProtoProgressNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// EventName
	EventName string `json:"eventName"`
	// TimeMax
	TimeMax float32 `json:"timeMax"`
	// TimeLeft
	TimeLeft float32 `json:"timeLeft"`
	// State
	State byte `json:"state"`
}

type S1BattleProtoRfidnotify struct {
	// S0Id
	S0Id byte `json:"s0Id"`
	// Flag
	Flag byte `json:"flag"`
	// Type
	Type byte `json:"type"`
	// Area
	Area byte `json:"area"`
	// Data0
	Data0 byte `json:"data0"`
	// Data1
	Data1 byte `json:"data1"`
	// Data2
	Data2 byte `json:"data2"`
	// Data3
	Data3 byte `json:"data3"`
	// Data4
	Data4 byte `json:"data4"`
	// Data5
	Data5 byte `json:"data5"`
}

type S1BattleProtoRfidreq struct {
	// Type
	Type byte `json:"type"`
	// Area
	Area byte `json:"area"`
	// Data0
	Data0 byte `json:"data0"`
	// Data1
	Data1 byte `json:"data1"`
	// Data2
	Data2 byte `json:"data2"`
	// Data3
	Data3 byte `json:"data3"`
	// Data4
	Data4 byte `json:"data4"`
	// Data5
	Data5 byte `json:"data5"`
	// Timestamp
	Timestamp uint64 `json:"timestamp"`
}

type S1BattleProtoReduceReliveTimeNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// TimeMax
	TimeMax float32 `json:"timeMax"`
	// TimeLeft
	TimeLeft float32 `json:"timeLeft"`
}

type S1BattleProtoReloginNotify struct {
	// SelfUid
	SelfUid uint64 `json:"selfUid"`
	// SelfTid
	SelfTid uint32 `json:"selfTid"`
	// SelfTeamId
	SelfTeamId byte `json:"selfTeamId"`
	// SelfSeatIndex
	SelfSeatIndex byte `json:"selfSeatIndex"`
	// SceneState
	SceneState int32 `json:"sceneState"`
	// SceneStateTimeLeft
	SceneStateTimeLeft int32 `json:"sceneStateTimeLeft"`
	// SceneTeamsCount
	SceneTeamsCount int32 `json:"sceneTeamsCount"`
	// SceneTeamsPlayerMax
	SceneTeamsPlayerMax int32 `json:"sceneTeamsPlayerMax"`
	// ScenePlayersCount
	ScenePlayersCount int32 `json:"scenePlayersCount"`
	// TeamsTotalDamage
	TeamsTotalDamage []uint32 `json:"teamsTotalDamage"`
	// PlayersTotalBuffCount
	PlayersTotalBuffCount uint32 `json:"playersTotalBuffCount"`
	// PlayersUid
	PlayersUid []uint64 `json:"playersUid"`
	// PlayersTid
	PlayersTid []uint32 `json:"playersTid"`
	// PlayersTeamId
	PlayersTeamId []byte `json:"playersTeamId"`
	// PlayersSeatIndex
	PlayersSeatIndex []byte `json:"playersSeatIndex"`
	// PlayersState
	PlayersState []byte `json:"playersState"`
	// PlayersCdLeftTime
	PlayersCdLeftTime []float32 `json:"playersCdLeftTime"`
	// PlayersCdMaxTime
	PlayersCdMaxTime []float32 `json:"playersCdMaxTime"`
	// PlayersHp
	PlayersHp []uint32 `json:"playersHp"`
	// PlayersMaxHp
	PlayersMaxHp []uint32 `json:"playersMaxHp"`
	// PlayersHeat
	PlayersHeat []float32 `json:"playersHeat"`
	// PlayersMaxHeat
	PlayersMaxHeat []float32 `json:"playersMaxHeat"`
	// PlayersBulletNum
	PlayersBulletNum []uint32 `json:"playersBulletNum"`
	// PlayersMaxBulletNum
	PlayersMaxBulletNum []uint32 `json:"playersMaxBulletNum"`
	// PlayersArmor
	PlayersArmor []int32 `json:"playersArmor"`
	// PlayersMaxArmor
	PlayersMaxArmor []int32 `json:"playersMaxArmor"`
	// PlayersDurability
	PlayersDurability []int32 `json:"playersDurability"`
	// PlayersMaxDurability
	PlayersMaxDurability []int32 `json:"playersMaxDurability"`
	// PlayersBuffCount
	PlayersBuffCount []uint32 `json:"playersBuffCount"`
	// PlayersBuffUid
	PlayersBuffUid [][]uint64 `json:"playersBuffUid"`
	// PlayersBuffTid
	PlayersBuffTid [][]uint32 `json:"playersBuffTid"`
	// PlayersBuffLevel
	PlayersBuffLevel [][]uint32 `json:"playersBuffLevel"`
	// PlayersBuffLeftTime
	PlayersBuffLeftTime [][]float32 `json:"playersBuffLeftTime"`
	// PlayersGunCtrlState
	PlayersGunCtrlState []byte `json:"playersGunCtrlState"`
	// PlayersMoveCtrlState
	PlayersMoveCtrlState []byte `json:"playersMoveCtrlState"`
	// PlayersUserInputState
	PlayersUserInputState []byte `json:"playersUserInputState"`
	// PlayersLockScreenState
	PlayersLockScreenState []byte `json:"playersLockScreenState"`
	// PlayersUserCustomActionState
	PlayersUserCustomActionState []byte `json:"playersUserCustomActionState"`
	// PlayersAutoModeCtrlState
	PlayersAutoModeCtrlState []byte `json:"playersAutoModeCtrlState"`
	// DevicesBaseStationState
	DevicesBaseStationState []byte `json:"devicesBaseStationState"`
	// DevicesBaseStationStateLeftTime
	DevicesBaseStationStateLeftTime []float32 `json:"devicesBaseStationStateLeftTime"`
	// DevicesRuneState
	DevicesRuneState []byte `json:"devicesRuneState"`
	// DevicesRuneStateLeftTime
	DevicesRuneStateLeftTime []float32 `json:"devicesRuneStateLeftTime"`
	// WifiName
	WifiName string `json:"wifiName"`
	// WifiPassword
	WifiPassword string `json:"wifiPassword"`
}

type S1BattleProtoSceneInfosNotify struct {
	// PlayersCount
	PlayersCount int32 `json:"playersCount"`
	// PlayersUid
	PlayersUid []uint64 `json:"playersUid"`
	// PlayersTid
	PlayersTid []uint32 `json:"playersTid"`
	// PlayersTeamId
	PlayersTeamId []byte `json:"playersTeamId"`
	// PlayersSeatIndex
	PlayersSeatIndex []byte `json:"playersSeatIndex"`
}

type S1BattleProtoSceneStateNotify struct {
	// State
	State ESceneStateType `json:"state"`
	// TimeLeft
	TimeLeft int32 `json:"timeLeft"`
	// TeamCount
	TeamCount int32 `json:"teamCount"`
	// TeamHurtSum
	TeamHurtSum []int32 `json:"teamHurtSum"`
	// TeamBekillCount
	TeamBekillCount []int32 `json:"teamBekillCount"`
}

type S1BattleProtoShareBulletReq struct {
	// TeamId
	TeamId byte `json:"teamId"`
	// BulletIndex
	BulletIndex byte `json:"bulletIndex"`
}

type S1BattleProtoTeamInfoNotify struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

type S1BattleProtoUploadModuleInfo struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

type S1BattleProtoWarningNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// WarningLevel
	WarningLevel byte `json:"warningLevel"`
}

type S1BattleProtoBaseArmorStateNotify struct {
	// PlayerUid
	PlayerUid uint64 `json:"playerUid"`
	// ArmorId
	ArmorId int16 `json:"armorId"`
	// ArmorCount
	ArmorCount int16 `json:"armorCount"`
	// ArmorValue
	ArmorValue int16 `json:"armorValue"`
}

type S1BattleProtot2022RobotArmorCfgData struct {
	// ArmorParaDataCount
	ArmorParaDataCount int32 `json:"armorParaDataCount"`
	// ArmorParaDatas
	ArmorParaDatas []S1BattleProto2022ArmorParaConfigItem `json:"armorParaDatas"`
}

type S1ClientTransferRobotProto struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data []byte `json:"data"`
}

type S1ProtoClientExceptionInfo struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

// S1ProtoClientNetworkInfoNotify S1ProtoClientNetworkInfoNotify
type S1ProtoClientNetworkInfoNotify struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// VtUplinkSpeed
	VtUplinkSpeed uint32 `json:"vtUplinkSpeed"`
	// VtDownlinkSpeed
	VtDownlinkSpeed uint32 `json:"vtDownlinkSpeed"`
}

type S1ProtoClientSyncNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// ConnClientState
	ConnClientState int32 `json:"connClientState"`
	// ConnRobotState
	ConnRobotState int32 `json:"connRobotState"`
	// Battery
	Battery int32 `json:"battery"`
	// SignalQuality
	SignalQuality int32 `json:"signalQuality"`
	// BattleMode
	BattleMode byte `json:"battleMode"`
	// OnlineModuleCount
	OnlineModuleCount int32 `json:"onlineModuleCount"`
	// OnlineModule
	OnlineModule []uint64 `json:"onlineModule"`
}

type S1ProtoClientSyncReq struct {
	// Uid
	Uid uint64 `json:"uid"`
	// ConnClientState
	ConnClientState int32 `json:"connClientState"`
	// ConnRobotState
	ConnRobotState int32 `json:"connRobotState"`
	// Battery
	Battery int32 `json:"battery"`
	// SignalQuality
	SignalQuality int32 `json:"signalQuality"`
	// BattleMode
	BattleMode byte `json:"battleMode"`
	// OnlineModuleCount
	OnlineModuleCount int32 `json:"onlineModuleCount"`
	// OnlineModule
	OnlineModule []uint64 `json:"onlineModule"`
}

type S1ProtoEnterRoomAck struct {
	// ResultId
	ResultId EEnterRoomAckResultType `json:"resultId"`
	// Uid
	Uid uint64 `json:"uid"`
}

type S1ProtoEnterRoomNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
}

type S1ProtoEnterRoomReq struct {
	// Nouse1
	Nouse1 byte `json:"nouse1"`
	// Nouse2
	Nouse2 byte `json:"nouse2"`
}

type S1ProtoGmcommandAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
}

type S1ProtoGmcommandReq struct {
	// Command
	Command string `json:"command"`
	// Pars
	Pars string `json:"pars"`
}

type S1ProtoHeader struct {
	// ProtoId
	ProtoId uint16 `json:"protoId"`
	// ProtoType
	ProtoType uint16 `json:"protoType"`
	// DataLen
	DataLen uint32 `json:"dataLen"`
	// SequenceId
	SequenceId byte `json:"sequenceId"`
	// AckType
	AckType byte `json:"ackType"`
	// Nouse1
	Nouse1 byte `json:"nouse1"`
	// Nouse2
	Nouse2 byte `json:"nouse2"`
}

type S1ProtoHeartBeatAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
	// S0Clientid
	S0Clientid byte `json:"s0Clientid"`
}

type S1ProtoHeartBeatReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
	// S0Clientid
	S0Clientid byte `json:"s0Clientid"`
}

type S1ProtoLeaveRoomAck struct {
	// ResultId
	ResultId ELeaveRoomAckResultType `json:"resultId"`
}

type S1ProtoLeaveRoomNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Willgo
	Willgo ERoleLocation `json:"willgo"`
}

type S1ProtoLeaveRoomReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

type S1ProtoLoginAck struct {
	// ResultId
	ResultId ELoginAckResultType `json:"resultId"`
	// Uid
	Uid uint64 `json:"uid"`
	// Token
	Token string `json:"token"`
}

type S1ProtoLoginReq struct {
	// Account
	Account string `json:"account"`
	// Password
	Password string `json:"password"`
	// Version
	Version float32 `json:"version"`
	// Tid
	Tid uint32 `json:"tid"`
	// Teamid
	Teamid uint32 `json:"teamid"`
	// Hash
	Hash int64 `json:"hash"`
}

type S1ProtoLogoutAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
}

type S1ProtoLogoutNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Reason
	Reason byte `json:"reason"`
}

type S1ProtoLogoutReq struct {
	// Account
	Account string `json:"account"`
}

// S1ProtoMarkDetectResult S1ProtoClientNetworkInfoNotify
type S1ProtoMarkDetectResult struct {
	// X
	X float32 `json:"x"`
	// Y
	Y float32 `json:"y"`
	// W
	W float32 `json:"w"`
	// H
	H float32 `json:"h"`
	// Pitch
	Pitch float32 `json:"pitch"`
	// Yaw
	Yaw float32 `json:"yaw"`
	// Roll
	Roll float32 `json:"roll"`
	// T44C2M
	T44C2M []float32 `json:"t44C2M"`
	// Color
	Color byte `json:"color"`
	// MarkerId
	MarkerId uint16 `json:"markerId"`
	// Distance
	Distance uint16 `json:"distance"`
}

// S1ProtoMarkDetectResultAck S1ProtoClientNetworkInfoNotify
type S1ProtoMarkDetectResultAck struct {
	// Result
	Result int32 `json:"result"`
	// X
	X float32 `json:"x"`
	// Y
	Y float32 `json:"y"`
	// W
	W float32 `json:"w"`
	// H
	H float32 `json:"h"`
	// Color
	Color byte `json:"color"`
	// MarkerId
	MarkerId uint16 `json:"markerId"`
	// Distance
	Distance uint16 `json:"distance"`
}

// S1ProtoNetworkInfoNotify S1ProtoNetworkInfoNotify
type S1ProtoNetworkInfoNotify struct {
	// Robotid
	Robotid byte `json:"robotid"`
	// WifiUplinkSpeed
	WifiUplinkSpeed uint32 `json:"wifiUplinkSpeed"`
	// WifiDownlinkSpeed
	WifiDownlinkSpeed uint32 `json:"wifiDownlinkSpeed"`
	// VtUplinkSpeed
	VtUplinkSpeed uint32 `json:"vtUplinkSpeed"`
	// VtDownlinkSpeed
	VtDownlinkSpeed uint32 `json:"vtDownlinkSpeed"`
	// WifiUplinkPlr
	WifiUplinkPlr uint16 `json:"wifiUplinkPlr"`
	// WifiDownlinkPlr
	WifiDownlinkPlr uint16 `json:"wifiDownlinkPlr"`
	// VtUplinkPlr
	VtUplinkPlr uint16 `json:"vtUplinkPlr"`
	// VtDownlinkPlr
	VtDownlinkPlr uint16 `json:"vtDownlinkPlr"`
	// OtherType
	OtherType byte `json:"otherType"`
	// OtherPlr
	OtherPlr uint16 `json:"otherPlr"`
	// OtherDelay
	OtherDelay uint16 `json:"otherDelay"`
	// Delay
	Delay uint32 `json:"delay"`
}

type S1ProtoPingAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
}

type S1ProtoPingReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

type S1ProtoPlayerCommandDisconnectNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Value
	Value byte `json:"value"`
}

type S1ProtoPlayerCommandUserInputNotify struct {
	// Value
	Value byte `json:"value"`
}

type S1ProtoReLoginAck struct {
	// ResultId
	ResultId ELoginAckResultType `json:"resultId"`
	// Uid
	Uid uint64 `json:"uid"`
	// Token
	Token string `json:"token"`
	// WifiName
	WifiName string `json:"wifiName"`
	// WifiPassword
	WifiPassword string `json:"wifiPassword"`
}

type S1ProtoReLoginReq struct {
	// Account
	Account string `json:"account"`
	// Password
	Password string `json:"password"`
	// Version
	Version float32 `json:"version"`
	// Tid
	Tid uint32 `json:"tid"`
}

type S1ProtoReLoginRoomNotify struct {
	// SelfUid
	SelfUid uint64 `json:"selfUid"`
	// SelfTid
	SelfTid uint32 `json:"selfTid"`
	// SelfTeamId
	SelfTeamId byte `json:"selfTeamId"`
	// SelfSeatIndex
	SelfSeatIndex byte `json:"selfSeatIndex"`
	// RoomState
	RoomState int32 `json:"roomState"`
	// RoomStateTimeLeft
	RoomStateTimeLeft int32 `json:"roomStateTimeLeft"`
	// RoomSeatIslock
	RoomSeatIslock int32 `json:"roomSeatIslock"`
	// RoomTeamsCount
	RoomTeamsCount int32 `json:"roomTeamsCount"`
	// RoomTeamsPlayerMax
	RoomTeamsPlayerMax int32 `json:"roomTeamsPlayerMax"`
	// RoomPlayersCount
	RoomPlayersCount int32 `json:"roomPlayersCount"`
	// PlayersUid
	PlayersUid []uint64 `json:"playersUid"`
	// PlayersTid
	PlayersTid []uint32 `json:"playersTid"`
	// PlayersTeamId
	PlayersTeamId []byte `json:"playersTeamId"`
	// PlayersSeatIndex
	PlayersSeatIndex []byte `json:"playersSeatIndex"`
	// WifiName
	WifiName string `json:"wifiName"`
	// WifiPassword
	WifiPassword string `json:"wifiPassword"`
}

type S1ProtoRobotNetworkStatusAck struct {
	// WifiDownlinkPlr
	WifiDownlinkPlr uint16 `json:"wifiDownlinkPlr"`
	// VtDownlinkPlr
	VtDownlinkPlr uint16 `json:"vtDownlinkPlr"`
}

type S1ProtoRobotNetworkStatusReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

// S1ProtoRobotStudentSerialPortInfos S1ProtoRobotStudentSerialPortInfos
type S1ProtoRobotStudentSerialPortInfos struct {
	// OtherType
	OtherType byte `json:"otherType"`
	// OtherPlr
	OtherPlr uint16 `json:"otherPlr"`
	// OtherDelay
	OtherDelay uint16 `json:"otherDelay"`
}

type S1ProtoRoomInfosNotify struct {
	// PlayersCount
	PlayersCount int32 `json:"playersCount"`
	// PlayersUid
	PlayersUid []uint64 `json:"playersUid"`
	// PlayersTid
	PlayersTid []uint32 `json:"playersTid"`
	// PlayersTeamId
	PlayersTeamId []byte `json:"playersTeamId"`
	// PlayersSeatIndex
	PlayersSeatIndex []byte `json:"playersSeatIndex"`
	// WifiName
	WifiName string `json:"wifiName"`
	// WifiPassword
	WifiPassword string `json:"wifiPassword"`
	// CurrMatch
	CurrMatch string `json:"currMatch"`
}

type S1ProtoRoomPauseNotify struct {
	// TimeElapse
	TimeElapse int32 `json:"timeElapse"`
}

type S1ProtoRoomStateNotify struct {
	// State
	State ERoomStateType `json:"state"`
	// TimeLeft
	TimeLeft int32 `json:"timeLeft"`
	// IsPause
	IsPause int32 `json:"isPause"`
}

type S1ProtoSetCurrMatchNotify struct {
	// CurrMatch
	CurrMatch string `json:"currMatch"`
	// CurrToken
	CurrToken string `json:"currToken"`
}

type S1ProtoSetLockSeatNotify struct {
	// LockSeat
	LockSeat byte `json:"lockSeat"`
}

type S1ProtoSetReadyAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
}

type S1ProtoSetReadyNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// State
	State byte `json:"state"`
}

type S1ProtoSetReadyReq struct {
	// State
	State byte `json:"state"`
}

type S1ProtoSetSvrInfoAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
}

type S1ProtoSetSvrInfoReq struct {
	// SvrName
	SvrName string `json:"svrName"`
}

type S1ProtoSetTeamAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
	// TeamId
	TeamId byte `json:"teamId"`
	// SeatIndex
	SeatIndex byte `json:"seatIndex"`
}

type S1ProtoSetTeamNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// TeamId
	TeamId byte `json:"teamId"`
	// SeatIndex
	SeatIndex byte `json:"seatIndex"`
}

type S1ProtoSetTeamPlayerNumNotify struct {
	// PlayerNum
	PlayerNum byte `json:"playerNum"`
}

type S1ProtoSetTeamReq struct {
	// TeamId
	TeamId byte `json:"teamId"`
	// SeatIndex
	SeatIndex byte `json:"seatIndex"`
}

type S1ProtoSetTidAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
	// Tid
	Tid uint32 `json:"tid"`
}

type S1ProtoSetTidNotify struct {
	// Uid
	Uid uint64 `json:"uid"`
	// Tid
	Tid uint32 `json:"tid"`
}

type S1ProtoSetTidReq struct {
	// Tid
	Tid uint32 `json:"tid"`
	// Token
	Token string `json:"token"`
}

type S1ProtoSetTokenAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
}

type S1ProtoSetTokenReq struct {
	// Token
	Token string `json:"token"`
}

type S1ProtoSetWifiInfoNotify struct {
	// WifiName
	WifiName string `json:"wifiName"`
	// WifiPassword
	WifiPassword string `json:"wifiPassword"`
}

type S1ProtoSvrStateAck struct {
	// State
	State ERoomStateType `json:"state"`
	// TimeLeft
	TimeLeft int32 `json:"timeLeft"`
	// IsPause
	IsPause int32 `json:"isPause"`
	// CurrMatch
	CurrMatch string `json:"currMatch"`
}

type S1ProtoSvrStateReq struct {
	// Nouse
	Nouse byte `json:"nouse"`
}

type S1ProtoTeamCampNotify struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

type S1ProtoTeamInfoNotify struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

type S1ProtoTeamLineupInfoNotify struct {
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

type S1ProtoTeamLogoNotify struct {
	// Teamid
	Teamid int32 `json:"teamid"`
	// DataLen
	DataLen int32 `json:"dataLen"`
	// Data
	Data string `json:"data"`
}

type S1ProtoTechnicalPauseInfoNotify struct {
	// PauseTimeType 当前暂停时间类型 0：未暂停 1：官方技术暂停中
	PauseTimeType byte `json:"pauseTimeType"`
	// PauseSide 当前暂停的发起方：0红，1蓝，2官方
	PauseSide byte `json:"pauseSide"`
	// RedShortPauseLeftCount 红方剩余2min暂停次数
	RedShortPauseLeftCount byte `json:"redShortPauseLeftCount"`
	// RedLongPauseLeftCount 红方剩余3min暂停次数
	RedLongPauseLeftCount byte `json:"redLongPauseLeftCount"`
	// BlueShortPauseLeftCount 蓝方剩余2min暂停次数
	BlueShortPauseLeftCount byte `json:"blueShortPauseLeftCount"`
	// BlueLongPauseLeftCount 蓝方剩余3min暂停次数
	BlueLongPauseLeftCount byte `json:"blueLongPauseLeftCount"`
	// PausedDuration 正向计时暂停的时长(ms)
	PausedDuration uint32 `json:"pausedDuration"`
}

type S1ProtoTestAck struct {
	// ResultId
	ResultId int32 `json:"resultId"`
	// Test
	Test string `json:"test"`
}

type S1ProtoTestReq struct {
	// Test
	Test string `json:"test"`
}

type S1ProtoUserStateAck struct {
	// Online
	Online int32 `json:"online"`
}

type S1ProtoUserStateReq struct {
	// Account
	Account string `json:"account"`
	// Password
	Password string `json:"password"`
}
