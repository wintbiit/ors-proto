package proto

const S1BattleProto2022EnvHeartBeatAckSize = 4

func (s *S1BattleProto2022EnvHeartBeatAck) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022EnvHeartBeatAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022EnvHeartBeatAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
