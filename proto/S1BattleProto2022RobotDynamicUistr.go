package proto

type S1BattleProto2022RobotDynamicUistr struct {
	EnvDeviceXnameEn string
	EnvDeviceXnameZh string
	EnvDeviceYnameZh string
	EnvDeviceZnameEn string
	ExtModuleAnameZh string
	ExtModuleCnameEn string
	EnvDeviceYnameEn string
	EnvDeviceZnameZh string
	ExtModuleAnameEn string
	ExtModuleBnameEn string
	ExtModuleBnameZh string
	ExtModuleCnameZh string
}

const S1BattleProto2022RobotDynamicUistrSize = 384

func (s *S1BattleProto2022RobotDynamicUistr) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022RobotDynamicUistrSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022RobotDynamicUistr) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
