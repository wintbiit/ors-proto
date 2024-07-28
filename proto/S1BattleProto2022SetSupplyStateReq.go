package proto

type S1BattleProto2022SetSupplyStateReq struct {
	ReportTime uint16
	ReportReq  byte
}

const S1BattleProto2022SetSupplyStateReqSize = 3

func (s *S1BattleProto2022SetSupplyStateReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SetSupplyStateReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SetSupplyStateReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
