package proto

const S1ProtoPlayerCommandDisconnectNotifySize = 9

func (s *S1ProtoPlayerCommandDisconnectNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoPlayerCommandDisconnectNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoPlayerCommandDisconnectNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
