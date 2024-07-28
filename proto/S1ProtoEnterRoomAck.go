package proto

type S1ProtoEnterRoomAck struct {
	ResultId e_EnterRoomAck_ResultType
	Uid      uint64
}

const S1ProtoEnterRoomAckSize = 12

func (s *S1ProtoEnterRoomAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoEnterRoomAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoEnterRoomAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
