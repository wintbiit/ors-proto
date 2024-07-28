package proto

type S1BattleProto2022QueryRobotInfoResult struct {
	LoaderVersion   uint32
	AppVersion      uint32
	DeviceIdListLen int32
	DeviceIdList    []int32
	Reserved        uint32
	ModuleType      byte
	ModuleId        byte
}

const S1BattleProto2022QueryRobotInfoResultSize = 66

func (s *S1BattleProto2022QueryRobotInfoResult) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022QueryRobotInfoResultSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022QueryRobotInfoResult) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
