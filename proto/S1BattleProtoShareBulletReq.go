package proto

const S1BattleProtoShareBulletReqSize = 2

func (s *S1BattleProtoShareBulletReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoShareBulletReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoShareBulletReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
