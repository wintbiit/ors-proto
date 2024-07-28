package proto

type S1ProtoReLoginReq struct {
	Account  string
	Password string
	Version  float32
	Tid      uint32
}

const S1ProtoReLoginReqSize = 72

func (s *S1ProtoReLoginReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoReLoginReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
