package proto

import (
	"encoding/binary"
	"math"
)

const S1BattleProto2022AirplaneStatusNotifySize = 29

func (s *S1BattleProto2022AirplaneStatusNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022AirplaneStatusNotifySize)
	bytes[0] = s.Robotid
	binary.LittleEndian.PutUint32(bytes[1:], uint32(s.Energy))
	bytes[5] = s.IsFire
	binary.LittleEndian.PutUint16(bytes[6:], uint16(s.Curbullet))
	binary.LittleEndian.PutUint16(bytes[8:], uint16(s.Maxbullet))
	binary.LittleEndian.PutUint32(bytes[10:], math.Float32bits(s.Curshoottime))
	binary.LittleEndian.PutUint32(bytes[14:], math.Float32bits(s.Fixshoottime))
	binary.LittleEndian.PutUint32(bytes[18:], math.Float32bits(s.Cdleft))
	binary.LittleEndian.PutUint16(bytes[22:], s.Countleft)
	binary.LittleEndian.PutUint32(bytes[24:], uint32(s.Cdrefreshcost))
	bytes[28] = s.Isincd

	return bytes
}

func (s *S1BattleProto2022AirplaneStatusNotify) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022AirplaneStatusNotifySize {
		return InvalidDataError
	}

	s.Robotid = bytes[0]
	s.Energy = int32(binary.LittleEndian.Uint32(bytes[1:]))
	s.IsFire = bytes[5]
	s.Curbullet = int16(binary.LittleEndian.Uint16(bytes[6:]))
	s.Maxbullet = int16(binary.LittleEndian.Uint16(bytes[8:]))
	s.Curshoottime = math.Float32frombits(binary.LittleEndian.Uint32(bytes[10:]))
	s.Fixshoottime = math.Float32frombits(binary.LittleEndian.Uint32(bytes[14:]))
	s.Cdleft = math.Float32frombits(binary.LittleEndian.Uint32(bytes[18:]))
	s.Countleft = binary.LittleEndian.Uint16(bytes[22:])
	s.Cdrefreshcost = int32(binary.LittleEndian.Uint32(bytes[24:]))
	s.Isincd = bytes[28]

	return nil
}
