package proto

const S1BattleProtoArmorHitReqSize = 24

func (s *S1BattleProtoArmorHitReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoArmorHitReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoArmorHitReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
