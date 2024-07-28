package proto

const S1ProtoSetTeamAckSize = 6

func (s *S1ProtoSetTeamAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoSetTeamAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
