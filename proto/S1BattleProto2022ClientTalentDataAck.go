package proto

type S1BattleProto2022ClientTalentDataAck struct {
	Result int32
}

const S1BattleProto2022ClientTalentDataAckSize = 4

func (s *S1BattleProto2022ClientTalentDataAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientTalentDataAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientTalentDataAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
