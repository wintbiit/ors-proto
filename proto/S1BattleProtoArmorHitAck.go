package proto

const S1BattleProtoArmorHitAckSize = 4

func (s *S1BattleProtoArmorHitAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoArmorHitAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoArmorHitAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
