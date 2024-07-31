package proto

import "encoding/binary"

const S1ProtoLeaveRoomNotifySize = 9

func (s *S1ProtoLeaveRoomNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoLeaveRoomNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:], s.Uid)
	bytes[8] = byte(s.Willgo)
	return bytes
}

func (s *S1ProtoLeaveRoomNotify) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:])
	s.Willgo = ERoleLocation(bytes[8])
	return nil
}
