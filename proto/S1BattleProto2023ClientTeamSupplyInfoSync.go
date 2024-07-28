package proto

type S1BattleProto2023ClientTeamSupplyInfoSync struct {
	Teamid                          byte
	ExchangeSmallBulletPackageCount byte
	ExchangeBigBulletPackageCount   byte
	SmallBulletPackageUnitPrice     float32
	SmallBulletUnitPrice            float32
	SentryControlPrice              float32
	ExchangeRemoteHealCount         byte
	SentryRemainRevivalCount        byte
	BigBulletPackageUnitPrice       float32
	BigBulletUnitPrice              float32
	CurrentRemoteHealPrice          float32
}

const S1BattleProto2023ClientTeamSupplyInfoSyncSize = 29

func (s *S1BattleProto2023ClientTeamSupplyInfoSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientTeamSupplyInfoSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientTeamSupplyInfoSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
