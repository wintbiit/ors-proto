package proto

import "encoding/binary"

const S1ProtoRoomPauseNotifySize = 4

func (s *S1ProtoRoomPauseNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoRoomPauseNotifySize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.TimeElapse))
	return bytes
}

func (s *S1ProtoRoomPauseNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoRoomPauseNotifySize {
		return InvalidDataError
	}

	s.TimeElapse = int32(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
