package proto

type S1ProtoRoomPauseNotify struct {
	TimeElapse int32
}

const S1ProtoRoomPauseNotifySize = 4

func (s *S1ProtoRoomPauseNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoRoomPauseNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoRoomPauseNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
