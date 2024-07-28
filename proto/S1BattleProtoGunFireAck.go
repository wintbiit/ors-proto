package proto

type S1BattleProtoGunFireAck struct {
	Result int32
}

const S1BattleProtoGunFireAckSize = 4

func (s *S1BattleProtoGunFireAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGunFireAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGunFireAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
