package proto

type S1BattleProto2022ClientBuyBulletReq struct {
	Type  byte
	Count byte
}

const S1BattleProto2022ClientBuyBulletReqSize = 2

func (s *S1BattleProto2022ClientBuyBulletReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBuyBulletReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientBuyBulletReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
