package proto

type S1BattleProtoDeviceModuleNotify struct {
	AllDevice           []uint64
	AllDeviceConnection []byte
	AllDevicePrority    []byte
	Uid                 uint64
	ProductType         uint16
	AllDeviceCount      uint16
}

const S1BattleProtoDeviceModuleNotifySize = 12

func (s *S1BattleProtoDeviceModuleNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoDeviceModuleNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoDeviceModuleNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
