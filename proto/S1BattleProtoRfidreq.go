package proto

type S1BattleProtoRfidreq struct {
	Data2     byte
	Data4     byte
	Timestamp uint64
	Type      byte
	Area      byte
	Data0     byte
	Data1     byte
	Data3     byte
	Data5     byte
}

const S1BattleProtoRfidreqSize = 16

func (s *S1BattleProtoRfidreq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoRfidreqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoRfidreq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
