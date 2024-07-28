package proto

const S1BattleProtoFullSceneDataReqSize = 1

func (s *S1BattleProtoFullSceneDataReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoFullSceneDataReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoFullSceneDataReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
