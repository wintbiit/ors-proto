package proto

import "encoding/binary"

const S1ProtoEnterRoomNotifySize = 8

func (s *S1ProtoEnterRoomNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoEnterRoomNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:], s.Uid)
	return bytes
}

func (s *S1ProtoEnterRoomNotify) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:])
	return nil
}
