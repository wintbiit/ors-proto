package proto

type S1BattleProto2022HeatLimit struct {
	HeatCdFreq       byte
	HeatCdStep       uint16
	HeatSpdThreshold uint16
	HeatSpdMaxHurt   uint16
	HeatHurtTabCount int32
	HeatHurtTab      []byte
}

const S1BattleProto2022HeatLimitSize = 23

func (s *S1BattleProto2022HeatLimit) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022HeatLimitSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022HeatLimit) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
