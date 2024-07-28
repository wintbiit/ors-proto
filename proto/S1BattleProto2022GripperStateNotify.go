package proto

type S1BattleProto2022GripperStateNotify struct {
	GrippersStatesListLen byte
	GripperStates         []byte
	IsConnect             byte
	IsHasMineralsListLen  byte
	IsHasMineralsList     []byte
	ErrorsListLen         byte
	ErrorsList            []byte
}

const S1BattleProto2022GripperStateNotifySize = 19

func (s *S1BattleProto2022GripperStateNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022GripperStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022GripperStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
