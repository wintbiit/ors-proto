package proto

type S1BattleProtoRfidnotify struct {
	Flag  byte
	Type  byte
	Data2 byte
	Data5 byte
	S0Id  byte
	Area  byte
	Data0 byte
	Data1 byte
	Data3 byte
	Data4 byte
}

const S1BattleProtoRfidnotifySize = 10

func (s *S1BattleProtoRfidnotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoRfidnotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoRfidnotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
