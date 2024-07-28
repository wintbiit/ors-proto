package proto

type S1ProtoReLoginAck struct {
	ResultId     e_LoginAck_ResultType
	Uid          uint64
	Token        string
	WifiName     string
	WifiPassword string
}

const S1ProtoReLoginAckSize = 108

func (s *S1ProtoReLoginAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoReLoginAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
