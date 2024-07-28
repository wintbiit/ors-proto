package proto

type S1BattleProto2022IoctrRmMotorsStatus struct {
	RealTemperatureListLen int32
	RealTemperatures       []byte
	ModuleId               byte
	ModuleType             byte
	RealPositionListLen    int32
	RealPositions          []int64
	RealSpeedListLen       int32
	RealSpeeds             []int16
}

const S1BattleProto2022IoctrRmMotorsStatusSize = 102

func (s *S1BattleProto2022IoctrRmMotorsStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrRmMotorsStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrRmMotorsStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
