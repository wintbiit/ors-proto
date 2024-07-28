package proto

type S1ProtoMarkDetectResultAck struct {
	Result   int32
	X        float32
	Y        float32
	W        float32
	H        float32
	Color    byte
	MarkerId uint16
	Distance uint16
}

const S1ProtoMarkDetectResultAckSize = 25

func (s *S1ProtoMarkDetectResultAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoMarkDetectResultAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoMarkDetectResultAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
