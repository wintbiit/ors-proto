package proto

type S1BattleProto2022ClientShielddata struct {
	RedShield     int32
	BlueShield    int32
	ShieldMax     int32
	RedHasShield  byte
	BlueHasShield byte
}

const S1BattleProto2022ClientShielddataSize = 14

func (s *S1BattleProto2022ClientShielddata) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientShielddataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientShielddata) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
