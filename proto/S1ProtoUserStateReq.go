package proto

type S1ProtoUserStateReq struct {
	Account  string
	Password string
}

const S1ProtoUserStateReqSize = 64

func (s *S1ProtoUserStateReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoUserStateReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoUserStateReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
