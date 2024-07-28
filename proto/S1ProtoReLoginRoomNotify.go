package proto

const S1ProtoReLoginRoomNotifySize = 102

func (s *S1ProtoReLoginRoomNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginRoomNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoReLoginRoomNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
