package proto

type S1BattleProto2022ClientMineralInfoSync struct {
	IsRedConnect             byte
	IsBlueConnect            byte
	RedGoldCount             byte
	BlueInfraredStateListLen byte
	BlueInfraredState        []byte
	RedSilverCount           byte
	BlueGoldCount            byte
	BlueSilverCount          byte
	RedInfraredStateListLen  byte
	RedInfraredState         []byte
}

const S1BattleProto2022ClientMineralInfoSyncSize = 12

func (s *S1BattleProto2022ClientMineralInfoSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientMineralInfoSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientMineralInfoSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
