package proto

type S1ProtoUserStateAck struct {
	Online int32
}

const S1ProtoUserStateAckSize = 4

func (s *S1ProtoUserStateAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoUserStateAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoUserStateAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
