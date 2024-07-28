package proto

type S1ProtoClientSyncReq struct {
	Battery           int32
	SignalQuality     int32
	BattleMode        byte
	OnlineModuleCount int32
	OnlineModule      []uint64
	Uid               uint64
	ConnClientState   int32
	ConnRobotState    int32
}

const S1ProtoClientSyncReqSize = 29

func (s *S1ProtoClientSyncReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoClientSyncReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoClientSyncReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
