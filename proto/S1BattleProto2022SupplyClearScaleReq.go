package proto

type S1BattleProto2022SupplyClearScaleReq struct {
	Placeholder byte
}

const S1BattleProto2022SupplyClearScaleReqSize = 1

func (s *S1BattleProto2022SupplyClearScaleReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SupplyClearScaleReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SupplyClearScaleReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
