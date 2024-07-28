package proto

type S1BattleProto2022StuIcrabuffDebuffZoneStatus struct {
	F2Status     byte
	F2StatusInfo byte
	F3Status     byte
	F6StatusInfo byte
	F1StatusInfo byte
	F1Status     byte
	F3StatusInfo byte
	F4Status     byte
	F4StatusInfo byte
	F5Status     byte
	F5StatusInfo byte
	F6Status     byte
	Cmd          uint16
}

const S1BattleProto2022StuIcrabuffDebuffZoneStatusSize = 14

func (s *S1BattleProto2022StuIcrabuffDebuffZoneStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuIcrabuffDebuffZoneStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuIcrabuffDebuffZoneStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
