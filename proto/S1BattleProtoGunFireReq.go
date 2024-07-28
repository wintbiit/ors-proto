package proto

const S1BattleProtoGunFireReqSize = 13

func (s *S1BattleProtoGunFireReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGunFireReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGunFireReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
