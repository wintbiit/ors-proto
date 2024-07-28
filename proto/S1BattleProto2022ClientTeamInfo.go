package proto

type S1BattleProto2022ClientTeamInfo struct {
	TeamAvaliableCoinsNum   int32
	BigBulletQuotaNum       int16
	BigBulletCanBuyNum      int16
	Warning2Count           uint16
	Warning3Count           uint16
	TotalHurt               uint16
	BigBulletAvailableNum   int16
	BloodPack               uint32
	TotalHp                 uint32
	Warning1Count           uint16
	CenterPointEnergy       uint32
	TotalCoins              int32
	CurrentHp               uint32
	AirSupportNum           uint16
	SmallBulletCanBuyNum    int16
	SmallBulletAvailableNum int16
	SmallBulletQuotaNum     int16
}

const S1BattleProto2022ClientTeamInfoSize = 46

func (s *S1BattleProto2022ClientTeamInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientTeamInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientTeamInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
