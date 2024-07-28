package proto

type S1BattleProto2022TestTimeDelayDownRecord struct {
	WifiLossTabS     []uint16
	GlobalDelay      uint32
	UartDelay        uint32
	UartLossCnt      byte
	WifiLossCnt      byte
	WifiLossTabCount int32
}

const S1BattleProto2022TestTimeDelayDownRecordSize = 388

func (s *S1BattleProto2022TestTimeDelayDownRecord) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022TestTimeDelayDownRecordSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022TestTimeDelayDownRecord) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
