package proto

const S1ProtoEnterRoomReqSize = 2

func (s *S1ProtoEnterRoomReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoEnterRoomReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoEnterRoomReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
