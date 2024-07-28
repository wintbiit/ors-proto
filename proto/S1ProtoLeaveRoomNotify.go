package proto

type S1ProtoLeaveRoomNotify struct {
	Uid    uint64
	Willgo e_Role_Location
}

const S1ProtoLeaveRoomNotifySize = 9

func (s *S1ProtoLeaveRoomNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoLeaveRoomNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLeaveRoomNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
