package proto

import "encoding/binary"

const S1BattleProto2022BaseLightingEffectSize = 4

func (s *S1BattleProto2022BaseLightingEffect) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022BaseLightingEffectSize)
	binary.LittleEndian.PutUint32(bytes[0:], s.LightColor)
	return bytes
}

func (s *S1BattleProto2022BaseLightingEffect) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022BaseLightingEffectSize {
		return InvalidDataError
	}

	s.LightColor = binary.LittleEndian.Uint32(bytes[0:])
	return nil
}
