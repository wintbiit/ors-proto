package proto

type S1BattleProtoSceneStateNotify struct {
	TimeLeft        int32
	TeamCount       int32
	TeamHurtSum     []int32
	TeamBekillCount []int32
	State           e_Scene_StateType
}

const S1BattleProtoSceneStateNotifySize = 12

func (s *S1BattleProtoSceneStateNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoSceneStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoSceneStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
