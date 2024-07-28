package proto

type S1BattleProto2022StuSiloInfo struct {
	LaunchOpeningStatus  byte
	AttackTarget         byte
	TargetChangeTime     uint16
	OperateLaunchCmdTime uint16
	Cmd                  uint16
}

const S1BattleProto2022StuSiloInfoSize = 8

func (s *S1BattleProto2022StuSiloInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuSiloInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuSiloInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
