package proto

type S1BattleProto2023ClientPenaltyTableInfos struct {
	UploadType byte
	InfosLen   byte
	Infos      []S1BattleProto2023ClientPenaltyInfo
}

const S1BattleProto2023ClientPenaltyTableInfosSize = 7282

func (s *S1BattleProto2023ClientPenaltyTableInfos) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientPenaltyTableInfosSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientPenaltyTableInfos) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
