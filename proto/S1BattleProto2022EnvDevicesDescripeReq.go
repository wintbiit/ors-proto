package proto

type S1BattleProto2022EnvDevicesDescripeReq struct {
	Res byte
}

const S1BattleProto2022EnvDevicesDescripeReqSize = 1

func (s *S1BattleProto2022EnvDevicesDescripeReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnvDevicesDescripeReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnvDevicesDescripeReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
