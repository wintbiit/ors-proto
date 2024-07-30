package proto

import (
	"reflect"
)

const (
	ProtoIDS1ProtoLoginReq                                      uint16 = 9000
	ProtoIDS1ProtoLoginAck                                      uint16 = 9001
	ProtoIDS1ProtoReLoginReq                                    uint16 = 9002
	ProtoIDS1ProtoReLoginAck                                    uint16 = 9003
	ProtoIDS1ProtoReLoginRoomNotify                             uint16 = 9004
	ProtoIDS1ProtoLogoutReq                                     uint16 = 9005
	ProtoIDS1ProtoLogoutAck                                     uint16 = 9006
	ProtoIDS1ProtoLogoutNotify                                  uint16 = 9007
	ProtoIDS1ProtoHeartBeatReq                                  uint16 = 9008
	ProtoIDS1ProtoHeartBeatAck                                  uint16 = 9009
	ProtoIDS1ProtoTestReq                                       uint16 = 9010
	ProtoIDS1ProtoTestAck                                       uint16 = 9011
	ProtoIDS1ProtoPingReq                                       uint16 = 9012
	ProtoIDS1ProtoPingAck                                       uint16 = 9013
	ProtoIDS1ProtoEnterRoomReq                                  uint16 = 9014
	ProtoIDS1ProtoEnterRoomAck                                  uint16 = 9015
	ProtoIDS1ProtoEnterRoomNotify                               uint16 = 9016
	ProtoIDS1ProtoRoomInfosNotify                               uint16 = 9017
	ProtoIDS1ProtoRoomStateNotify                               uint16 = 9018
	ProtoIDS1ProtoRoomPauseNotify                               uint16 = 9019
	ProtoIDS1ProtoLeaveRoomReq                                  uint16 = 9020
	ProtoIDS1ProtoLeaveRoomAck                                  uint16 = 9021
	ProtoIDS1ProtoLeaveRoomNotify                               uint16 = 9022
	ProtoIDS1ProtoSetTeamReq                                    uint16 = 9023
	ProtoIDS1ProtoSetTeamAck                                    uint16 = 9024
	ProtoIDS1ProtoSetTeamNotify                                 uint16 = 9025
	ProtoIDS1ProtoSetTidReq                                     uint16 = 9026
	ProtoIDS1ProtoSetTidAck                                     uint16 = 9027
	ProtoIDS1ProtoSetTidNotify                                  uint16 = 9028
	ProtoIDS1ProtoSetReadyReq                                   uint16 = 9029
	ProtoIDS1ProtoSetReadyAck                                   uint16 = 9030
	ProtoIDS1ProtoSetReadyNotify                                uint16 = 9031
	ProtoIDS1ProtoSetTokenReq                                   uint16 = 9032
	ProtoIDS1ProtoSetTokenAck                                   uint16 = 9033
	ProtoIDS1ProtoSetSvrInfoReq                                 uint16 = 9034
	ProtoIDS1ProtoSetSvrInfoAck                                 uint16 = 9035
	ProtoIDS1ProtoSetWifiInfoNotify                             uint16 = 9036
	ProtoIDS1ProtoSetLockSeatNotify                             uint16 = 9037
	ProtoIDS1ProtoSetCurrMatchNotify                            uint16 = 9038
	ProtoIDS1ProtoSetTeamPlayerNumNotify                        uint16 = 9039
	ProtoIDS1ProtoGMCommandReq                                  uint16 = 9040
	ProtoIDS1ProtoGMCommandAck                                  uint16 = 9041
	ProtoIDS1ProtoPlayerCommand_UserInput_Notify                uint16 = 9042
	ProtoIDS1ProtoPlayerCommand_Disconnect_Notify               uint16 = 9043
	ProtoIDS1ProtoClientSyncReq                                 uint16 = 9044
	ProtoIDS1ProtoClientSyncNotify                              uint16 = 9045
	ProtoIDS1ProtoTeamInfoNotify                                uint16 = 9046
	ProtoIDS1ProtoTeamLogoNotify                                uint16 = 9047
	ProtoIDS1ProtoTeamCampNotify                                uint16 = 9048
	ProtoIDS1ProtoSvrStateReq                                   uint16 = 9049
	ProtoIDS1ProtoSvrStateAck                                   uint16 = 9050
	ProtoIDS1ProtoUserStateReq                                  uint16 = 9051
	ProtoIDS1ProtoUserStateAck                                  uint16 = 9052
	ProtoIDS1ClientTransferRobotProto                           uint16 = 9053
	ProtoIDS1ProtoRobotNetworkStatusReq                         uint16 = 9054
	ProtoIDS1ProtoRobotNetworkStatusAck                         uint16 = 9055
	ProtoIDS1ProtoClientNetworkInfoNotify                       uint16 = 9056
	ProtoIDS1ProtoRobotStudentSerialPortInfos                   uint16 = 9057
	ProtoIDS1ProtoNetworkInfoNotify                             uint16 = 9058
	ProtoIDS1ProtoTechnicalPauseInfoNotify                      uint16 = 9059
	ProtoIDS1ProtoMarkDetectResult                              uint16 = 9060
	ProtoIDS1ProtoMarkDetectResultAck                           uint16 = 9061
	ProtoIDS1ProtoClientExceptionInfo                           uint16 = 9062
	ProtoIDS1BattleProtoAutoStateNotify                         uint16 = 10000
	ProtoIDS1BattleProtoArmorHitReq                             uint16 = 10001
	ProtoIDS1BattleProtoArmorHitAck                             uint16 = 10002
	ProtoIDS1BattleProtoBeAttackNotify                          uint16 = 10003
	ProtoIDS1BattleProtoGunFireReq                              uint16 = 10004
	ProtoIDS1BattleProtoGunFireAck                              uint16 = 10005
	ProtoIDS1BattleProtoGunFireNotify                           uint16 = 10006
	ProtoIDS1BattleProtoGunHeatNotify                           uint16 = 10007
	ProtoIDS1BattleProtoGunBulletNotify                         uint16 = 10008
	ProtoIDS1BattleProtoAttrsNotify                             uint16 = 10009
	ProtoIDS1BattleProtoPlayerInitOkNotify                      uint16 = 10010
	ProtoIDS1BattleProtoPlayerDeadNotify                        uint16 = 10011
	ProtoIDS1BattleProtoPlayerReliveNotify                      uint16 = 10012
	ProtoIDS1BattleProtoBuffAddNotify                           uint16 = 10013
	ProtoIDS1BattleProtoBuffDelNotify                           uint16 = 10014
	ProtoIDS1BattleProtoBuffModifyNotify                        uint16 = 10015
	ProtoIDS1BattleProtoMarkerDetectReq                         uint16 = 10016
	ProtoIDS1BattleProtoMarkerDetectAck                         uint16 = 10017
	ProtoIDS1BattleProtoProgressNotify                          uint16 = 10018
	ProtoIDS1BattleProtoDeviceModuleReq                         uint16 = 10019
	ProtoIDS1BattleProtoDeviceModuleNotify                      uint16 = 10020
	ProtoIDS1BattleProtoPoJiaNotify                             uint16 = 10021
	ProtoIDS1BattleProtoPlayerCommand_GunFire_Notify            uint16 = 10022
	ProtoIDS1BattleProtoPlayerCommand_Move_Notify               uint16 = 10023
	ProtoIDS1BattleProtoPlayerCommand_LockScreen_Notify         uint16 = 10024
	ProtoIDS1BattleProtoPlayerCommand_AutoControl_Notify        uint16 = 10025
	ProtoIDS1BattleProtoPlayerCommand_UserInput_Notify          uint16 = 10026
	ProtoIDS1BattleProtoPlayerCommand_UserCustomAction_Notify   uint16 = 10027
	ProtoIDS1BattleProtoPlayerCommand_SetDisconnectNotify       uint16 = 10028
	ProtoIDS1BattleProtoSceneInfosNotify                        uint16 = 10029
	ProtoIDS1BattleProtoSceneStateNotify                        uint16 = 10030
	ProtoIDS1BattleProtoGameSettlementNotify                    uint16 = 10031
	ProtoIDS1BattleProtoReloginNotify                           uint16 = 10032
	ProtoIDS1BattleProtoIsCanUseATKNotify                       uint16 = 10033
	ProtoIDS1BattleProtoIsCanUsePJNotify                        uint16 = 10034
	ProtoIDS1BattleProtoClientSyncReq                           uint16 = 10035
	ProtoIDS1BattleProtoClientSyncNotify                        uint16 = 10036
	ProtoIDS1BattleProtoDeviceBulletNotify                      uint16 = 10037
	ProtoIDS1BattleProtoWarningNotify                           uint16 = 10038
	ProtoIDS1BattleProtoPlayerResetNotify                       uint16 = 10039
	ProtoIDS1BattleProtoReduceReliveTimeNotify                  uint16 = 10040
	ProtoIDS1BattleProtoFullSceneDataReq                        uint16 = 10041
	ProtoIDS1BattleProtoTeamInfoNotify                          uint16 = 10042
	ProtoIDS1BattleProtoPlaceholderNotify                       uint16 = 10043
	ProtoIDS1BattleProtoModuleOffLineInfoNotify                 uint16 = 10044
	ProtoIDS1BattleProtoAddBulletNotify                         uint16 = 10045
	ProtoIDS1BattleProtoShareBulletReq                          uint16 = 10046
	ProtoIDS1BattleProtoUploadModuleInfo                        uint16 = 10047
	ProtoIDS1BattleProto2022RobotInitCfgTab                     uint16 = 10048
	ProtoIDS1BattleProto2022ConfigTabAck                        uint16 = 10049
	ProtoIDS1BattleProto2022RobotStatus                         uint16 = 10050
	ProtoIDS1BattleProto2022RobotsDataSync                      uint16 = 10051
	ProtoIDS1BattleProto2022RobotsWeaponFire                    uint16 = 10052
	ProtoIDS1BattleProto2022RobotHited                          uint16 = 10053
	ProtoIDS1BattleProto2022VtmSpeedSync                        uint16 = 10054
	ProtoIDS1BattleProto2022ClientsServerBaseSync               uint16 = 10055
	ProtoIDS1BattleProto2022ClientsServerSync                   uint16 = 10056
	ProtoIDS1BattleProto2022ClientsFirstSync                    uint16 = 10057
	ProtoIDS1BattleProto2022ClientsRobotStatusNotify            uint16 = 10058
	ProtoIDS1BattleProto2022ClientsRobotModuleErrorInfo         uint16 = 10059
	ProtoIDS1BattleProto2022TestTimeDelayDownData               uint16 = 10060
	ProtoIDS1BattleProto2022TestTimeDelayDownComplete           uint16 = 10061
	ProtoIDS1BattleProto2022TestDelayDownDataAck                uint16 = 10062
	ProtoIDS1BattleProto2022CONTROL_CMD_T                       uint16 = 10063
	ProtoIDS1BattleProto2022ClientGMCommand                     uint16 = 10064
	ProtoIDS1BattleProto2022ClientCurrentProgress               uint16 = 10065
	ProtoIDS1BattleProto2022ClientHitNotify                     uint16 = 10066
	ProtoIDS1BattleProto2022ReLoginDataSync                     uint16 = 10067
	ProtoIDS1BattleProto2022ClientResultDataInfo                uint16 = 10068
	ProtoIDS1BattleProto2022EnvBaseShellControl                 uint16 = 10069
	ProtoIDS1BattleProto2022RFIDReq                             uint16 = 10070
	ProtoIDS1BattleProto2022RFIDNotify                          uint16 = 10071
	ProtoIDS1BattleProto2022IOCtrCfgSet                         uint16 = 10072
	ProtoIDS1BattleProto2022IOCtrCfgSetAck                      uint16 = 10073
	ProtoIDS1BattleProto2022IOCtrSetVal                         uint16 = 10074
	ProtoIDS1BattleProto2022IOCtrSetValAck                      uint16 = 10075
	ProtoIDS1BattleProto2022IOCtrState                          uint16 = 10076
	ProtoIDS1BattleProto2022IOCtrTriggerVal                     uint16 = 10077
	ProtoIDS1BattleProto2022EnvDevicesDescripeReq               uint16 = 10078
	ProtoIDS1BattleProto2022EnvDevicesDescripeAck               uint16 = 10079
	ProtoIDS1BattleProto2022S2C_PowerStateNotify                uint16 = 10080
	ProtoIDS1BattleProto2022Energy2019StateReport               uint16 = 10081
	ProtoIDS1BattleProto2022Energy2019SetArmLight               uint16 = 10082
	ProtoIDS1BattleProto2022Energy2020SetArmRotate              uint16 = 10083
	ProtoIDS1BattleProto2022Energy2019ArmorHitUpload            uint16 = 10084
	ProtoIDS1BattleProto2022Energy2019AmorSelfCHeck             uint16 = 10085
	ProtoIDS1BattleProto2022Env_Heart_Beat_Req                  uint16 = 10086
	ProtoIDS1BattleProto2022Env_Heart_Beat_Ack                  uint16 = 10087
	ProtoIDS1BattleProto2022_IOCtrSetRmMotorsCfg                uint16 = 10088
	ProtoIDS1BattleProto2022_IOCtrSetRmMotorsCfgAck             uint16 = 10089
	ProtoIDS1BattleProto2022_IOCtrSetRmMotorsMoveVal            uint16 = 10090
	ProtoIDS1BattleProto2022_IOCtrRmMotorsStatus                uint16 = 10091
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType1             uint16 = 10092
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType1Ack          uint16 = 10093
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType2             uint16 = 10094
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType2Ack          uint16 = 10095
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType3             uint16 = 10096
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType3Ack          uint16 = 10097
	ProtoIDS1BattleProto2022_IOCtrSetActuator                   uint16 = 10098
	ProtoIDS1BattleProto2022_IOCtrSetActuatorAck                uint16 = 10099
	ProtoIDS1BattleProto2022_SiloEnvStatus                      uint16 = 10100
	ProtoIDS1BattleProto2022_SiloEnvCtrlDoor                    uint16 = 10101
	ProtoIDS1BattleProto2022ClientSiloEnvCtrReq                 uint16 = 10102
	ProtoIDS1BattleProto2022SendDartRobotStatus                 uint16 = 10103
	ProtoIDS1BattleProto2022SiloCtrLedData                      uint16 = 10104
	ProtoIDS1BattleProto2022ClientBuyBulletReq                  uint16 = 10105
	ProtoIDS1BattleProto2022ClientBuyBulletAck                  uint16 = 10106
	ProtoIDS1BattleProto2022ClientEconomyNotify                 uint16 = 10107
	ProtoIDS1BattleProto2022ClientAirplaneCtrlReq               uint16 = 10108
	ProtoIDS1BattleProto2022ClientAirplaneStatusInfo            uint16 = 10109
	ProtoIDS1BattleProto2022ClientShowMessage                   uint16 = 10110
	ProtoIDS1BattleProto2022ClientTalentDataReq                 uint16 = 10111
	ProtoIDS1BattleProto2022ClientTalentDataAck                 uint16 = 10112
	ProtoIDS1BattleProto2022ClientTalentDataNotify              uint16 = 10113
	ProtoIDS1BattleProto2022SetSupplyStateReq                   uint16 = 10114
	ProtoIDS1BattleProto2022SupplyClearScaleAck                 uint16 = 10115
	ProtoIDS1BattleProto2022SupplyClearScaleReq                 uint16 = 10116
	ProtoIDS1BattleProto2022SupplyExportAck                     uint16 = 10117
	ProtoIDS1BattleProto2022SupplyExportReq                     uint16 = 10118
	ProtoIDS1BattleProto2022SupplyFreedAck                      uint16 = 10119
	ProtoIDS1BattleProto2022SupplyFreedReq                      uint16 = 10120
	ProtoIDS1BattleProto2022SupplyReloadAck                     uint16 = 10121
	ProtoIDS1BattleProto2022SupplyReloadReq                     uint16 = 10122
	ProtoIDS1BattleProto2022SupplyReport                        uint16 = 10123
	ProtoIDS1BattleProto2022SupplyResetReq                      uint16 = 10124
	ProtoIDS1BattleProto2022SupplyRestartReq                    uint16 = 10125
	ProtoIDS1BattleProto2022ClientSupplySync                    uint16 = 10126
	ProtoIDS1BattleProto2022SetFlyCatStatus                     uint16 = 10127
	ProtoIDS1BattleProto2022SetFlyCatStatusAck                  uint16 = 10128
	ProtoIDS1BattleProto2022FlyCatStatus                        uint16 = 10129
	ProtoIDS1BattleProto2022SetFlyCatLight                      uint16 = 10130
	ProtoIDS1BattleProto2022SetFlyCatLightAck                   uint16 = 10131
	ProtoIDS1BattleProto2022SiloDevStatusSyncData               uint16 = 10132
	ProtoIDS1BattleProto2022ClientWarningNotify                 uint16 = 10133
	ProtoIDS1BattleProto2022ClientSHIELDDATA                    uint16 = 10134
	ProtoIDS1BattleProto2022ClientRobotDeathNotify              uint16 = 10135
	ProtoIDS1BattleProto2022BaseLightingEffect                  uint16 = 10136
	ProtoIDS1BattleProto2022RobotIRCheckReq                     uint16 = 10137
	ProtoIDS1BattleProto2022EnergySetID                         uint16 = 10138
	ProtoIDS1BattleProto2022EnerySetLogoLight                   uint16 = 10139
	ProtoIDS1BattleProto2022EnergyStateChangeNotify             uint16 = 10140
	ProtoIDS1BattleProto2022EnergyStateSync                     uint16 = 10141
	ProtoIDS1BattleProto2022QueryRobotInfo                      uint16 = 10142
	ProtoIDS1BattleProto2022QueryRobotInfoResult                uint16 = 10143
	ProtoIDS1BattleProto2022ClientKickOffRobotNotify            uint16 = 10144
	ProtoIDS1BattleProto2022RobotResurgenceNotify               uint16 = 10145
	ProtoIDS1BattleProto2022GripperStateNotify                  uint16 = 10146
	ProtoIDS1BattleProto2022BaseStatus                          uint16 = 10147
	ProtoIDS1BattleProto2022ClientBaseRobotDevStatusSyncData    uint16 = 10148
	ProtoIDS1BattleProto2022EnergyReset                         uint16 = 10149
	ProtoIDS1BattleProto2022ExchangeProCtrCMD                   uint16 = 10150
	ProtoIDS1BattleProto2022ExchangeProCtrCMDAck                uint16 = 10151
	ProtoIDS1BattleProto2022ExchangeProLidarInfo                uint16 = 10152
	ProtoIDS1BattleProto2022ExchangeProPosition                 uint16 = 10153
	ProtoIDS1BattleProto2022ExchangeProRealPosition             uint16 = 10154
	ProtoIDS1BattleProto2022ClientClientStatusSync              uint16 = 10155
	ProtoIDS1BattleProto2022ClientAllClientStatusSync           uint16 = 10156
	ProtoIDS1BattleProto2022ClientOutpostDeviceStatusSync       uint16 = 10157
	ProtoIDS1BattleProto2022ClientMineralInfoSync               uint16 = 10158
	ProtoIDS1BattleProto2022ClientMineralExchangeNotify         uint16 = 10159
	ProtoIDS1BattleProto2022ClientSiloEnvDoorOpenCloseNotify    uint16 = 10160
	ProtoIDS1BattleProto2022ClientFlycatStatusSync              uint16 = 10161
	ProtoIDS1BattleProto2022StuGameStatus                       uint16 = 10162
	ProtoIDS1BattleProto2022StuGameResult                       uint16 = 10163
	ProtoIDS1BattleProto2022StuRobotCurrentHP                   uint16 = 10164
	ProtoIDS1BattleProto2022StuICRABuffDebuffZoneStatus         uint16 = 10165
	ProtoIDS1BattleProto2022StuEnvStatus                        uint16 = 10166
	ProtoIDS1BattleProto2022StuSupplyStatus                     uint16 = 10167
	ProtoIDS1BattleProto2022StuWarning                          uint16 = 10168
	ProtoIDS1BattleProto2022StuMissileRemainingTime             uint16 = 10169
	ProtoIDS1BattleProto2022StuRobotBuff                        uint16 = 10170
	ProtoIDS1BattleProto2022StuAreialEnergy                     uint16 = 10171
	ProtoIDS1BattleProto2022StuRobotHurt                        uint16 = 10172
	ProtoIDS1BattleProto2022StuLeftBullet                       uint16 = 10173
	ProtoIDS1BattleProto2022StuRFIDStatus                       uint16 = 10174
	ProtoIDS1BattleProto2022StuSiloInfo                         uint16 = 10175
	ProtoIDS1BattleProto2022StuCustomControlData                uint16 = 10176
	ProtoIDS1BattleProto2022StuClientSendCMD                    uint16 = 10177
	ProtoIDS1BattleProto2022StuClientRecvInfo                   uint16 = 10178
	ProtoIDS1BattleProto2022StuCommunicationByteData            uint16 = 10179
	ProtoIDS1BattleProto2022StuCommunication                    uint16 = 10180
	ProtoIDS1BattleProto2022ClientRaderInfoToClient             uint16 = 10181
	ProtoIDS1BattleProto2022ClientCustomControlData             uint16 = 10182
	ProtoIDS1BattleProto2022MapClickInfoNotify                  uint16 = 10183
	ProtoIDS1BattleProto2022RobotNtpServerIpReq                 uint16 = 10184
	ProtoIDS1BattleProto2022RobotNtpServerIpAck                 uint16 = 10185
	ProtoIDS1BattleProto2022RobotMoudleLife                     uint16 = 10186
	ProtoIDS1BattleProto2022ClientCheckInTimeStampNotify        uint16 = 10187
	ProtoIDS1BattleProto2022RobotCheckinTimeStampReq            uint16 = 10188
	ProtoIDS1BattleProto2022RobotCheckinTimeStampAck            uint16 = 10189
	ProtoIDS1BattleProto2022RobotMeasureStartReq                uint16 = 10190
	ProtoIDS1BattleProto2022RobotMeasureStartRsp                uint16 = 10191
	ProtoIDS1BattleProto2022RobotMeasureStopReq                 uint16 = 10192
	ProtoIDS1BattleProto2022RobotMeasureStopRsp                 uint16 = 10193
	ProtoIDS1BattleProto2022RobotPushCAPInfo                    uint16 = 10194
	ProtoIDS1BattleProto2022RobotPushCAPRTInfo                  uint16 = 10195
	ProtoIDS1BattleProto2022RobotPushSensorInfo                 uint16 = 10196
	ProtoIDS1BattleProto2022RobotGetCheckCapReq                 uint16 = 10197
	ProtoIDS1BattleProto2022RobotGetCheckCapAck                 uint16 = 10198
	ProtoIDS1BattleProto2022ClientExceptionCapDataNotify        uint16 = 10199
	ProtoIDS1BattleProto2022ClientExceptionRecoverCardNotify    uint16 = 10200
	ProtoIDS1BattleProto2022ClientRecoverCardStatus             uint16 = 10201
	ProtoIDS1BattleProto2022ClientVtmSpeedSync                  uint16 = 10202
	ProtoIDS1BattleProto2022ClientServerMapSync                 uint16 = 10203
	ProtoIDS1BattleProto2022ClientServerUIMessage               uint16 = 10204
	ProtoIDS1BattleProto2022RobotBaseDevCtlCmdAck               uint16 = 10205
	ProtoIDS1BattleProto2022ClientWeaponFireNotify              uint16 = 10206
	ProtoIDS1BattleProto2022RobotVTMSetting                     uint16 = 10207
	ProtoIDS1BattleProto2022ClientExchangeProStateNotify        uint16 = 10208
	ProtoIDS1BattleProto2022RobotNewHeartBeatReq                uint16 = 10209
	ProtoIDS1BattleProto2022RobotNewHeartBeatAck                uint16 = 10210
	ProtoIDS1BattleProto2022_IOCtrCfgReq                        uint16 = 10211
	ProtoIDS1BattleProto2022ExchangeProClearOreRes              uint16 = 10212
	ProtoIDS1BattleProto2022ClientHoldCenterPointCompleteNotify uint16 = 10213
	ProtoIDS1BattleProto_Base_Armor_State_Notify                uint16 = 10214
	ProtoIDS1BattleProto2023ClientServerRobotsBattleInfoSync    uint16 = 10215
	ProtoIDS1BattleProto2023ClientTeamSupplyInfoSync            uint16 = 10216
	ProtoIDS1BattleProto2022EnergyArmorLife                     uint16 = 10217
	ProtoIDS1BattleProto2022EnergyArmorLifeAck                  uint16 = 10218
	ProtoIDS1BattleProto2023_OutPostArmorResetReq               uint16 = 10219
	ProtoIDS1BattleProto2023_OutPostArmorResetAck               uint16 = 10220
	ProtoIDS1BattleProto2023_OutPostArmorResetFinish            uint16 = 10221
	ProtoIDS1BattleProto2023ClientBuyReviveReq                  uint16 = 10222
	ProtoIDS1BattleProto2023ClientBuyReviveAck                  uint16 = 10223
	ProtoIDS1BattleProto2023ClientGameSysUploadResultNotify     uint16 = 10224
	ProtoIDS1BattleProto2023ClientTeamControlZoneInfoSync       uint16 = 10225
	ProtoIDS1BattleProto2023ClientGuardInPatrolInfoSync         uint16 = 10226
	ProtoIDS1BattleProto2023ClientRobotPathPlanInfo             uint16 = 10227
	ProtoIDS1BattleProto2023ClientSingleLadderRewardNotify      uint16 = 10228
	ProtoIDS1BattleProto2022StuTeamRobotsPos                    uint16 = 10229
	ProtoIDS1BattleProto2023StuRadarMarkProgress                uint16 = 10230
	ProtoIDS1BattleProto2022ClientHitRuneInfoSync               uint16 = 10231
	ProtoIDS1BattleProto2022ClientHitRuneNotifySync             uint16 = 10232
	ProtoIDS1BattleProto2022ExchangeProSetArmRunSpeed           uint16 = 10233
	ProtoIDS1BattleProto2022ExchangeProSetArmRunSpeedAck        uint16 = 10234
	ProtoIDS1BattleProto2023ExchangeProErrorInfo                uint16 = 10235
	ProtoIDS1BattleProto2023ExchangeProErrorInfoNotify          uint16 = 10236
	ProtoIDS1BattleProto2023CustomControlDataLenInfo            uint16 = 10237
	ProtoIDS1BattleProto2023BigRuneHitRecordNotify              uint16 = 10238
	ProtoIDS1BattleProto2023ClientPenaltyTableInfos             uint16 = 10239
	ProtoIDS1BattleProto2023ClientPenaltyTableInfosAck          uint16 = 10240
	ProtoIDS1BattleProto2023ClientBeAttackedDamageNotify        uint16 = 10241
	ProtoIDS1BattleProto2022ClientSeverAlertNotify              uint16 = 10242
	ProtoIDS1BattleProto2023EnergyAngleReport                   uint16 = 10243
	ProtoIDS1BattleProto2023EnergyTestModeSet                   uint16 = 10244
	ProtoIDS1BattleProto2022ClientRobotArmorLifeQueryResult     uint16 = 10245
	ProtoIDS1BattleProto2022ClientArmorLifeInfo                 uint16 = 10246
	ProtoIDS1BattleProto2023RobotArmorLife                      uint16 = 10247
	ProtoIDS1BattleProto2023IRArmorLightCtrl                    uint16 = 10248
	ProtoIDS1BattleProto2023IRArmorSelfCheck                    uint16 = 10249
	ProtoIDS1BattleProto2023IRArmorSelfCheckAck                 uint16 = 10250
	ProtoIDS1BattleProto2022BaseDartTargetControl               uint16 = 10251
	ProtoIDS1BattleProto2022BaseDartTargetControlAck            uint16 = 10252
	ProtoIDS1BattleProto2023ClientRobotCustomMessage            uint16 = 10253
	ProtoIDS1BattleProto2023ExchangeProMineralValueSync         uint16 = 10254
	ProtoIDS1BattleProto2022ClientMapClickTransResult           uint16 = 10255
	ProtoIDS1BattleProto2022StuGuardDecision                    uint16 = 10256
	ProtoIDS1BattleProto2022StuGuardDecisionInfo                uint16 = 10257
	ProtoIDS1BattleProto2022StuRadarDecision                    uint16 = 10258
	ProtoIDS1BattleProto2022StuRadarDecisionInfo                uint16 = 10259
	ProtoIDS1BattleProto2022ClientRobotExpAddData               uint16 = 10260
	ProtoIDS1BattleProto2023_OutPostStatus                      uint16 = 10261
	ProtoIDS1BattleProto2023ClientModuleVersionState            uint16 = 10262
	ProtoIDS1BattleProto2023ClientInitRuleParms                 uint16 = 10263
	ProtoIDS1BattleProto2022StuRadarMarkPositionInfo            uint16 = 10264
	ProtoIDS1BattleProto2023ClientGameSystemInfoNotify          uint16 = 10265
	ProtoIDS1BattleProto2023ClientChangeTokenCMD                uint16 = 10266
)

