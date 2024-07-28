package proto

type S1ProtoMarkDetectResult struct {
	X        float32
	Y        float32
	H        float32
	T44C2M   []float32
	Color    byte
	MarkerId uint16
	W        float32
	Pitch    float32
	Yaw      float32
	Roll     float32
	Distance uint16
}

const S1ProtoMarkDetectResultSize = 97

func (s *S1ProtoMarkDetectResult) Serialize() []byte {
	bytes := make([]byte, S1ProtoMarkDetectResultSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoMarkDetectResult) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
