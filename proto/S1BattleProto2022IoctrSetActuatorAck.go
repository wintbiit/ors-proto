package proto

type S1BattleProto2022IoctrSetActuatorAck struct {
	ModuleType byte
	Actuator0  byte
	Actuator1  byte
	ModuleId   byte
}

const S1BattleProto2022IoctrSetActuatorAckSize = 4

func (s *S1BattleProto2022IoctrSetActuatorAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrSetActuatorAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrSetActuatorAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
