package proto

type S1BattleProtoClientSyncReq struct {
	Uid             uint64
	ConnClientState int32
	ConnRobotState  int32
	Battery         int32
	SignalQuality   int32
	BattleMode      byte
}

const S1BattleProtoClientSyncReqSize = 25

func (s *S1BattleProtoClientSyncReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoClientSyncReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoClientSyncReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
