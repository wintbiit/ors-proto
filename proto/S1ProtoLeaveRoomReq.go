package proto

const S1ProtoLeaveRoomReqSize = 1

func (s *S1ProtoLeaveRoomReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLeaveRoomReqSize)
	bytes[0] = s.Nouse
	return bytes
}

func (s *S1ProtoLeaveRoomReq) Deserialize(bytes []byte) error {
	s.Nouse = bytes[0]
	return nil
}
