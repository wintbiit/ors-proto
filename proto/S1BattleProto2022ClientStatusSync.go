package proto

type S1BattleProto2022ClientStatusSync struct {
	ClientId int32
	Status   int32
}

const S1BattleProto2022ClientStatusSyncSize = 8

func (s *S1BattleProto2022ClientStatusSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientStatusSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientStatusSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
