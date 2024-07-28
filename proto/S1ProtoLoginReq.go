package proto

type S1ProtoLoginReq struct {
	Account  string
	Password string
	Version  float32
	Tid      uint32
	Teamid   uint32
	Hash     int64
}

const S1ProtoLoginReqSize = 84

func (s *S1ProtoLoginReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLoginReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLoginReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
