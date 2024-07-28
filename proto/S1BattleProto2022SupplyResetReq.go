package proto

type S1BattleProto2022SupplyResetReq struct {
	Placeholder byte
}

const S1BattleProto2022SupplyResetReqSize = 1

func (s *S1BattleProto2022SupplyResetReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SupplyResetReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SupplyResetReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
