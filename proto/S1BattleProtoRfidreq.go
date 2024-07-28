package proto

const S1BattleProtoRfidreqSize = 16

func (s *S1BattleProtoRfidreq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoRfidreqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoRfidreq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
