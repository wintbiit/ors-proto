package proto

type S1BattleProto2022ClientBuffSlot struct {
	BuffItemsLen int32
	BuffItems    []S1BattleProto2022ClientBuffItem
}

const S1BattleProto2022ClientBuffSlotSize = 44

func (s *S1BattleProto2022ClientBuffSlot) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBuffSlotSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientBuffSlot) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
