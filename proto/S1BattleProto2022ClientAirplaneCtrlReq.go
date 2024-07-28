package proto

type S1BattleProto2022ClientAirplaneCtrlReq struct {
	Robotid     byte
	ControlCode byte
}

const S1BattleProto2022ClientAirplaneCtrlReqSize = 2

func (s *S1BattleProto2022ClientAirplaneCtrlReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientAirplaneCtrlReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientAirplaneCtrlReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
