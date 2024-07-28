package proto

type S1BattleProtoClientSyncNotify struct {
	ConnRobotState  int32
	Battery         int32
	SignalQuality   int32
	BattleMode      byte
	Uid             uint64
	ConnClientState int32
}

const S1BattleProtoClientSyncNotifySize = 25

func (s *S1BattleProtoClientSyncNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoClientSyncNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoClientSyncNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
