package proto

type S1BattleProto2022ClientBaseRobotDevStatus struct {
	BaseShellMotoOneOverHeat       byte
	BaseDartBoardMotoOverCurrent   byte
	BaseShellOneSensorStatus       byte
	BaseDartboardSelfcheckStatus   byte
	DartBoardBrokenErr             byte
	BaseDartBoardMotoOverSpeed     byte
	BaseShellMotoOneOnlineStatus   byte
	BaseShellMotoTwoOnlineStatus   byte
	DartTargetIsInplace            byte
	BaseShellTwoSensorStatus       byte
	BaseShellThreeSensorStatus     byte
	BaseShellCloseTimeout          byte
	BaseShellMotoThreeOverHeat     byte
	ShellStatus                    byte
	Online                         byte
	BaseShellOpenTimeout           byte
	IrledStatusRight               uint32
	BaseShellMotoTwoOverCurrent    byte
	BaseDartboardSensorStatus      byte
	BaseShellMotoThreeBlockStatus  byte
	BaseDartboardMoveTimeout       byte
	BaseShellMotoTwoOverHeat       byte
	BaseShellMotoThreeOverCurrent  byte
	BaseShellMotoOneOverSpeed      byte
	BaseShellMotoThreeOverSpeed    byte
	BaseDartboardMotoBlockStatus   byte
	BaseDartboardMotoOnlineStatus  byte
	BaseShellMotoTwoOverSpeed      byte
	BaseShellMotoThreeOnlineStatus byte
	BaseShellSelfcheckStatus       byte
	DartBoardIronline              byte
	IrledStatusLeft                uint32
	BaseDartBoardMotoOverHeat      byte
	BaseShellMotoOneOverCurrent    byte
	BaseShellMotoOneBlockStatus    byte
	BaseShellMotoTwoBlockStatus    byte
	DartTargetPosition             int16
}

const S1BattleProto2022ClientBaseRobotDevStatusSize = 44

func (s *S1BattleProto2022ClientBaseRobotDevStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBaseRobotDevStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientBaseRobotDevStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