var ProtoIdMap = map[uint16]string{
	9000:  "ProtoIDS1ProtoLoginReq",
	9001:  "ProtoIDS1ProtoLoginAck",
	9002:  "ProtoIDS1ProtoReLoginReq",
	9003:  "ProtoIDS1ProtoReLoginAck",
	9004:  "ProtoIDS1ProtoReLoginRoomNotify",
	9005:  "ProtoIDS1ProtoLogoutReq",
	9006:  "ProtoIDS1ProtoLogoutAck",
	9007:  "ProtoIDS1ProtoLogoutNotify",
	9008:  "ProtoIDS1ProtoHeartBeatReq",
	9009:  "ProtoIDS1ProtoHeartBeatAck",
	9010:  "ProtoIDS1ProtoTestReq",
	9011:  "ProtoIDS1ProtoTestAck",
	9012:  "ProtoIDS1ProtoPingReq",
	9013:  "ProtoIDS1ProtoPingAck",
	9014:  "ProtoIDS1ProtoEnterRoomReq",
	9015:  "ProtoIDS1ProtoEnterRoomAck",
	9016:  "ProtoIDS1ProtoEnterRoomNotify",
	9017:  "ProtoIDS1ProtoRoomInfosNotify",
	9018:  "ProtoIDS1ProtoRoomStateNotify",
	9019:  "ProtoIDS1ProtoRoomPauseNotify",
	9020:  "ProtoIDS1ProtoLeaveRoomReq",
	9021:  "ProtoIDS1ProtoLeaveRoomAck",
	9022:  "ProtoIDS1ProtoLeaveRoomNotify",
	9023:  "ProtoIDS1ProtoSetTeamReq",
	9024:  "ProtoIDS1ProtoSetTeamAck",
	9025:  "ProtoIDS1ProtoSetTeamNotify",
	9026:  "ProtoIDS1ProtoSetTidReq",
	9027:  "ProtoIDS1ProtoSetTidAck",
	9028:  "ProtoIDS1ProtoSetTidNotify",
	9029:  "ProtoIDS1ProtoSetReadyReq",
	9030:  "ProtoIDS1ProtoSetReadyAck",
	9031:  "ProtoIDS1ProtoSetReadyNotify",
	9032:  "ProtoIDS1ProtoSetTokenReq",
	9033:  "ProtoIDS1ProtoSetTokenAck",
	9034:  "ProtoIDS1ProtoSetSvrInfoReq",
	9035:  "ProtoIDS1ProtoSetSvrInfoAck",
	9036:  "ProtoIDS1ProtoSetWifiInfoNotify",
	9037:  "ProtoIDS1ProtoSetLockSeatNotify",
	9038:  "ProtoIDS1ProtoSetCurrMatchNotify",
	9039:  "ProtoIDS1ProtoSetTeamPlayerNumNotify",
	9040:  "ProtoIDS1ProtoGMCommandReq",
	9041:  "ProtoIDS1ProtoGMCommandAck",
	9042:  "ProtoIDS1ProtoPlayerCommand_UserInput_Notify",
	9043:  "ProtoIDS1ProtoPlayerCommand_Disconnect_Notify",
	9044:  "ProtoIDS1ProtoClientSyncReq",
	9045:  "ProtoIDS1ProtoClientSyncNotify",
	9046:  "ProtoIDS1ProtoTeamInfoNotify",
	9047:  "ProtoIDS1ProtoTeamLogoNotify",
	9048:  "ProtoIDS1ProtoTeamCampNotify",
	9049:  "ProtoIDS1ProtoSvrStateReq",
	9050:  "ProtoIDS1ProtoSvrStateAck",
	9051:  "ProtoIDS1ProtoUserStateReq",
	9052:  "ProtoIDS1ProtoUserStateAck",
	9053:  "ProtoIDS1ClientTransferRobotProto",
	9054:  "ProtoIDS1ProtoRobotNetworkStatusReq",
	9055:  "ProtoIDS1ProtoRobotNetworkStatusAck",
	9056:  "ProtoIDS1ProtoClientNetworkInfoNotify",
	9057:  "ProtoIDS1ProtoRobotStudentSerialPortInfos",
	9058:  "ProtoIDS1ProtoNetworkInfoNotify",
	9059:  "ProtoIDS1ProtoTechnicalPauseInfoNotify",
	9060:  "ProtoIDS1ProtoMarkDetectResult",
	9061:  "ProtoIDS1ProtoMarkDetectResultAck",
	9062:  "ProtoIDS1ProtoClientExceptionInfo",
	10000: "ProtoIDS1BattleProtoAutoStateNotify",
	10001: "ProtoIDS1BattleProtoArmorHitReq",
	10002: "ProtoIDS1BattleProtoArmorHitAck",
	10003: "ProtoIDS1BattleProtoBeAttackNotify",
	10004: "ProtoIDS1BattleProtoGunFireReq",
	10005: "ProtoIDS1BattleProtoGunFireAck",
	10006: "ProtoIDS1BattleProtoGunFireNotify",
	10007: "ProtoIDS1BattleProtoGunHeatNotify",
	10008: "ProtoIDS1BattleProtoGunBulletNotify",
	10009: "ProtoIDS1BattleProtoAttrsNotify",
	10010: "ProtoIDS1BattleProtoPlayerInitOkNotify",
	10011: "ProtoIDS1BattleProtoPlayerDeadNotify",
	10012: "ProtoIDS1BattleProtoPlayerReliveNotify",
	10013: "ProtoIDS1BattleProtoBuffAddNotify",
	10014: "ProtoIDS1BattleProtoBuffDelNotify",
	10015: "ProtoIDS1BattleProtoBuffModifyNotify",
	10016: "ProtoIDS1BattleProtoMarkerDetectReq",
	10017: "ProtoIDS1BattleProtoMarkerDetectAck",
	10018: "ProtoIDS1BattleProtoProgressNotify",
	10019: "ProtoIDS1BattleProtoDeviceModuleReq",
	10020: "ProtoIDS1BattleProtoDeviceModuleNotify",
	10021: "ProtoIDS1BattleProtoPoJiaNotify",
	10022: "ProtoIDS1BattleProtoPlayerCommand_GunFire_Notify",
	10023: "ProtoIDS1BattleProtoPlayerCommand_Move_Notify",
	10024: "ProtoIDS1BattleProtoPlayerCommand_LockScreen_Notify",
	10025: "ProtoIDS1BattleProtoPlayerCommand_AutoControl_Notify",
	10026: "ProtoIDS1BattleProtoPlayerCommand_UserInput_Notify",
	10027: "ProtoIDS1BattleProtoPlayerCommand_UserCustomAction_Notify",
	10028: "ProtoIDS1BattleProtoPlayerCommand_SetDisconnectNotify",
	10029: "ProtoIDS1BattleProtoSceneInfosNotify",
	10030: "ProtoIDS1BattleProtoSceneStateNotify",
	10031: "ProtoIDS1BattleProtoGameSettlementNotify",
	10032: "ProtoIDS1BattleProtoReloginNotify",
	10033: "ProtoIDS1BattleProtoIsCanUseATKNotify",
	10034: "ProtoIDS1BattleProtoIsCanUsePJNotify",
	10035: "ProtoIDS1BattleProtoClientSyncReq",
	10036: "ProtoIDS1BattleProtoClientSyncNotify",
	10037: "ProtoIDS1BattleProtoDeviceBulletNotify",
	10038: "ProtoIDS1BattleProtoWarningNotify",
	10039: "ProtoIDS1BattleProtoPlayerResetNotify",
	10040: "ProtoIDS1BattleProtoReduceReliveTimeNotify",
	10041: "ProtoIDS1BattleProtoFullSceneDataReq",
	10042: "ProtoIDS1BattleProtoTeamInfoNotify",
	10043: "ProtoIDS1BattleProtoPlaceholderNotify",
	10044: "ProtoIDS1BattleProtoModuleOffLineInfoNotify",
	10045: "ProtoIDS1BattleProtoAddBulletNotify",
	10046: "ProtoIDS1BattleProtoShareBulletReq",
	10047: "ProtoIDS1BattleProtoUploadModuleInfo",
	10048: "ProtoIDS1BattleProto2022RobotInitCfgTab",
	10049: "ProtoIDS1BattleProto2022ConfigTabAck",
	10050: "ProtoIDS1BattleProto2022RobotStatus",
	10051: "ProtoIDS1BattleProto2022RobotsDataSync",
	10052: "ProtoIDS1BattleProto2022RobotsWeaponFire",
	10053: "ProtoIDS1BattleProto2022RobotHited",
	10054: "ProtoIDS1BattleProto2022VtmSpeedSync",
	10055: "ProtoIDS1BattleProto2022ClientsServerBaseSync",
	10056: "ProtoIDS1BattleProto2022ClientsServerSync",
	10057: "ProtoIDS1BattleProto2022ClientsFirstSync",
	10058: "ProtoIDS1BattleProto2022ClientsRobotStatusNotify",
	10059: "ProtoIDS1BattleProto2022ClientsRobotModuleErrorInfo",
	10060: "ProtoIDS1BattleProto2022TestTimeDelayDownData",
	10061: "ProtoIDS1BattleProto2022TestTimeDelayDownComplete",
	10062: "ProtoIDS1BattleProto2022TestDelayDownDataAck",
	10063: "ProtoIDS1BattleProto2022CONTROL_CMD_T",
	10064: "ProtoIDS1BattleProto2022ClientGMCommand",
	10065: "ProtoIDS1BattleProto2022ClientCurrentProgress",
	10066: "ProtoIDS1BattleProto2022ClientHitNotify",
	10067: "ProtoIDS1BattleProto2022ReLoginDataSync",
	10068: "ProtoIDS1BattleProto2022ClientResultDataInfo",
	10069: "ProtoIDS1BattleProto2022EnvBaseShellControl",
	10070: "ProtoIDS1BattleProto2022RFIDReq",
	10071: "ProtoIDS1BattleProto2022RFIDNotify",
	10072: "ProtoIDS1BattleProto2022IOCtrCfgSet",
	10073: "ProtoIDS1BattleProto2022IOCtrCfgSetAck",
	10074: "ProtoIDS1BattleProto2022IOCtrSetVal",
	10075: "ProtoIDS1BattleProto2022IOCtrSetValAck",
	10076: "ProtoIDS1BattleProto2022IOCtrState",
	10077: "ProtoIDS1BattleProto2022IOCtrTriggerVal",
	10078: "ProtoIDS1BattleProto2022EnvDevicesDescripeReq",
	10079: "ProtoIDS1BattleProto2022EnvDevicesDescripeAck",
	10080: "ProtoIDS1BattleProto2022S2C_PowerStateNotify",
	10081: "ProtoIDS1BattleProto2022Energy2019StateReport",
	10082: "ProtoIDS1BattleProto2022Energy2019SetArmLight",
	10083: "ProtoIDS1BattleProto2022Energy2020SetArmRotate",
	10084: "ProtoIDS1BattleProto2022Energy2019ArmorHitUpload",
	10085: "ProtoIDS1BattleProto2022Energy2019AmorSelfCHeck",
	10086: "ProtoIDS1BattleProto2022Env_Heart_Beat_Req",
	10087: "ProtoIDS1BattleProto2022Env_Heart_Beat_Ack",
	10088: "ProtoIDS1BattleProto2022_IOCtrSetRmMotorsCfg",
	10089: "ProtoIDS1BattleProto2022_IOCtrSetRmMotorsCfgAck",
	10090: "ProtoIDS1BattleProto2022_IOCtrSetRmMotorsMoveVal",
	10091: "ProtoIDS1BattleProto2022_IOCtrRmMotorsStatus",
	10092: "ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType1",
	10093: "ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType1Ack",
	10094: "ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType2",
	10095: "ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType2Ack",
	10096: "ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType3",
	10097: "ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType3Ack",
	10098: "ProtoIDS1BattleProto2022_IOCtrSetActuator",
	10099: "ProtoIDS1BattleProto2022_IOCtrSetActuatorAck",
	10100: "ProtoIDS1BattleProto2022_SiloEnvStatus",
	10101: "ProtoIDS1BattleProto2022_SiloEnvCtrlDoor",
	10102: "ProtoIDS1BattleProto2022ClientSiloEnvCtrReq",
	10103: "ProtoIDS1BattleProto2022SendDartRobotStatus",
	10104: "ProtoIDS1BattleProto2022SiloCtrLedData",
	10105: "ProtoIDS1BattleProto2022ClientBuyBulletReq",
	10106: "ProtoIDS1BattleProto2022ClientBuyBulletAck",
	10107: "ProtoIDS1BattleProto2022ClientEconomyNotify",
	10108: "ProtoIDS1BattleProto2022ClientAirplaneCtrlReq",
	10109: "ProtoIDS1BattleProto2022ClientAirplaneStatusInfo",
	10110: "ProtoIDS1BattleProto2022ClientShowMessage",
	10111: "ProtoIDS1BattleProto2022ClientTalentDataReq",
	10112: "ProtoIDS1BattleProto2022ClientTalentDataAck",
	10113: "ProtoIDS1BattleProto2022ClientTalentDataNotify",
	10114: "ProtoIDS1BattleProto2022SetSupplyStateReq",
	10115: "ProtoIDS1BattleProto2022SupplyClearScaleAck",
	10116: "ProtoIDS1BattleProto2022SupplyClearScaleReq",
	10117: "ProtoIDS1BattleProto2022SupplyExportAck",
	10118: "ProtoIDS1BattleProto2022SupplyExportReq",
	10119: "ProtoIDS1BattleProto2022SupplyFreedAck",
	10120: "ProtoIDS1BattleProto2022SupplyFreedReq",
	10121: "ProtoIDS1BattleProto2022SupplyReloadAck",
	10122: "ProtoIDS1BattleProto2022SupplyReloadReq",
	10123: "ProtoIDS1BattleProto2022SupplyReport",
	10124: "ProtoIDS1BattleProto2022SupplyResetReq",
	10125: "ProtoIDS1BattleProto2022SupplyRestartReq",
	10126: "ProtoIDS1BattleProto2022ClientSupplySync",
	10127: "ProtoIDS1BattleProto2022SetFlyCatStatus",
	10128: "ProtoIDS1BattleProto2022SetFlyCatStatusAck",
	10129: "ProtoIDS1BattleProto2022FlyCatStatus",
	10130: "ProtoIDS1BattleProto2022SetFlyCatLight",
	10131: "ProtoIDS1BattleProto2022SetFlyCatLightAck",
	10132: "ProtoIDS1BattleProto2022SiloDevStatusSyncData",
	10133: "ProtoIDS1BattleProto2022ClientWarningNotify",
	10134: "ProtoIDS1BattleProto2022ClientSHIELDDATA",
	10135: "ProtoIDS1BattleProto2022ClientRobotDeathNotify",
	10136: "ProtoIDS1BattleProto2022BaseLightingEffect",
	10137: "ProtoIDS1BattleProto2022RobotIRCheckReq",
	10138: "ProtoIDS1BattleProto2022EnergySetID",
	10139: "ProtoIDS1BattleProto2022EnerySetLogoLight",
	10140: "ProtoIDS1BattleProto2022EnergyStateChangeNotify",
	10141: "ProtoIDS1BattleProto2022EnergyStateSync",
	10142: "ProtoIDS1BattleProto2022QueryRobotInfo",
	10143: "ProtoIDS1BattleProto2022QueryRobotInfoResult",
	10144: "ProtoIDS1BattleProto2022ClientKickOffRobotNotify",
	10145: "ProtoIDS1BattleProto2022RobotResurgenceNotify",
	10146: "ProtoIDS1BattleProto2022GripperStateNotify",
	10147: "ProtoIDS1BattleProto2022BaseStatus",
	10148: "ProtoIDS1BattleProto2022ClientBaseRobotDevStatusSyncData",
	10149: "ProtoIDS1BattleProto2022EnergyReset",
	10150: "ProtoIDS1BattleProto2022ExchangeProCtrCMD",
	10151: "ProtoIDS1BattleProto2022ExchangeProCtrCMDAck",
	10152: "ProtoIDS1BattleProto2022ExchangeProLidarInfo",
	10153: "ProtoIDS1BattleProto2022ExchangeProPosition",
	10154: "ProtoIDS1BattleProto2022ExchangeProRealPosition",
	10155: "ProtoIDS1BattleProto2022ClientClientStatusSync",
	10156: "ProtoIDS1BattleProto2022ClientAllClientStatusSync",
	10157: "ProtoIDS1BattleProto2022ClientOutpostDeviceStatusSync",
	10158: "ProtoIDS1BattleProto2022ClientMineralInfoSync",
	10159: "ProtoIDS1BattleProto2022ClientMineralExchangeNotify",
	10160: "ProtoIDS1BattleProto2022ClientSiloEnvDoorOpenCloseNotify",
	10161: "ProtoIDS1BattleProto2022ClientFlycatStatusSync",
	10162: "ProtoIDS1BattleProto2022StuGameStatus",
	10163: "ProtoIDS1BattleProto2022StuGameResult",
	10164: "ProtoIDS1BattleProto2022StuRobotCurrentHP",
	10165: "ProtoIDS1BattleProto2022StuICRABuffDebuffZoneStatus",
	10166: "ProtoIDS1BattleProto2022StuEnvStatus",
	10167: "ProtoIDS1BattleProto2022StuSupplyStatus",
	10168: "ProtoIDS1BattleProto2022StuWarning",
	10169: "ProtoIDS1BattleProto2022StuMissileRemainingTime",
	10170: "ProtoIDS1BattleProto2022StuRobotBuff",
	10171: "ProtoIDS1BattleProto2022StuAreialEnergy",
	10172: "ProtoIDS1BattleProto2022StuRobotHurt",
	10173: "ProtoIDS1BattleProto2022StuLeftBullet",
	10174: "ProtoIDS1BattleProto2022StuRFIDStatus",
	10175: "ProtoIDS1BattleProto2022StuSiloInfo",
	10176: "ProtoIDS1BattleProto2022StuCustomControlData",
	10177: "ProtoIDS1BattleProto2022StuClientSendCMD",
	10178: "ProtoIDS1BattleProto2022StuClientRecvInfo",
	10179: "ProtoIDS1BattleProto2022StuCommunicationByteData",
	10180: "ProtoIDS1BattleProto2022StuCommunication",
	10181: "ProtoIDS1BattleProto2022ClientRaderInfoToClient",
	10182: "ProtoIDS1BattleProto2022ClientCustomControlData",
	10183: "ProtoIDS1BattleProto2022MapClickInfoNotify",
	10184: "ProtoIDS1BattleProto2022RobotNtpServerIpReq",
	10185: "ProtoIDS1BattleProto2022RobotNtpServerIpAck",
	10186: "ProtoIDS1BattleProto2022RobotMoudleLife",
	10187: "ProtoIDS1BattleProto2022ClientCheckInTimeStampNotify",
	10188: "ProtoIDS1BattleProto2022RobotCheckinTimeStampReq",
	10189: "ProtoIDS1BattleProto2022RobotCheckinTimeStampAck",
	10190: "ProtoIDS1BattleProto2022RobotMeasureStartReq",
	10191: "ProtoIDS1BattleProto2022RobotMeasureStartRsp",
	10192: "ProtoIDS1BattleProto2022RobotMeasureStopReq",
	10193: "ProtoIDS1BattleProto2022RobotMeasureStopRsp",
	10194: "ProtoIDS1BattleProto2022RobotPushCAPInfo",
	10195: "ProtoIDS1BattleProto2022RobotPushCAPRTInfo",
	10196: "ProtoIDS1BattleProto2022RobotPushSensorInfo",
	10197: "ProtoIDS1BattleProto2022RobotGetCheckCapReq",
	10198: "ProtoIDS1BattleProto2022RobotGetCheckCapAck",
	10199: "ProtoIDS1BattleProto2022ClientExceptionCapDataNotify",
	10200: "ProtoIDS1BattleProto2022ClientExceptionRecoverCardNotify",
	10201: "ProtoIDS1BattleProto2022ClientRecoverCardStatus",
	10202: "ProtoIDS1BattleProto2022ClientVtmSpeedSync",
	10203: "ProtoIDS1BattleProto2022ClientServerMapSync",
	10204: "ProtoIDS1BattleProto2022ClientServerUIMessage",
	10205: "ProtoIDS1BattleProto2022RobotBaseDevCtlCmdAck",
	10206: "ProtoIDS1BattleProto2022ClientWeaponFireNotify",
	10207: "ProtoIDS1BattleProto2022RobotVTMSetting",
	10208: "ProtoIDS1BattleProto2022ClientExchangeProStateNotify",
	10209: "ProtoIDS1BattleProto2022RobotNewHeartBeatReq",
	10210: "ProtoIDS1BattleProto2022RobotNewHeartBeatAck",
	10211: "ProtoIDS1BattleProto2022_IOCtrCfgReq",
	10212: "ProtoIDS1BattleProto2022ExchangeProClearOreRes",
	10213: "ProtoIDS1BattleProto2022ClientHoldCenterPointCompleteNotify",
	10214: "ProtoIDS1BattleProto_Base_Armor_State_Notify",
	10215: "ProtoIDS1BattleProto2023ClientServerRobotsBattleInfoSync",
	10216: "ProtoIDS1BattleProto2023ClientTeamSupplyInfoSync",
	10217: "ProtoIDS1BattleProto2022EnergyArmorLife",
	10218: "ProtoIDS1BattleProto2022EnergyArmorLifeAck",
	10219: "ProtoIDS1BattleProto2023_OutPostArmorResetReq",
	10220: "ProtoIDS1BattleProto2023_OutPostArmorResetAck",
	10221: "ProtoIDS1BattleProto2023_OutPostArmorResetFinish",
	10222: "ProtoIDS1BattleProto2023ClientBuyReviveReq",
	10223: "ProtoIDS1BattleProto2023ClientBuyReviveAck",
	10224: "ProtoIDS1BattleProto2023ClientGameSysUploadResultNotify",
	10225: "ProtoIDS1BattleProto2023ClientTeamControlZoneInfoSync",
	10226: "ProtoIDS1BattleProto2023ClientGuardInPatrolInfoSync",
	10227: "ProtoIDS1BattleProto2023ClientRobotPathPlanInfo",
	10228: "ProtoIDS1BattleProto2023ClientSingleLadderRewardNotify",
	10229: "ProtoIDS1BattleProto2022StuTeamRobotsPos",
	10230: "ProtoIDS1BattleProto2023StuRadarMarkProgress",
	10231: "ProtoIDS1BattleProto2022ClientHitRuneInfoSync",
	10232: "ProtoIDS1BattleProto2022ClientHitRuneNotifySync",
	10233: "ProtoIDS1BattleProto2022ExchangeProSetArmRunSpeed",
	10234: "ProtoIDS1BattleProto2022ExchangeProSetArmRunSpeedAck",
	10235: "ProtoIDS1BattleProto2023ExchangeProErrorInfo",
	10236: "ProtoIDS1BattleProto2023ExchangeProErrorInfoNotify",
	10237: "ProtoIDS1BattleProto2023CustomControlDataLenInfo",
	10238: "ProtoIDS1BattleProto2023BigRuneHitRecordNotify",
	10239: "ProtoIDS1BattleProto2023ClientPenaltyTableInfos",
	10240: "ProtoIDS1BattleProto2023ClientPenaltyTableInfosAck",
	10241: "ProtoIDS1BattleProto2023ClientBeAttackedDamageNotify",
	10242: "ProtoIDS1BattleProto2022ClientSeverAlertNotify",
	10243: "ProtoIDS1BattleProto2023EnergyAngleReport",
	10244: "ProtoIDS1BattleProto2023EnergyTestModeSet",
	10245: "ProtoIDS1BattleProto2022ClientRobotArmorLifeQueryResult",
	10246: "ProtoIDS1BattleProto2022ClientArmorLifeInfo",
	10247: "ProtoIDS1BattleProto2023RobotArmorLife",
	10248: "ProtoIDS1BattleProto2023IRArmorLightCtrl",
	10249: "ProtoIDS1BattleProto2023IRArmorSelfCheck",
	10250: "ProtoIDS1BattleProto2023IRArmorSelfCheckAck",
	10251: "ProtoIDS1BattleProto2022BaseDartTargetControl",
	10252: "ProtoIDS1BattleProto2022BaseDartTargetControlAck",
	10253: "ProtoIDS1BattleProto2023ClientRobotCustomMessage",
	10254: "ProtoIDS1BattleProto2023ExchangeProMineralValueSync",
	10255: "ProtoIDS1BattleProto2022ClientMapClickTransResult",
	10256: "ProtoIDS1BattleProto2022StuGuardDecision",
	10257: "ProtoIDS1BattleProto2022StuGuardDecisionInfo",
	10258: "ProtoIDS1BattleProto2022StuRadarDecision",
	10259: "ProtoIDS1BattleProto2022StuRadarDecisionInfo",
	10260: "ProtoIDS1BattleProto2022ClientRobotExpAddData",
	10261: "ProtoIDS1BattleProto2023_OutPostStatus",
	10262: "ProtoIDS1BattleProto2023ClientModuleVersionState",
	10263: "ProtoIDS1BattleProto2023ClientInitRuleParms",
	10264: "ProtoIDS1BattleProto2022StuRadarMarkPositionInfo",
	10265: "ProtoIDS1BattleProto2023ClientGameSystemInfoNotify",
	10266: "ProtoIDS1BattleProto2023ClientChangeTokenCMD",
}

