package proto

type S1BattleProtoMarkerDetectReq struct {
	MarkerId  byte
	MarkerStr string
	Uid       uint64
	Act       byte
}

const S1BattleProtoMarkerDetectReqSize = 42

func (s *S1BattleProtoMarkerDetectReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoMarkerDetectReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoMarkerDetectReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
