package proto

type S1BattleProto2022StuClientSendCmd struct {
	TargetPosY    float32
	TargetPosZ    float32
	CmdKeyboard   byte
	TargetRobotId uint16
	Cmd           uint16
	TargetPosX    float32
}

const S1BattleProto2022StuClientSendCmdSize = 17

func (s *S1BattleProto2022StuClientSendCmd) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuClientSendCmdSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuClientSendCmd) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
