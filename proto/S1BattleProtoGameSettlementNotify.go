package proto

type S1BattleProtoGameSettlementNotify struct {
	PlayersHpMax    []int32
	PlayersBekilled []int32
	TeamHurtSum     []int32
	DurationTime    int32
	PlayersUid      []uint64
	PlayersHp       []int32
	Winreason       int32
	TeamCount       int32
	TeamAtkCount    []int32
	TeamPojiaCount  []int32
	TeamidWin       int32
	PlayersCount    int32
}

const S1BattleProtoGameSettlementNotifySize = 20

func (s *S1BattleProtoGameSettlementNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGameSettlementNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGameSettlementNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
