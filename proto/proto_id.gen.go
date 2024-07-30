package proto

const (
	ProtoIDS1ProtoLoginReq                                      = 9000
	ProtoIDS1ProtoLoginAck                                      = 9001
	ProtoIDS1ProtoReLoginReq                                    = 9002
	ProtoIDS1ProtoReLoginAck                                    = 9003
	ProtoIDS1ProtoReLoginRoomNotify                             = 9004
	ProtoIDS1ProtoLogoutReq                                     = 9005
	ProtoIDS1ProtoLogoutAck                                     = 9006
	ProtoIDS1ProtoLogoutNotify                                  = 9007
	ProtoIDS1ProtoHeartBeatReq                                  = 9008
	ProtoIDS1ProtoHeartBeatAck                                  = 9009
	ProtoIDS1ProtoTestReq                                       = 9010
	ProtoIDS1ProtoTestAck                                       = 9011
	ProtoIDS1ProtoPingReq                                       = 9012
	ProtoIDS1ProtoPingAck                                       = 9013
	ProtoIDS1ProtoEnterRoomReq                                  = 9014
	ProtoIDS1ProtoEnterRoomAck                                  = 9015
	ProtoIDS1ProtoEnterRoomNotify                               = 9016
	ProtoIDS1ProtoRoomInfosNotify                               = 9017
	ProtoIDS1ProtoRoomStateNotify                               = 9018
	ProtoIDS1ProtoRoomPauseNotify                               = 9019
	ProtoIDS1ProtoLeaveRoomReq                                  = 9020
	ProtoIDS1ProtoLeaveRoomAck                                  = 9021
	ProtoIDS1ProtoLeaveRoomNotify                               = 9022
	ProtoIDS1ProtoSetTeamReq                                    = 9023
	ProtoIDS1ProtoSetTeamAck                                    = 9024
	ProtoIDS1ProtoSetTeamNotify                                 = 9025
	ProtoIDS1ProtoSetTidReq                                     = 9026
	ProtoIDS1ProtoSetTidAck                                     = 9027
	ProtoIDS1ProtoSetTidNotify                                  = 9028
	ProtoIDS1ProtoSetReadyReq                                   = 9029
	ProtoIDS1ProtoSetReadyAck                                   = 9030
	ProtoIDS1ProtoSetReadyNotify                                = 9031
	ProtoIDS1ProtoSetTokenReq                                   = 9032
	ProtoIDS1ProtoSetTokenAck                                   = 9033
	ProtoIDS1ProtoSetSvrInfoReq                                 = 9034
	ProtoIDS1ProtoSetSvrInfoAck                                 = 9035
	ProtoIDS1ProtoSetWifiInfoNotify                             = 9036
	ProtoIDS1ProtoSetLockSeatNotify                             = 9037
	ProtoIDS1ProtoSetCurrMatchNotify                            = 9038
	ProtoIDS1ProtoSetTeamPlayerNumNotify                        = 9039
	ProtoIDS1ProtoGMCommandReq                                  = 9040
	ProtoIDS1ProtoGMCommandAck                                  = 9041
	ProtoIDS1ProtoPlayerCommand_UserInput_Notify                = 9042
	ProtoIDS1ProtoPlayerCommand_Disconnect_Notify               = 9043
	ProtoIDS1ProtoClientSyncReq                                 = 9044
	ProtoIDS1ProtoClientSyncNotify                              = 9045
	ProtoIDS1ProtoTeamInfoNotify                                = 9046
	ProtoIDS1ProtoTeamLogoNotify                                = 9047
	ProtoIDS1ProtoTeamCampNotify                                = 9048
	ProtoIDS1ProtoSvrStateReq                                   = 9049
	ProtoIDS1ProtoSvrStateAck                                   = 9050
	ProtoIDS1ProtoUserStateReq                                  = 9051
	ProtoIDS1ProtoUserStateAck                                  = 9052
	ProtoIDS1ClientTransferRobotProto                           = 9053
	ProtoIDS1ProtoRobotNetworkStatusReq                         = 9054
	ProtoIDS1ProtoRobotNetworkStatusAck                         = 9055
	ProtoIDS1ProtoClientNetworkInfoNotify                       = 9056
	ProtoIDS1ProtoRobotStudentSerialPortInfos                   = 9057
	ProtoIDS1ProtoNetworkInfoNotify                             = 9058
	ProtoIDS1ProtoTechnicalPauseInfoNotify                      = 9059
	ProtoIDS1ProtoMarkDetectResult                              = 9060
	ProtoIDS1ProtoMarkDetectResultAck                           = 9061
	ProtoIDS1ProtoClientExceptionInfo                           = 9062
	ProtoIDS1BattleProtoAutoStateNotify                         = 10000
	ProtoIDS1BattleProtoArmorHitReq                             = 10001
	ProtoIDS1BattleProtoArmorHitAck                             = 10002
	ProtoIDS1BattleProtoBeAttackNotify                          = 10003
	ProtoIDS1BattleProtoGunFireReq                              = 10004
	ProtoIDS1BattleProtoGunFireAck                              = 10005
	ProtoIDS1BattleProtoGunFireNotify                           = 10006
	ProtoIDS1BattleProtoGunHeatNotify                           = 10007
	ProtoIDS1BattleProtoGunBulletNotify                         = 10008
	ProtoIDS1BattleProtoAttrsNotify                             = 10009
	ProtoIDS1BattleProtoPlayerInitOkNotify                      = 10010
	ProtoIDS1BattleProtoPlayerDeadNotify                        = 10011
	ProtoIDS1BattleProtoPlayerReliveNotify                      = 10012
	ProtoIDS1BattleProtoBuffAddNotify                           = 10013
	ProtoIDS1BattleProtoBuffDelNotify                           = 10014
	ProtoIDS1BattleProtoBuffModifyNotify                        = 10015
	ProtoIDS1BattleProtoMarkerDetectReq                         = 10016
	ProtoIDS1BattleProtoMarkerDetectAck                         = 10017
	ProtoIDS1BattleProtoProgressNotify                          = 10018
	ProtoIDS1BattleProtoDeviceModuleReq                         = 10019
	ProtoIDS1BattleProtoDeviceModuleNotify                      = 10020
	ProtoIDS1BattleProtoPoJiaNotify                             = 10021
	ProtoIDS1BattleProtoPlayerCommand_GunFire_Notify            = 10022
	ProtoIDS1BattleProtoPlayerCommand_Move_Notify               = 10023
	ProtoIDS1BattleProtoPlayerCommand_LockScreen_Notify         = 10024
	ProtoIDS1BattleProtoPlayerCommand_AutoControl_Notify        = 10025
	ProtoIDS1BattleProtoPlayerCommand_UserInput_Notify          = 10026
	ProtoIDS1BattleProtoPlayerCommand_UserCustomAction_Notify   = 10027
	ProtoIDS1BattleProtoPlayerCommand_SetDisconnectNotify       = 10028
	ProtoIDS1BattleProtoSceneInfosNotify                        = 10029
	ProtoIDS1BattleProtoSceneStateNotify                        = 10030
	ProtoIDS1BattleProtoGameSettlementNotify                    = 10031
	ProtoIDS1BattleProtoReloginNotify                           = 10032
	ProtoIDS1BattleProtoIsCanUseATKNotify                       = 10033
	ProtoIDS1BattleProtoIsCanUsePJNotify                        = 10034
	ProtoIDS1BattleProtoClientSyncReq                           = 10035
	ProtoIDS1BattleProtoClientSyncNotify                        = 10036
	ProtoIDS1BattleProtoDeviceBulletNotify                      = 10037
	ProtoIDS1BattleProtoWarningNotify                           = 10038
	ProtoIDS1BattleProtoPlayerResetNotify                       = 10039
	ProtoIDS1BattleProtoReduceReliveTimeNotify                  = 10040
	ProtoIDS1BattleProtoFullSceneDataReq                        = 10041
	ProtoIDS1BattleProtoTeamInfoNotify                          = 10042
	ProtoIDS1BattleProtoPlaceholderNotify                       = 10043
	ProtoIDS1BattleProtoModuleOffLineInfoNotify                 = 10044
	ProtoIDS1BattleProtoAddBulletNotify                         = 10045
	ProtoIDS1BattleProtoShareBulletReq                          = 10046
	ProtoIDS1BattleProtoUploadModuleInfo                        = 10047
	ProtoIDS1BattleProto2022RobotInitCfgTab                     = 10048
	ProtoIDS1BattleProto2022ConfigTabAck                        = 10049
	ProtoIDS1BattleProto2022RobotStatus                         = 10050
	ProtoIDS1BattleProto2022RobotsDataSync                      = 10051
	ProtoIDS1BattleProto2022RobotsWeaponFire                    = 10052
	ProtoIDS1BattleProto2022RobotHited                          = 10053
	ProtoIDS1BattleProto2022VtmSpeedSync                        = 10054
	ProtoIDS1BattleProto2022ClientsServerBaseSync               = 10055
	ProtoIDS1BattleProto2022ClientsServerSync                   = 10056
	ProtoIDS1BattleProto2022ClientsFirstSync                    = 10057
	ProtoIDS1BattleProto2022ClientsRobotStatusNotify            = 10058
	ProtoIDS1BattleProto2022ClientsRobotModuleErrorInfo         = 10059
	ProtoIDS1BattleProto2022TestTimeDelayDownData               = 10060
	ProtoIDS1BattleProto2022TestTimeDelayDownComplete           = 10061
	ProtoIDS1BattleProto2022TestDelayDownDataAck                = 10062
	ProtoIDS1BattleProto2022CONTROL_CMD_T                       = 10063
	ProtoIDS1BattleProto2022ClientGMCommand                     = 10064
	ProtoIDS1BattleProto2022ClientCurrentProgress               = 10065
	ProtoIDS1BattleProto2022ClientHitNotify                     = 10066
	ProtoIDS1BattleProto2022ReLoginDataSync                     = 10067
	ProtoIDS1BattleProto2022ClientResultDataInfo                = 10068
	ProtoIDS1BattleProto2022EnvBaseShellControl                 = 10069
	ProtoIDS1BattleProto2022RFIDReq                             = 10070
	ProtoIDS1BattleProto2022RFIDNotify                          = 10071
	ProtoIDS1BattleProto2022IOCtrCfgSet                         = 10072
	ProtoIDS1BattleProto2022IOCtrCfgSetAck                      = 10073
	ProtoIDS1BattleProto2022IOCtrSetVal                         = 10074
	ProtoIDS1BattleProto2022IOCtrSetValAck                      = 10075
	ProtoIDS1BattleProto2022IOCtrState                          = 10076
	ProtoIDS1BattleProto2022IOCtrTriggerVal                     = 10077
	ProtoIDS1BattleProto2022EnvDevicesDescripeReq               = 10078
	ProtoIDS1BattleProto2022EnvDevicesDescripeAck               = 10079
	ProtoIDS1BattleProto2022S2C_PowerStateNotify                = 10080
	ProtoIDS1BattleProto2022Energy2019StateReport               = 10081
	ProtoIDS1BattleProto2022Energy2019SetArmLight               = 10082
	ProtoIDS1BattleProto2022Energy2020SetArmRotate              = 10083
	ProtoIDS1BattleProto2022Energy2019ArmorHitUpload            = 10084
	ProtoIDS1BattleProto2022Energy2019AmorSelfCHeck             = 10085
	ProtoIDS1BattleProto2022Env_Heart_Beat_Req                  = 10086
	ProtoIDS1BattleProto2022Env_Heart_Beat_Ack                  = 10087
	ProtoIDS1BattleProto2022_IOCtrSetRmMotorsCfg                = 10088
	ProtoIDS1BattleProto2022_IOCtrSetRmMotorsCfgAck             = 10089
	ProtoIDS1BattleProto2022_IOCtrSetRmMotorsMoveVal            = 10090
	ProtoIDS1BattleProto2022_IOCtrRmMotorsStatus                = 10091
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType1             = 10092
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType1Ack          = 10093
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType2             = 10094
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType2Ack          = 10095
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType3             = 10096
	ProtoIDS1BattleProto2022_IOCtrSetLightsRgbType3Ack          = 10097
	ProtoIDS1BattleProto2022_IOCtrSetActuator                   = 10098
	ProtoIDS1BattleProto2022_IOCtrSetActuatorAck                = 10099
	ProtoIDS1BattleProto2022_SiloEnvStatus                      = 10100
	ProtoIDS1BattleProto2022_SiloEnvCtrlDoor                    = 10101
	ProtoIDS1BattleProto2022ClientSiloEnvCtrReq                 = 10102
	ProtoIDS1BattleProto2022SendDartRobotStatus                 = 10103
	ProtoIDS1BattleProto2022SiloCtrLedData                      = 10104
	ProtoIDS1BattleProto2022ClientBuyBulletReq                  = 10105
	ProtoIDS1BattleProto2022ClientBuyBulletAck                  = 10106
	ProtoIDS1BattleProto2022ClientEconomyNotify                 = 10107
	ProtoIDS1BattleProto2022ClientAirplaneCtrlReq               = 10108
	ProtoIDS1BattleProto2022ClientAirplaneStatusInfo            = 10109
	ProtoIDS1BattleProto2022ClientShowMessage                   = 10110
	ProtoIDS1BattleProto2022ClientTalentDataReq                 = 10111
	ProtoIDS1BattleProto2022ClientTalentDataAck                 = 10112
	ProtoIDS1BattleProto2022ClientTalentDataNotify              = 10113
	ProtoIDS1BattleProto2022SetSupplyStateReq                   = 10114
	ProtoIDS1BattleProto2022SupplyClearScaleAck                 = 10115
	ProtoIDS1BattleProto2022SupplyClearScaleReq                 = 10116
	ProtoIDS1BattleProto2022SupplyExportAck                     = 10117
	ProtoIDS1BattleProto2022SupplyExportReq                     = 10118
	ProtoIDS1BattleProto2022SupplyFreedAck                      = 10119
	ProtoIDS1BattleProto2022SupplyFreedReq                      = 10120
	ProtoIDS1BattleProto2022SupplyReloadAck                     = 10121
	ProtoIDS1BattleProto2022SupplyReloadReq                     = 10122
	ProtoIDS1BattleProto2022SupplyReport                        = 10123
	ProtoIDS1BattleProto2022SupplyResetReq                      = 10124
	ProtoIDS1BattleProto2022SupplyRestartReq                    = 10125
	ProtoIDS1BattleProto2022ClientSupplySync                    = 10126
	ProtoIDS1BattleProto2022SetFlyCatStatus                     = 10127
	ProtoIDS1BattleProto2022SetFlyCatStatusAck                  = 10128
	ProtoIDS1BattleProto2022FlyCatStatus                        = 10129
	ProtoIDS1BattleProto2022SetFlyCatLight                      = 10130
	ProtoIDS1BattleProto2022SetFlyCatLightAck                   = 10131
	ProtoIDS1BattleProto2022SiloDevStatusSyncData               = 10132
	ProtoIDS1BattleProto2022ClientWarningNotify                 = 10133
	ProtoIDS1BattleProto2022ClientSHIELDDATA                    = 10134
	ProtoIDS1BattleProto2022ClientRobotDeathNotify              = 10135
	ProtoIDS1BattleProto2022BaseLightingEffect                  = 10136
	ProtoIDS1BattleProto2022RobotIRCheckReq                     = 10137
	ProtoIDS1BattleProto2022EnergySetID                         = 10138
	ProtoIDS1BattleProto2022EnerySetLogoLight                   = 10139
	ProtoIDS1BattleProto2022EnergyStateChangeNotify             = 10140
	ProtoIDS1BattleProto2022EnergyStateSync                     = 10141
	ProtoIDS1BattleProto2022QueryRobotInfo                      = 10142
	ProtoIDS1BattleProto2022QueryRobotInfoResult                = 10143
	ProtoIDS1BattleProto2022ClientKickOffRobotNotify            = 10144
	ProtoIDS1BattleProto2022RobotResurgenceNotify               = 10145
	ProtoIDS1BattleProto2022GripperStateNotify                  = 10146
	ProtoIDS1BattleProto2022BaseStatus                          = 10147
	ProtoIDS1BattleProto2022ClientBaseRobotDevStatusSyncData    = 10148
	ProtoIDS1BattleProto2022EnergyReset                         = 10149
	ProtoIDS1BattleProto2022ExchangeProCtrCMD                   = 10150
	ProtoIDS1BattleProto2022ExchangeProCtrCMDAck                = 10151
	ProtoIDS1BattleProto2022ExchangeProLidarInfo                = 10152
	ProtoIDS1BattleProto2022ExchangeProPosition                 = 10153
	ProtoIDS1BattleProto2022ExchangeProRealPosition             = 10154
	ProtoIDS1BattleProto2022ClientClientStatusSync              = 10155
	ProtoIDS1BattleProto2022ClientAllClientStatusSync           = 10156
	ProtoIDS1BattleProto2022ClientOutpostDeviceStatusSync       = 10157
	ProtoIDS1BattleProto2022ClientMineralInfoSync               = 10158
	ProtoIDS1BattleProto2022ClientMineralExchangeNotify         = 10159
	ProtoIDS1BattleProto2022ClientSiloEnvDoorOpenCloseNotify    = 10160
	ProtoIDS1BattleProto2022ClientFlycatStatusSync              = 10161
	ProtoIDS1BattleProto2022StuGameStatus                       = 10162
	ProtoIDS1BattleProto2022StuGameResult                       = 10163
	ProtoIDS1BattleProto2022StuRobotCurrentHP                   = 10164
	ProtoIDS1BattleProto2022StuICRABuffDebuffZoneStatus         = 10165
	ProtoIDS1BattleProto2022StuEnvStatus                        = 10166
	ProtoIDS1BattleProto2022StuSupplyStatus                     = 10167
	ProtoIDS1BattleProto2022StuWarning                          = 10168
	ProtoIDS1BattleProto2022StuMissileRemainingTime             = 10169
	ProtoIDS1BattleProto2022StuRobotBuff                        = 10170
	ProtoIDS1BattleProto2022StuAreialEnergy                     = 10171
	ProtoIDS1BattleProto2022StuRobotHurt                        = 10172
	ProtoIDS1BattleProto2022StuLeftBullet                       = 10173
	ProtoIDS1BattleProto2022StuRFIDStatus                       = 10174
	ProtoIDS1BattleProto2022StuSiloInfo                         = 10175
	ProtoIDS1BattleProto2022StuCustomControlData                = 10176
	ProtoIDS1BattleProto2022StuClientSendCMD                    = 10177
	ProtoIDS1BattleProto2022StuClientRecvInfo                   = 10178
	ProtoIDS1BattleProto2022StuCommunicationByteData            = 10179
	ProtoIDS1BattleProto2022StuCommunication                    = 10180
	ProtoIDS1BattleProto2022ClientRaderInfoToClient             = 10181
	ProtoIDS1BattleProto2022ClientCustomControlData             = 10182
	ProtoIDS1BattleProto2022MapClickInfoNotify                  = 10183
	ProtoIDS1BattleProto2022RobotNtpServerIpReq                 = 10184
	ProtoIDS1BattleProto2022RobotNtpServerIpAck                 = 10185
	ProtoIDS1BattleProto2022RobotMoudleLife                     = 10186
	ProtoIDS1BattleProto2022ClientCheckInTimeStampNotify        = 10187
	ProtoIDS1BattleProto2022RobotCheckinTimeStampReq            = 10188
	ProtoIDS1BattleProto2022RobotCheckinTimeStampAck            = 10189
	ProtoIDS1BattleProto2022RobotMeasureStartReq                = 10190
	ProtoIDS1BattleProto2022RobotMeasureStartRsp                = 10191
	ProtoIDS1BattleProto2022RobotMeasureStopReq                 = 10192
	ProtoIDS1BattleProto2022RobotMeasureStopRsp                 = 10193
	ProtoIDS1BattleProto2022RobotPushCAPInfo                    = 10194
	ProtoIDS1BattleProto2022RobotPushCAPRTInfo                  = 10195
	ProtoIDS1BattleProto2022RobotPushSensorInfo                 = 10196
	ProtoIDS1BattleProto2022RobotGetCheckCapReq                 = 10197
	ProtoIDS1BattleProto2022RobotGetCheckCapAck                 = 10198
	ProtoIDS1BattleProto2022ClientExceptionCapDataNotify        = 10199
	ProtoIDS1BattleProto2022ClientExceptionRecoverCardNotify    = 10200
	ProtoIDS1BattleProto2022ClientRecoverCardStatus             = 10201
	ProtoIDS1BattleProto2022ClientVtmSpeedSync                  = 10202
	ProtoIDS1BattleProto2022ClientServerMapSync                 = 10203
	ProtoIDS1BattleProto2022ClientServerUIMessage               = 10204
	ProtoIDS1BattleProto2022RobotBaseDevCtlCmdAck               = 10205
	ProtoIDS1BattleProto2022ClientWeaponFireNotify              = 10206
	ProtoIDS1BattleProto2022RobotVTMSetting                     = 10207
	ProtoIDS1BattleProto2022ClientExchangeProStateNotify        = 10208
	ProtoIDS1BattleProto2022RobotNewHeartBeatReq                = 10209
	ProtoIDS1BattleProto2022RobotNewHeartBeatAck                = 10210
	ProtoIDS1BattleProto2022_IOCtrCfgReq                        = 10211
	ProtoIDS1BattleProto2022ExchangeProClearOreRes              = 10212
	ProtoIDS1BattleProto2022ClientHoldCenterPointCompleteNotify = 10213
	ProtoIDS1BattleProto_Base_Armor_State_Notify                = 10214
	ProtoIDS1BattleProto2023ClientServerRobotsBattleInfoSync    = 10215
	ProtoIDS1BattleProto2023ClientTeamSupplyInfoSync            = 10216
	ProtoIDS1BattleProto2022EnergyArmorLife                     = 10217
	ProtoIDS1BattleProto2022EnergyArmorLifeAck                  = 10218
	ProtoIDS1BattleProto2023_OutPostArmorResetReq               = 10219
	ProtoIDS1BattleProto2023_OutPostArmorResetAck               = 10220
	ProtoIDS1BattleProto2023_OutPostArmorResetFinish            = 10221
	ProtoIDS1BattleProto2023ClientBuyReviveReq                  = 10222
	ProtoIDS1BattleProto2023ClientBuyReviveAck                  = 10223
	ProtoIDS1BattleProto2023ClientGameSysUploadResultNotify     = 10224
	ProtoIDS1BattleProto2023ClientTeamControlZoneInfoSync       = 10225
	ProtoIDS1BattleProto2023ClientGuardInPatrolInfoSync         = 10226
	ProtoIDS1BattleProto2023ClientRobotPathPlanInfo             = 10227
	ProtoIDS1BattleProto2023ClientSingleLadderRewardNotify      = 10228
	ProtoIDS1BattleProto2022StuTeamRobotsPos                    = 10229
	ProtoIDS1BattleProto2023StuRadarMarkProgress                = 10230
	ProtoIDS1BattleProto2022ClientHitRuneInfoSync               = 10231
	ProtoIDS1BattleProto2022ClientHitRuneNotifySync             = 10232
	ProtoIDS1BattleProto2022ExchangeProSetArmRunSpeed           = 10233
	ProtoIDS1BattleProto2022ExchangeProSetArmRunSpeedAck        = 10234
	ProtoIDS1BattleProto2023ExchangeProErrorInfo                = 10235
	ProtoIDS1BattleProto2023ExchangeProErrorInfoNotify          = 10236
	ProtoIDS1BattleProto2023CustomControlDataLenInfo            = 10237
	ProtoIDS1BattleProto2023BigRuneHitRecordNotify              = 10238
	ProtoIDS1BattleProto2023ClientPenaltyTableInfos             = 10239
	ProtoIDS1BattleProto2023ClientPenaltyTableInfosAck          = 10240
	ProtoIDS1BattleProto2023ClientBeAttackedDamageNotify        = 10241
	ProtoIDS1BattleProto2022ClientSeverAlertNotify              = 10242
	ProtoIDS1BattleProto2023EnergyAngleReport                   = 10243
	ProtoIDS1BattleProto2023EnergyTestModeSet                   = 10244
	ProtoIDS1BattleProto2022ClientRobotArmorLifeQueryResult     = 10245
	ProtoIDS1BattleProto2022ClientArmorLifeInfo                 = 10246
	ProtoIDS1BattleProto2023RobotArmorLife                      = 10247
	ProtoIDS1BattleProto2023IRArmorLightCtrl                    = 10248
	ProtoIDS1BattleProto2023IRArmorSelfCheck                    = 10249
	ProtoIDS1BattleProto2023IRArmorSelfCheckAck                 = 10250
	ProtoIDS1BattleProto2022BaseDartTargetControl               = 10251
	ProtoIDS1BattleProto2022BaseDartTargetControlAck            = 10252
	ProtoIDS1BattleProto2023ClientRobotCustomMessage            = 10253
	ProtoIDS1BattleProto2023ExchangeProMineralValueSync         = 10254
	ProtoIDS1BattleProto2022ClientMapClickTransResult           = 10255
	ProtoIDS1BattleProto2022StuGuardDecision                    = 10256
	ProtoIDS1BattleProto2022StuGuardDecisionInfo                = 10257
	ProtoIDS1BattleProto2022StuRadarDecision                    = 10258
	ProtoIDS1BattleProto2022StuRadarDecisionInfo                = 10259
	ProtoIDS1BattleProto2022ClientRobotExpAddData               = 10260
	ProtoIDS1BattleProto2023_OutPostStatus                      = 10261
	ProtoIDS1BattleProto2023ClientModuleVersionState            = 10262
	ProtoIDS1BattleProto2023ClientInitRuleParms                 = 10263
	ProtoIDS1BattleProto2022StuRadarMarkPositionInfo            = 10264
	ProtoIDS1BattleProto2023ClientGameSystemInfoNotify          = 10265
	ProtoIDS1BattleProto2023ClientChangeTokenCMD                = 10266
)

var ProtoIdMap = map[int]string{
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
