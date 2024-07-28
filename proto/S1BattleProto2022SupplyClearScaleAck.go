package proto

type S1BattleProto2022SupplyClearScaleAck struct {
	Placeholder byte
}

const S1BattleProto2022SupplyClearScaleAckSize = 1

func (s *S1BattleProto2022SupplyClearScaleAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SupplyClearScaleAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SupplyClearScaleAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
