package proto

type S1BattleProto2022EnvDevicesDescripeAck struct {
	Name S1BattleProto2022RobotDynamicUIStr
}

const S1BattleProto2022EnvDevicesDescripeAckSize = 384

func (s *S1BattleProto2022EnvDevicesDescripeAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnvDevicesDescripeAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnvDevicesDescripeAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
