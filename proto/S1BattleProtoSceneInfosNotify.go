package proto

type S1BattleProtoSceneInfosNotify struct {
	PlayersCount     int32
	PlayersUid       []uint64
	PlayersTid       []uint32
	PlayersTeamId    []byte
	PlayersSeatIndex []byte
}

const S1BattleProtoSceneInfosNotifySize = 4

func (s *S1BattleProtoSceneInfosNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoSceneInfosNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoSceneInfosNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
