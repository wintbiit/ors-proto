package proto

const S1ProtoRoomInfosNotifySize = 100

func (s *S1ProtoRoomInfosNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoRoomInfosNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoRoomInfosNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
