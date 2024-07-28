package proto

type S1ProtoPingAck struct {
	ResultId int32
}

const S1ProtoPingAckSize = 4

func (s *S1ProtoPingAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoPingAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoPingAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
