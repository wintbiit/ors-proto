package proto

type S1BattleProto2022UwbData struct {
	Rsv1       byte
	X          float32
	Y          float32
	Z          float32
	Compass    float32
	AnchorMask byte
	WifiEn     byte
	Rsv0       byte
}

const S1BattleProto2022UwbDataSize = 20

func (s *S1BattleProto2022UwbData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022UwbDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022UwbData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
