package proto

type S1BattleProto2022SupplyExportReq struct {
	ExportEnable byte
}

const S1BattleProto2022SupplyExportReqSize = 1

func (s *S1BattleProto2022SupplyExportReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SupplyExportReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SupplyExportReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
