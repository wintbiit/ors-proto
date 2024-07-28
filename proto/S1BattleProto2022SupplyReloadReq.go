package proto

type S1BattleProto2022SupplyReloadReq struct {
	ReloadEnable byte
}

const S1BattleProto2022SupplyReloadReqSize = 1

func (s *S1BattleProto2022SupplyReloadReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SupplyReloadReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SupplyReloadReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
