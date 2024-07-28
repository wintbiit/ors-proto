package proto

type S1BattleProto2022ClientHpChangeDetailVal struct {
	ByGolf          int32
	ByOverSpeed     int32
	ByoverFreq      int32
	ByOverPower     int32
	ByOverHeat      int32
	ByModuleOffline int32
	ByPunish        int32
	ByCrash         int32
	ByBullet        int32
	ByMissile       int32
	ByKill          int32
	ByWifiOffline   int32
	All             int32
}

const S1BattleProto2022ClientHpChangeDetailValSize = 52

func (s *S1BattleProto2022ClientHpChangeDetailVal) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientHpChangeDetailValSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientHpChangeDetailVal) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
