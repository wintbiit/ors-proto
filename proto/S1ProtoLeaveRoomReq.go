package proto

const S1ProtoLeaveRoomReqSize = 1

func (s *S1ProtoLeaveRoomReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLeaveRoomReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLeaveRoomReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
