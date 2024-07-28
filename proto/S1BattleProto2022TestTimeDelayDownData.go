package proto

type S1BattleProto2022TestTimeDelayDownData struct {
	Seq      uint32
	CurTime  int64
	WifiTime uint32
	DataLen  int32
	Datas    []byte
}

const S1BattleProto2022TestTimeDelayDownDataSize = 388

func (s *S1BattleProto2022TestTimeDelayDownData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022TestTimeDelayDownDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022TestTimeDelayDownData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
