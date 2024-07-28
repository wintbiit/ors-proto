package proto

type S1ProtoHeader struct {
	AckType    byte
	Nouse1     byte
	Nouse2     byte
	ProtoId    uint16
	ProtoType  uint16
	DataLen    uint32
	SequenceId byte
}

const S1ProtoHeaderSize = 12

func (s *S1ProtoHeader) Serialize() []byte {
	bytes := make([]byte, S1ProtoHeaderSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoHeader) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
