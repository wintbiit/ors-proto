package proto

type S1ProtoClientSyncNotify struct {
	OnlineModule      []uint64
	Uid               uint64
	ConnClientState   int32
	ConnRobotState    int32
	Battery           int32
	SignalQuality     int32
	BattleMode        byte
	OnlineModuleCount int32
}

const S1ProtoClientSyncNotifySize = 29

func (s *S1ProtoClientSyncNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoClientSyncNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoClientSyncNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