var ProtoTypeMap = map[uint16]reflect.Type{
	9000:  reflect.TypeOf(S1ProtoLoginReq{}),
	9001:  reflect.TypeOf(S1ProtoLoginAck{}),
	9002:  reflect.TypeOf(S1ProtoReLoginReq{}),
	9003:  reflect.TypeOf(S1ProtoReLoginAck{}),
	9004:  reflect.TypeOf(S1ProtoReLoginRoomNotify{}),
	9005:  reflect.TypeOf(S1ProtoLogoutReq{}),
	9006:  reflect.TypeOf(S1ProtoLogoutAck{}),
	9007:  reflect.TypeOf(S1ProtoLogoutNotify{}),
	9008:  reflect.TypeOf(S1ProtoHeartBeatReq{}),
	9009:  reflect.TypeOf(S1ProtoHeartBeatAck{}),
	9010:  reflect.TypeOf(S1ProtoTestReq{}),
	9011:  reflect.TypeOf(S1ProtoTestAck{}),
	9012:  reflect.TypeOf(S1ProtoPingReq{}),
	9013:  reflect.TypeOf(S1ProtoPingAck{}),
	9014:  reflect.TypeOf(S1ProtoEnterRoomReq{}),
	9015:  reflect.TypeOf(S1ProtoEnterRoomAck{}),
	9016:  reflect.TypeOf(S1ProtoEnterRoomNotify{}),
	9017:  reflect.TypeOf(S1ProtoRoomInfosNotify{}),
	9018:  reflect.TypeOf(S1ProtoRoomStateNotify{}),
	9019:  reflect.TypeOf(S1ProtoRoomPauseNotify{}),
	9020:  reflect.TypeOf(S1ProtoLeaveRoomReq{}),
	9021:  reflect.TypeOf(S1ProtoLeaveRoomAck{}),
	9022:  reflect.TypeOf(S1ProtoLeaveRoomNotify{}),
	9023:  reflect.TypeOf(S1ProtoSetTeamReq{}),
	9024:  reflect.TypeOf(S1ProtoSetTeamAck{}),
	9025:  reflect.TypeOf(S1ProtoSetTeamNotify{}),
	9026:  reflect.TypeOf(S1ProtoSetTidReq{}),
	9027:  reflect.TypeOf(S1ProtoSetTidAck{}),
	9028:  reflect.TypeOf(S1ProtoSetTidNotify{}),
	9029:  reflect.TypeOf(S1ProtoSetReadyReq{}),
	9030:  reflect.TypeOf(S1ProtoSetReadyAck{}),
	9031:  reflect.TypeOf(S1ProtoSetReadyNotify{}),
	9032:  reflect.TypeOf(S1ProtoSetTokenReq{}),
	9033:  reflect.TypeOf(S1ProtoSetTokenAck{}),
	9034:  reflect.TypeOf(S1ProtoSetSvrInfoReq{}),
	9035:  reflect.TypeOf(S1ProtoSetSvrInfoAck{}),
	9036:  reflect.TypeOf(S1ProtoSetWifiInfoNotify{}),
	9037:  reflect.TypeOf(S1ProtoSetLockSeatNotify{}),
	9038:  reflect.TypeOf(S1ProtoSetCurrMatchNotify{}),
	9039:  reflect.TypeOf(S1ProtoSetTeamPlayerNumNotify{}),
	9040:  reflect.TypeOf(S1ProtoGmcommandReq{}),
	9041:  reflect.TypeOf(S1ProtoGmcommandAck{}),
	9042:  reflect.TypeOf(S1ProtoPlayerCommandUserInputNotify{}),
	9043:  reflect.TypeOf(S1ProtoPlayerCommandDisconnectNotify{}),
	9044:  reflect.TypeOf(S1ProtoClientSyncReq{}),
	9045:  reflect.TypeOf(S1ProtoClientSyncNotify{}),
	9046:  reflect.TypeOf(S1ProtoTeamInfoNotify{}),
	9047:  reflect.TypeOf(S1ProtoTeamLogoNotify{}),
	9048:  reflect.TypeOf(S1ProtoTeamCampNotify{}),
	9049:  reflect.TypeOf(S1ProtoSvrStateReq{}),
	9050:  reflect.TypeOf(S1ProtoSvrStateAck{}),
	9051:  reflect.TypeOf(S1ProtoUserStateReq{}),
	9052:  reflect.TypeOf(S1ProtoUserStateAck{}),
	9053:  reflect.TypeOf(S1ClientTransferRobotProto{}),
	9054:  reflect.TypeOf(S1ProtoRobotNetworkStatusReq{}),
	9055:  reflect.TypeOf(S1ProtoRobotNetworkStatusAck{}),
	9056:  reflect.TypeOf(S1ProtoClientNetworkInfoNotify{}),
	9057:  reflect.TypeOf(S1ProtoRobotStudentSerialPortInfos{}),
	9058:  reflect.TypeOf(S1ProtoNetworkInfoNotify{}),
	9059:  reflect.TypeOf(S1ProtoTechnicalPauseInfoNotify{}),
	9060:  reflect.TypeOf(S1ProtoMarkDetectResult{}),
	9061:  reflect.TypeOf(S1ProtoMarkDetectResultAck{}),
	9062:  reflect.TypeOf(S1ProtoClientExceptionInfo{}),
	10000: reflect.TypeOf(S1BattleProtoAutoStateNotify{}),
	10001: reflect.TypeOf(S1BattleProtoArmorHitReq{}),
	10002: reflect.TypeOf(S1BattleProtoArmorHitAck{}),
	10003: reflect.TypeOf(S1BattleProtoBeAttackNotify{}),
	10004: reflect.TypeOf(S1BattleProtoGunFireReq{}),
	10005: reflect.TypeOf(S1BattleProtoGunFireAck{}),
	10006: reflect.TypeOf(S1BattleProtoGunFireNotify{}),
	10007: reflect.TypeOf(S1BattleProtoGunHeatNotify{}),
	10008: reflect.TypeOf(S1BattleProtoGunBulletNotify{}),
	10009: reflect.TypeOf(S1BattleProtoAttrsNotify{}),
	10010: reflect.TypeOf(S1BattleProtoPlayerInitOkNotify{}),
	10011: reflect.TypeOf(S1BattleProtoPlayerDeadNotify{}),
	10012: reflect.TypeOf(S1BattleProtoPlayerReliveNotify{}),
	10013: reflect.TypeOf(S1BattleProtoBuffAddNotify{}),
	10014: reflect.TypeOf(S1BattleProtoBuffDelNotify{}),
	10015: reflect.TypeOf(S1BattleProtoBuffModifyNotify{}),
	10016: reflect.TypeOf(S1BattleProtoMarkerDetectReq{}),
	10017: reflect.TypeOf(S1BattleProtoMarkerDetectAck{}),
	10018: reflect.TypeOf(S1BattleProtoProgressNotify{}),
	10019: reflect.TypeOf(S1BattleProtoDeviceModuleReq{}),
	10020: reflect.TypeOf(S1BattleProtoDeviceModuleNotify{}),
	10021: reflect.TypeOf(S1BattleProtoPoJiaNotify{}),
	10022: reflect.TypeOf(S1BattleProtoPlayerCommandGunFireNotify{}),
	10023: reflect.TypeOf(S1BattleProtoPlayerCommandMoveNotify{}),
	10024: reflect.TypeOf(S1BattleProtoPlayerCommandLockScreenNotify{}),
	10025: reflect.TypeOf(S1BattleProtoPlayerCommandAutoControlNotify{}),
	10026: reflect.TypeOf(S1BattleProtoPlayerCommandUserInputNotify{}),
	10027: reflect.TypeOf(S1BattleProtoPlayerCommandUserCustomActionNotify{}),
	10028: reflect.TypeOf(S1BattleProtoPlayerCommandSetDisconnectNotify{}),
	10029: reflect.TypeOf(S1BattleProtoSceneInfosNotify{}),
	10030: reflect.TypeOf(S1BattleProtoSceneStateNotify{}),
	10031: reflect.TypeOf(S1BattleProtoGameSettlementNotify{}),
	10032: reflect.TypeOf(S1BattleProtoReloginNotify{}),
	10033: reflect.TypeOf(S1BattleProtoIsCanUseAtknotify{}),
	10034: reflect.TypeOf(S1BattleProtoIsCanUsePjnotify{}),
	10035: reflect.TypeOf(S1BattleProtoClientSyncReq{}),
	10036: reflect.TypeOf(S1BattleProtoClientSyncNotify{}),
	10037: reflect.TypeOf(S1BattleProtoDeviceBulletNotify{}),
	10038: reflect.TypeOf(S1BattleProtoWarningNotify{}),
	10039: reflect.TypeOf(S1BattleProtoPlayerResetNotify{}),
	10040: reflect.TypeOf(S1BattleProtoReduceReliveTimeNotify{}),
	10041: reflect.TypeOf(S1BattleProtoFullSceneDataReq{}),
	10042: reflect.TypeOf(S1BattleProtoTeamInfoNotify{}),
	10043: reflect.TypeOf(S1BattleProtoPlaceholderNotify{}),
	10044: reflect.TypeOf(S1BattleProtoModuleOffLineInfoNotify{}),
	10045: reflect.TypeOf(S1BattleProtoAddBulletNotify{}),
	10046: reflect.TypeOf(S1BattleProtoShareBulletReq{}),
	10047: reflect.TypeOf(S1BattleProtoUploadModuleInfo{}),
	10048: reflect.TypeOf(S1BattleProto2022RobotInitCfgTab{}),
	10049: reflect.TypeOf(S1BattleProto2022ConfigTabAck{}),
	10050: reflect.TypeOf(S1BattleProto2022RobotStatus{}),
	10051: reflect.TypeOf(S1BattleProto2022RobotsDataSync{}),
	//	10052: reflect.TypeOf(S1BattleProto2022RobotsWeaponFire{}),
	10053: reflect.TypeOf(S1BattleProto2022RobotHited{}),
	10054: reflect.TypeOf(S1BattleProto2022VtmSpeedSync{}),
	10055: reflect.TypeOf(S1BattleProto2022ClientsServerBaseSync{}),
	10056: reflect.TypeOf(S1BattleProto2022ClientsServerSync{}),
	10057: reflect.TypeOf(S1BattleProto2022ClientsFirstSync{}),
	//	10058: reflect.TypeOf(S1BattleProto2022ClientsRobotStatusNotify{}),
	10059: reflect.TypeOf(S1BattleProto2022ClientsRobotModuleErrorInfo{}),
	10060: reflect.TypeOf(S1BattleProto2022TestTimeDelayDownData{}),
	10061: reflect.TypeOf(S1BattleProto2022TestTimeDelayDownComplete{}),
	//	10062: reflect.TypeOf(S1BattleProto2022TestDelayDownDataAck{}),
	10063: reflect.TypeOf(S1BattleProto2022ControlCmdT{}),
	10064: reflect.TypeOf(S1BattleProto2022ClientGmcommand{}),
	10065: reflect.TypeOf(S1BattleProto2022ClientCurrentProgress{}),
	10066: reflect.TypeOf(S1BattleProto2022ClientHitNotify{}),
	10067: reflect.TypeOf(S1BattleProto2022ReLoginDataSync{}),
	10068: reflect.TypeOf(S1BattleProto2022ClientResultDataInfo{}),
	10069: reflect.TypeOf(S1BattleProto2022EnvBaseShellControl{}),
	//	10070: reflect.TypeOf(S1BattleProto2022Rfidreq{}),
	//	10071: reflect.TypeOf(S1BattleProto2022Rfidnotify{}),
	10072: reflect.TypeOf(S1BattleProto2022IoctrCfgSet{}),
	10073: reflect.TypeOf(S1BattleProto2022IoctrCfgSetAck{}),
	10074: reflect.TypeOf(S1BattleProto2022IoctrSetVal{}),
	10075: reflect.TypeOf(S1BattleProto2022IoctrSetValAck{}),
	10076: reflect.TypeOf(S1BattleProto2022IoctrState{}),
	10077: reflect.TypeOf(S1BattleProto2022IoctrTriggerVal{}),
	10078: reflect.TypeOf(S1BattleProto2022EnvDevicesDescripeReq{}),
	10079: reflect.TypeOf(S1BattleProto2022EnvDevicesDescripeAck{}),
	10080: reflect.TypeOf(S1BattleProto2022S2CPowerStateNotify{}),
	10081: reflect.TypeOf(S1BattleProto2022Energy2019StateReport{}),
	10082: reflect.TypeOf(S1BattleProto2022Energy2019SetArmLight{}),
	10083: reflect.TypeOf(S1BattleProto2022Energy2020SetArmRotate{}),
	10084: reflect.TypeOf(S1BattleProto2022Energy2019ArmorHitUpload{}),
	10085: reflect.TypeOf(S1BattleProto2022Energy2019AmorSelfCheck{}),
	10086: reflect.TypeOf(S1BattleProto2022EnvHeartBeatReq{}),
	10087: reflect.TypeOf(S1BattleProto2022EnvHeartBeatAck{}),
	10088: reflect.TypeOf(S1BattleProto2022IoctrSetRmMotorsCfg{}),
	10089: reflect.TypeOf(S1BattleProto2022IoctrSetRmMotorsCfgAck{}),
	10090: reflect.TypeOf(S1BattleProto2022IoctrSetRmMotorsMoveVal{}),
	10091: reflect.TypeOf(S1BattleProto2022IoctrRmMotorsStatus{}),
	10092: reflect.TypeOf(S1BattleProto2022IoctrSetLightsRgbType1{}),
	10093: reflect.TypeOf(S1BattleProto2022IoctrSetLightsRgbType1Ack{}),
	10094: reflect.TypeOf(S1BattleProto2022IoctrSetLightsRgbType2{}),
	10095: reflect.TypeOf(S1BattleProto2022IoctrSetLightsRgbType2Ack{}),
	10096: reflect.TypeOf(S1BattleProto2022IoctrSetLightsRgbType3{}),
	10097: reflect.TypeOf(S1BattleProto2022IoctrSetLightsRgbType3Ack{}),
	10098: reflect.TypeOf(S1BattleProto2022IoctrSetActuator{}),
	10099: reflect.TypeOf(S1BattleProto2022IoctrSetActuatorAck{}),
	//	10100: reflect.TypeOf(S1BattleProto2022SiloEnvStatus{}),
	//	10101: reflect.TypeOf(S1BattleProto2022SiloEnvCtrlDoor{}),
	10102: reflect.TypeOf(S1BattleProto2022ClientSiloEnvCtrReq{}),
	10103: reflect.TypeOf(S1BattleProto2022SendDartRobotStatus{}),
	10104: reflect.TypeOf(S1BattleProto2022SiloCtrLedData{}),
	10105: reflect.TypeOf(S1BattleProto2022ClientBuyBulletReq{}),
	10106: reflect.TypeOf(S1BattleProto2022ClientBuyBulletAck{}),
	10107: reflect.TypeOf(S1BattleProto2022ClientEconomyNotify{}),
	10108: reflect.TypeOf(S1BattleProto2022ClientAirplaneCtrlReq{}),
	//	10109: reflect.TypeOf(S1BattleProto2022ClientAirplaneStatusInfo{}),
	10110: reflect.TypeOf(S1BattleProto2022ClientShowMessage{}),
	10111: reflect.TypeOf(S1BattleProto2022ClientTalentDataReq{}),
	10112: reflect.TypeOf(S1BattleProto2022ClientTalentDataAck{}),
	10113: reflect.TypeOf(S1BattleProto2022ClientTalentDataNotify{}),
	10114: reflect.TypeOf(S1BattleProto2022SetSupplyStateReq{}),
	10115: reflect.TypeOf(S1BattleProto2022SupplyClearScaleAck{}),
	10116: reflect.TypeOf(S1BattleProto2022SupplyClearScaleReq{}),
	10117: reflect.TypeOf(S1BattleProto2022SupplyExportAck{}),
	10118: reflect.TypeOf(S1BattleProto2022SupplyExportReq{}),
	10119: reflect.TypeOf(S1BattleProto2022SupplyFreedAck{}),
	10120: reflect.TypeOf(S1BattleProto2022SupplyFreedReq{}),
	10121: reflect.TypeOf(S1BattleProto2022SupplyReloadAck{}),
	10122: reflect.TypeOf(S1BattleProto2022SupplyReloadReq{}),
	10123: reflect.TypeOf(S1BattleProto2022SupplyReport{}),
	10124: reflect.TypeOf(S1BattleProto2022SupplyResetReq{}),
	10125: reflect.TypeOf(S1BattleProto2022SupplyRestartReq{}),
	10126: reflect.TypeOf(S1BattleProto2022ClientSupplySync{}),
	10127: reflect.TypeOf(S1BattleProto2022SetFlyCatStatus{}),
	10128: reflect.TypeOf(S1BattleProto2022SetFlyCatStatusAck{}),
	10129: reflect.TypeOf(S1BattleProto2022FlyCatStatus{}),
	10130: reflect.TypeOf(S1BattleProto2022SetFlyCatLight{}),
	10131: reflect.TypeOf(S1BattleProto2022SetFlyCatLightAck{}),
	10132: reflect.TypeOf(S1BattleProto2022SiloDevStatusSyncData{}),
	10133: reflect.TypeOf(S1BattleProto2022ClientWarningNotify{}),
	10134: reflect.TypeOf(S1BattleProto2022ClientShielddata{}),
	10135: reflect.TypeOf(S1BattleProto2022ClientRobotDeathNotify{}),
	10136: reflect.TypeOf(S1BattleProto2022BaseLightingEffect{}),
	10137: reflect.TypeOf(S1BattleProto2022RobotIrcheckReq{}),
	10138: reflect.TypeOf(S1BattleProto2022EnergySetId{}),
	10139: reflect.TypeOf(S1BattleProto2022EnerySetLogoLight{}),
	10140: reflect.TypeOf(S1BattleProto2022EnergyStateChangeNotify{}),
	10141: reflect.TypeOf(S1BattleProto2022EnergyStateSync{}),
	10142: reflect.TypeOf(S1BattleProto2022QueryRobotInfo{}),
	10143: reflect.TypeOf(S1BattleProto2022QueryRobotInfoResult{}),
	10144: reflect.TypeOf(S1BattleProto2022ClientKickOffRobotNotify{}),
	10145: reflect.TypeOf(S1BattleProto2022RobotResurgenceNotify{}),
	10146: reflect.TypeOf(S1BattleProto2022GripperStateNotify{}),
	//	10147: reflect.TypeOf(S1BattleProto2022BaseStatus{}),
	10148: reflect.TypeOf(S1BattleProto2022ClientBaseRobotDevStatusSyncData{}),
	10149: reflect.TypeOf(S1BattleProto2022EnergyReset{}),
	10150: reflect.TypeOf(S1BattleProto2022ExchangeProCtrCmd{}),
	10151: reflect.TypeOf(S1BattleProto2022ExchangeProCtrCmdack{}),
	10152: reflect.TypeOf(S1BattleProto2022ExchangeProLidarInfo{}),
	10153: reflect.TypeOf(S1BattleProto2022ExchangeProPosition{}),
	10154: reflect.TypeOf(S1BattleProto2022ExchangeProRealPosition{}),
	//	10155: reflect.TypeOf(S1BattleProto2022ClientClientStatusSync{}),
	10156: reflect.TypeOf(S1BattleProto2022ClientAllClientStatusSync{}),
	10157: reflect.TypeOf(S1BattleProto2022ClientOutpostDeviceStatusSync{}),
	10158: reflect.TypeOf(S1BattleProto2022ClientMineralInfoSync{}),
	10159: reflect.TypeOf(S1BattleProto2022ClientMineralExchangeNotify{}),
	10160: reflect.TypeOf(S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify{}),
	10161: reflect.TypeOf(S1BattleProto2022ClientFlycatStatusSync{}),
	10162: reflect.TypeOf(S1BattleProto2022StuGameStatus{}),
	10163: reflect.TypeOf(S1BattleProto2022StuGameResult{}),
	10164: reflect.TypeOf(S1BattleProto2022StuRobotCurrentHp{}),
	10165: reflect.TypeOf(S1BattleProto2022StuIcrabuffDebuffZoneStatus{}),
	10166: reflect.TypeOf(S1BattleProto2022StuEnvStatus{}),
	10167: reflect.TypeOf(S1BattleProto2022StuSupplyStatus{}),
	10168: reflect.TypeOf(S1BattleProto2022StuWarning{}),
	10169: reflect.TypeOf(S1BattleProto2022StuMissileRemainingTime{}),
	10170: reflect.TypeOf(S1BattleProto2022StuRobotBuff{}),
	10171: reflect.TypeOf(S1BattleProto2022StuAreialEnergy{}),
	10172: reflect.TypeOf(S1BattleProto2022StuRobotHurt{}),
	10173: reflect.TypeOf(S1BattleProto2022StuLeftBullet{}),
	10174: reflect.TypeOf(S1BattleProto2022StuRfidstatus{}),
	10175: reflect.TypeOf(S1BattleProto2022StuSiloInfo{}),
	10176: reflect.TypeOf(S1BattleProto2022StuCustomControlData{}),
	10177: reflect.TypeOf(S1BattleProto2022StuClientSendCmd{}),
	10178: reflect.TypeOf(S1BattleProto2022StuClientRecvInfo{}),
	10179: reflect.TypeOf(S1BattleProto2022StuCommunicationByteData{}),
	10180: reflect.TypeOf(S1BattleProto2022StuCommunication{}),
	10181: reflect.TypeOf(S1BattleProto2022ClientRaderInfoToClient{}),
	10182: reflect.TypeOf(S1BattleProto2022ClientCustomControlData{}),
	10183: reflect.TypeOf(S1BattleProto2022MapClickInfoNotify{}),
	10184: reflect.TypeOf(S1BattleProto2022RobotNtpServerIpReq{}),
	10185: reflect.TypeOf(S1BattleProto2022RobotNtpServerIpAck{}),
	10186: reflect.TypeOf(S1BattleProto2022RobotMoudleLife{}),
	10187: reflect.TypeOf(S1BattleProto2022ClientCheckInTimeStampNotify{}),
	//	10188: reflect.TypeOf(S1BattleProto2022RobotCheckinTimeStampReq{}),
	//	10189: reflect.TypeOf(S1BattleProto2022RobotCheckinTimeStampAck{}),
	10190: reflect.TypeOf(S1BattleProto2022RobotMeasureStartReq{}),
	10191: reflect.TypeOf(S1BattleProto2022RobotMeasureStartRsp{}),
	10192: reflect.TypeOf(S1BattleProto2022RobotMeasureStopReq{}),
	10193: reflect.TypeOf(S1BattleProto2022RobotMeasureStopRsp{}),
	10194: reflect.TypeOf(S1BattleProto2022RobotPushCapinfo{}),
	10195: reflect.TypeOf(S1BattleProto2022RobotPushCaprtinfo{}),
	10196: reflect.TypeOf(S1BattleProto2022RobotPushSensorInfo{}),
	10197: reflect.TypeOf(S1BattleProto2022RobotGetCheckCapReq{}),
	10198: reflect.TypeOf(S1BattleProto2022RobotGetCheckCapAck{}),
	10199: reflect.TypeOf(S1BattleProto2022ClientExceptionCapDataNotify{}),
	10200: reflect.TypeOf(S1BattleProto2022ClientExceptionRecoverCardNotify{}),
	10201: reflect.TypeOf(S1BattleProto2022ClientRecoverCardStatus{}),
	10202: reflect.TypeOf(S1BattleProto2022ClientVtmSpeedSync{}),
	10203: reflect.TypeOf(S1BattleProto2022ClientServerMapSync{}),
	10204: reflect.TypeOf(S1BattleProto2022ClientServerUimessage{}),
	10205: reflect.TypeOf(S1BattleProto2022RobotBaseDevCtlCmdAck{}),
	10206: reflect.TypeOf(S1BattleProto2022ClientWeaponFireNotify{}),
	//	10207: reflect.TypeOf(S1BattleProto2022RobotVtmsetting{}),
	10208: reflect.TypeOf(S1BattleProto2022ClientExchangeProStateNotify{}),
	10209: reflect.TypeOf(S1BattleProto2022RobotNewHeartBeatReq{}),
	10210: reflect.TypeOf(S1BattleProto2022RobotNewHeartBeatAck{}),
	10211: reflect.TypeOf(S1BattleProto2022IoctrCfgReq{}),
	10212: reflect.TypeOf(S1BattleProto2022ExchangeProClearOreRes{}),
	10213: reflect.TypeOf(S1BattleProto2022ClientHoldCenterPointCompleteNotify{}),
	10214: reflect.TypeOf(S1BattleProtoBaseArmorStateNotify{}),
	10215: reflect.TypeOf(S1BattleProto2023ClientServerRobotsBattleInfoSync{}),
	10216: reflect.TypeOf(S1BattleProto2023ClientTeamSupplyInfoSync{}),
	//	10217: reflect.TypeOf(S1BattleProto2022EnergyArmorLife{}),
	//	10218: reflect.TypeOf(S1BattleProto2022EnergyArmorLifeAck{}),
	//	10219: reflect.TypeOf(S1BattleProto2023OutPostArmorResetReq{}),
	//	10220: reflect.TypeOf(S1BattleProto2023OutPostArmorResetAck{}),
	//	10221: reflect.TypeOf(S1BattleProto2023OutPostArmorResetFinish{}),
	//	10222: reflect.TypeOf(S1BattleProto2023ClientBuyReviveReq{}),
	//	10223: reflect.TypeOf(S1BattleProto2023ClientBuyReviveAck{}),
	//	10224: reflect.TypeOf(S1BattleProto2023ClientGameSysUploadResultNotify{}),
	//	10225: reflect.TypeOf(S1BattleProto2023ClientTeamControlZoneInfoSync{}),
	10226: reflect.TypeOf(S1BattleProto2023ClientGuardInPatrolInfoSync{}),
	//	10227: reflect.TypeOf(S1BattleProto2023ClientRobotPathPlanInfo{}),
	//	10228: reflect.TypeOf(S1BattleProto2023ClientSingleLadderRewardNotify{}),
	//	10229: reflect.TypeOf(S1BattleProto2022StuTeamRobotsPos{}),
	//	10230: reflect.TypeOf(S1BattleProto2023StuRadarMarkProgress{}),
	//	10231: reflect.TypeOf(S1BattleProto2022ClientHitRuneInfoSync{}),
	//	10232: reflect.TypeOf(S1BattleProto2022ClientHitRuneNotifySync{}),
	//	10233: reflect.TypeOf(S1BattleProto2022ExchangeProSetArmRunSpeed{}),
	//	10234: reflect.TypeOf(S1BattleProto2022ExchangeProSetArmRunSpeedAck{}),
	//	10235: reflect.TypeOf(S1BattleProto2023ExchangeProErrorInfo{}),
	10236: reflect.TypeOf(S1BattleProto2023ExchangeProErrorInfoNotify{}),
	//	10237: reflect.TypeOf(S1BattleProto2023CustomControlDataLenInfo{}),
	//	10238: reflect.TypeOf(S1BattleProto2023BigRuneHitRecordNotify{}),
	10239: reflect.TypeOf(S1BattleProto2023ClientPenaltyTableInfos{}),
	10240: reflect.TypeOf(S1BattleProto2023ClientPenaltyTableInfosAck{}),
	//	10241: reflect.TypeOf(S1BattleProto2023ClientBeAttackedDamageNotify{}),
	10242: reflect.TypeOf(S1BattleProto2022ClientSeverAlertNotify{}),
	//	10243: reflect.TypeOf(S1BattleProto2023EnergyAngleReport{}),
	//	10244: reflect.TypeOf(S1BattleProto2023EnergyTestModeSet{}),
	10245: reflect.TypeOf(S1BattleProto2022ClientRobotArmorLifeQueryResult{}),
	10246: reflect.TypeOf(S1BattleProto2022ClientArmorLifeInfo{}),
	//	10247: reflect.TypeOf(S1BattleProto2023RobotArmorLife{}),
	//	10248: reflect.TypeOf(S1BattleProto2023IrarmorLightCtrl{}),
	//	10249: reflect.TypeOf(S1BattleProto2023IrarmorSelfCheck{}),
	//	10250: reflect.TypeOf(S1BattleProto2023IrarmorSelfCheckAck{}),
	//	10251: reflect.TypeOf(S1BattleProto2022BaseDartTargetControl{}),
	//	10252: reflect.TypeOf(S1BattleProto2022BaseDartTargetControlAck{}),
	//	10253: reflect.TypeOf(S1BattleProto2023ClientRobotCustomMessage{}),
	//	10254: reflect.TypeOf(S1BattleProto2023ExchangeProMineralValueSync{}),
	//	10255: reflect.TypeOf(S1BattleProto2022ClientMapClickTransResult{}),
	//	10256: reflect.TypeOf(S1BattleProto2022StuGuardDecision{}),
	//	10257: reflect.TypeOf(S1BattleProto2022StuGuardDecisionInfo{}),
	//	10258: reflect.TypeOf(S1BattleProto2022StuRadarDecision{}),
	//	10259: reflect.TypeOf(S1BattleProto2022StuRadarDecisionInfo{}),
	//	10260: reflect.TypeOf(S1BattleProto2022ClientRobotExpAddData{}),
	//	10261: reflect.TypeOf(S1BattleProto2023OutPostStatus{}),
	10262: reflect.TypeOf(S1BattleProto2023ClientModuleVersionState{}),
	//	10263: reflect.TypeOf(S1BattleProto2023ClientInitRuleParms{}),
	//	10264: reflect.TypeOf(S1BattleProto2022StuRadarMarkPositionInfo{}),
	10265: reflect.TypeOf(S1BattleProto2023ClientGameSystemInfoNotify{}),
	10266: reflect.TypeOf(S1BattleProto2023ClientChangeTokenCmd{}),
}
