package proto

type S1ProtoReLoginRoomNotify struct {
	RoomStateTimeLeft  int32
	PlayersTeamId      []byte
	WifiName           string
	SelfTeamId         byte
	RoomPlayersCount   int32
	PlayersSeatIndex   []byte
	SelfUid            uint64
	SelfTid            uint32
	SelfSeatIndex      byte
	RoomState          int32
	RoomSeatIslock     int32
	RoomTeamsCount     int32
	RoomTeamsPlayerMax int32
	PlayersUid         []uint64
	PlayersTid         []uint32
	WifiPassword       string
}

const S1ProtoReLoginRoomNotifySize = 102

func (s *S1ProtoReLoginRoomNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginRoomNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoReLoginRoomNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
