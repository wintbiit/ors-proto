package proto

import (
	"encoding/binary"
	"math"
)

const S1BattleProto2023ClientTeamSupplyInfoSyncSize = 29

func (s *S1BattleProto2023ClientTeamSupplyInfoSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientTeamSupplyInfoSyncSize)
	bytes[0] = s.Teamid
	bytes[1] = s.ExchangeSmallBulletPackageCount
	bytes[2] = s.ExchangeBigBulletPackageCount
	bytes[3] = s.ExchangeRemoteHealCount
	bytes[4] = s.SentryRemainRevivalCount
	binary.LittleEndian.PutUint32(bytes[5:], math.Float32bits(s.SmallBulletPackageUnitPrice))
	binary.LittleEndian.PutUint32(bytes[9:], math.Float32bits(s.BigBulletPackageUnitPrice))
	binary.LittleEndian.PutUint32(bytes[13:], math.Float32bits(s.BigBulletUnitPrice))
	binary.LittleEndian.PutUint32(bytes[17:], math.Float32bits(s.SmallBulletUnitPrice))
	binary.LittleEndian.PutUint32(bytes[21:], math.Float32bits(s.CurrentRemoteHealPrice))
	binary.LittleEndian.PutUint32(bytes[25:], math.Float32bits(s.SentryControlPrice))
	return bytes
}

func (s *S1BattleProto2023ClientTeamSupplyInfoSync) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2023ClientTeamSupplyInfoSyncSize {
		return InvalidDataError
	}

	s.Teamid = bytes[0]
	s.ExchangeSmallBulletPackageCount = bytes[1]
	s.ExchangeBigBulletPackageCount = bytes[2]
	s.ExchangeRemoteHealCount = bytes[3]
	s.SentryRemainRevivalCount = bytes[4]
	s.SmallBulletPackageUnitPrice = math.Float32frombits(binary.LittleEndian.Uint32(bytes[5:]))
	s.BigBulletPackageUnitPrice = math.Float32frombits(binary.LittleEndian.Uint32(bytes[9:]))
	s.BigBulletUnitPrice = math.Float32frombits(binary.LittleEndian.Uint32(bytes[13:]))
	s.SmallBulletUnitPrice = math.Float32frombits(binary.LittleEndian.Uint32(bytes[17:]))
	s.CurrentRemoteHealPrice = math.Float32frombits(binary.LittleEndian.Uint32(bytes[21:]))
	s.SentryControlPrice = math.Float32frombits(binary.LittleEndian.Uint32(bytes[25:]))
	return nil
}
