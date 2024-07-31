package proto

import (
	"encoding/binary"
	"math"
)

const S1BattleProto2022ArmorParaConfigItemSize = 12

func (s *S1BattleProto2022ArmorParaConfigItem) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ArmorParaConfigItemSize)
	binary.LittleEndian.PutUint32(bytes[0:], math.Float32bits(s.TiggerPress))
	binary.LittleEndian.PutUint32(bytes[4:], math.Float32bits(s.BulletMaxPress))
	binary.LittleEndian.PutUint32(bytes[8:], math.Float32bits(s.GolfMinPress))

	return bytes
}

func (s *S1BattleProto2022ArmorParaConfigItem) Deserialize(bytes []byte) error {
	s.TiggerPress = math.Float32frombits(binary.LittleEndian.Uint32(bytes[0:]))
	s.BulletMaxPress = math.Float32frombits(binary.LittleEndian.Uint32(bytes[4:]))
	s.GolfMinPress = math.Float32frombits(binary.LittleEndian.Uint32(bytes[8:]))
	return nil
}
