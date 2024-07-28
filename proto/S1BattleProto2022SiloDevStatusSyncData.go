package proto

type S1BattleProto2022SiloDevStatusSyncData struct {
	SiloDevStatusSyncDataListLen int32
	SiloDevStatusDataList        []S1BattleProto2022SiloDevStatus
}

const S1BattleProto2022SiloDevStatusSyncDataSize = 92

func (s *S1BattleProto2022SiloDevStatusSyncData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022SiloDevStatusSyncDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022SiloDevStatusSyncData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
