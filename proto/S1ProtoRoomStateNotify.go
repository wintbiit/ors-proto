package proto

import "encoding/binary"

const S1ProtoRoomStateNotifySize = 12

func (s *S1ProtoRoomStateNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoRoomStateNotifySize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.State))
	binary.LittleEndian.PutUint32(bytes[4:], uint32(s.TimeLeft))
	binary.LittleEndian.PutUint32(bytes[8:], uint32(s.IsPause))
	return bytes
}

func (s *S1ProtoRoomStateNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoRoomStateNotifySize {
		return InvalidDataError
	}

	s.State = ERoomStateType(binary.LittleEndian.Uint32(bytes[0:]))
	s.TimeLeft = int32(binary.LittleEndian.Uint32(bytes[4:]))
	s.IsPause = int32(binary.LittleEndian.Uint32(bytes[8:]))
	return nil
}
