package proto

type S1BattleProto2022MapClickInfoNotify struct {
	TeamId         byte
	RobotidListLen byte
	RobotidList    []byte
	Mode           byte
	Ascii          byte
	ScreenX        uint16
	IsSendAll      byte
	EnemyId        byte
	Type           byte
	ScreenY        uint16
	MapX           float32
	MapY           float32
}

const S1BattleProto2022MapClickInfoNotifySize = 25

func (s *S1BattleProto2022MapClickInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022MapClickInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022MapClickInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
