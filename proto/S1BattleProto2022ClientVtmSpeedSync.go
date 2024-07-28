package proto

type S1BattleProto2022ClientVtmSpeedSync struct {
	CurrentTxTemperature int32
	Clientid             int32
	CurrentFreq          int32
	TxConnect            int32
	CurrentSpeedRate     int32
}

const S1BattleProto2022ClientVtmSpeedSyncSize = 20

func (s *S1BattleProto2022ClientVtmSpeedSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientVtmSpeedSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientVtmSpeedSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
