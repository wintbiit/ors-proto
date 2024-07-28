package proto

type S1BattleProto2022ClientSafetySync struct {
	Lift2Connect  byte
	FlycatConnect byte
	Lift1State    byte
	Lift2State    byte
	CatError      byte
	FlycatBattery byte
	Id            byte
	Lift1Connect  byte
	FlycatState   byte
	Lift1Error    byte
	Lift2Error    byte
}

const S1BattleProto2022ClientSafetySyncSize = 11

func (s *S1BattleProto2022ClientSafetySync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientSafetySyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientSafetySync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
