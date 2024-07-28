package proto

type S1ProtoRoomStateNotify struct {
	State    e_Room_StateType
	TimeLeft int32
	IsPause  int32
}

const S1ProtoRoomStateNotifySize = 12

func (s *S1ProtoRoomStateNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoRoomStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoRoomStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
