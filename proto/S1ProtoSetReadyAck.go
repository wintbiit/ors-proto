package proto

type S1ProtoSetReadyAck struct {
	ResultId int32
}

const S1ProtoSetReadyAckSize = 4

func (s *S1ProtoSetReadyAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetReadyAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetReadyAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
