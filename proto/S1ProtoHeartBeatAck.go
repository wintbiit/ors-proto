package proto

type S1ProtoHeartBeatAck struct {
	S0Clientid byte
	ResultId   int32
}

const S1ProtoHeartBeatAckSize = 5

func (s *S1ProtoHeartBeatAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoHeartBeatAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoHeartBeatAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
