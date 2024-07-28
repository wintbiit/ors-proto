package proto

type S1BattleProto2023ClientRobotBattleInfo struct {
	IsOnHpPoint                  byte
	GimbalPowerCtrl              byte
	Robots0Id                    byte
	BattleState                  byte
	CanRemoteExchangeSmallBullet byte
	RemaindBuyRevivalCount       byte
	BuyRevivalPrice              int32
	BattleStateCountDown         float32
	CanRemoteHeal                byte
	CanRemoteExchangeBigBullet   byte
	ChassisPowerCtrl             byte
	ShooterPowerCtrl             byte
}

const S1BattleProto2023ClientRobotBattleInfoSize = 18

func (s *S1BattleProto2023ClientRobotBattleInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientRobotBattleInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientRobotBattleInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
