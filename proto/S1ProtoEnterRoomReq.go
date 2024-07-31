package proto

const S1ProtoEnterRoomReqSize = 2

func (s *S1ProtoEnterRoomReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoEnterRoomReqSize)
	bytes[0] = s.Nouse1
	bytes[1] = s.Nouse2
	return bytes
}

func (s *S1ProtoEnterRoomReq) Deserialize(bytes []byte) error {
	s.Nouse1 = bytes[0]
	s.Nouse2 = bytes[1]
	return nil
}
