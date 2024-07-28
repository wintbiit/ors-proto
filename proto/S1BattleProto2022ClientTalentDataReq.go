package proto

type S1BattleProto2022ClientTalentDataReq struct {
	目标类型    string
	性能类型    string
	是否设置副武器 byte
	Robotid byte
}

const S1BattleProto2022ClientTalentDataReqSize = 66

func (s *S1BattleProto2022ClientTalentDataReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientTalentDataReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientTalentDataReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
