package proto

const S1BattleProtoDeviceModuleReqSize = 12

func (s *S1BattleProtoDeviceModuleReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoDeviceModuleReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoDeviceModuleReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
