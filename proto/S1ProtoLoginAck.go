package proto

type S1ProtoLoginAck struct {
	Uid      uint64
	Token    string
	ResultId e_LoginAck_ResultType
}

const S1ProtoLoginAckSize = 44

func (s *S1ProtoLoginAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLoginAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLoginAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
