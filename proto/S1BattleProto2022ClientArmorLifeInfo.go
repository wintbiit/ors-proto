package proto

import "encoding/binary"

const S1BattleProto2022ClientArmorLifeInfoSize = 3

func (s *S1BattleProto2022ClientArmorLifeInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientArmorLifeInfoSize)
	binary.LittleEndian.PutUint16(bytes[0:], uint16(s.ModuleId))
	bytes[2] = s.LifeState
	return bytes
}

func (s *S1BattleProto2022ClientArmorLifeInfo) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022ClientArmorLifeInfoSize {
		return InvalidDataError
	}

	s.ModuleId = int16(binary.LittleEndian.Uint16(bytes[0:]))
	s.LifeState = bytes[2]
	return nil
}
