package proto

type S1BattleProto2022IoctrSetActuator struct {
	Actuator0  byte
	Actuator1  byte
	ModuleId   byte
	ModuleType byte
}

const S1BattleProto2022IoctrSetActuatorSize = 4

func (s *S1BattleProto2022IoctrSetActuator) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetActuatorSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetActuator) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
