package proto

type S1BattleProtoDeviceModuleReq struct {
	ConnectedDevice      []uint64
	Uid                  uint64
	ProductType          uint16
	ConnectedDeviceCount uint16
}

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
