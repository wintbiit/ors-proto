package proto

type S1BattleProto2022ReLoginDataSync struct {
	BuffVisibles []byte
	BuffMaxTime  []float32
	BuffLeftTime []float32
	RoleS0Id     uint64
	BuffCount    byte
	BuffUids     []uint64
	BuffTids     []uint32
	BuffLevels   []uint32
}

const S1BattleProto2022ReLoginDataSyncSize = 9

func (s *S1BattleProto2022ReLoginDataSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ReLoginDataSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ReLoginDataSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
