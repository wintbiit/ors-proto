package proto

type S1BattleProto2022VtmSpeedSync struct {
	CurrentCountryCode int32
	Clientid           int32
	CurrentFreq        int32
	TxConnect          int32
	CurrentSpeedRate   int32
}

const S1BattleProto2022VtmSpeedSyncSize = 20

func (s *S1BattleProto2022VtmSpeedSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022VtmSpeedSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022VtmSpeedSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
