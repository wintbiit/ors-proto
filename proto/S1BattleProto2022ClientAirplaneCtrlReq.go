package proto

const S1BattleProto2022ClientAirplaneCtrlReqSize = 2

func (s *S1BattleProto2022ClientAirplaneCtrlReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientAirplaneCtrlReqSize)
	bytes[0] = s.Robotid
	bytes[1] = s.ControlCode
	return bytes
}

func (s *S1BattleProto2022ClientAirplaneCtrlReq) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022ClientAirplaneCtrlReqSize {
		return InvalidDataError
	}

	s.Robotid = bytes[0]
	s.ControlCode = bytes[1]
	return nil
}
