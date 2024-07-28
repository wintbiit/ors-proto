package proto

type S1ProtoEnterRoomNotify struct {
	Uid uint64
}

const S1ProtoEnterRoomNotifySize = 8

func (s *S1ProtoEnterRoomNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoEnterRoomNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoEnterRoomNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
