package proto

type S1BattleProtoReloginNotify struct {
	PlayersMaxDurability            []int32
	PlayersUserInputState           []byte
	SelfTeamId                      byte
	SceneStateTimeLeft              int32
	SceneTeamsPlayerMax             int32
	SceneState                      int32
	PlayersTid                      []uint32
	PlayersUserCustomActionState    []byte
	PlayersLockScreenState          []byte
	SelfUid                         uint64
	PlayersHeat                     []float32
	PlayersMaxArmor                 []int32
	DevicesBaseStationState         []byte
	DevicesBaseStationStateLeftTime []float32
	DevicesRuneState                []byte
	PlayersTotalBuffCount           uint32
	PlayersTeamId                   []byte
	PlayersSeatIndex                []byte
	PlayersBuffCount                []uint32
	PlayersBuffUid                  [][]uint64
	PlayersMoveCtrlState            []byte
	PlayersAutoModeCtrlState        []byte
	DevicesRuneStateLeftTime        []float32
	PlayersState                    []byte
	PlayersMaxHeat                  []float32
	PlayersBulletNum                []uint32
	WifiName                        string
	WifiPassword                    string
	PlayersCdMaxTime                []float32
	PlayersMaxHp                    []uint32
	PlayersArmor                    []int32
	PlayersBuffLeftTime             [][]float32
	ScenePlayersCount               int32
	TeamsTotalDamage                []uint32
	PlayersUid                      []uint64
	SelfTid                         uint32
	SceneTeamsCount                 int32
	PlayersCdLeftTime               []float32
	PlayersDurability               []int32
	PlayersBuffTid                  [][]uint32
	PlayersBuffLevel                [][]uint32
	PlayersGunCtrlState             []byte
	SelfSeatIndex                   byte
	PlayersHp                       []uint32
	PlayersMaxBulletNum             []uint32
}

const S1BattleProtoReloginNotifySize = 102

func (s *S1BattleProtoReloginNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoReloginNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoReloginNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
