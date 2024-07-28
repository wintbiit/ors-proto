package proto

type S1BattleProto2022TestTimeDelayDownComplete struct {
	Count uint32
}

const S1BattleProto2022TestTimeDelayDownCompleteSize = 4

func (s *S1BattleProto2022TestTimeDelayDownComplete) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022TestTimeDelayDownCompleteSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022TestTimeDelayDownComplete) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
