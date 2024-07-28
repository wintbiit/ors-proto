package proto

const S1BattleProtoMarkerDetectAckSize = 1

func (s *S1BattleProtoMarkerDetectAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoMarkerDetectAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoMarkerDetectAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
