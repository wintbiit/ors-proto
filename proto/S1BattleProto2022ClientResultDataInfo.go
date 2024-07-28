package proto

type S1BattleProto2022ClientResultDataInfo struct {
	ScoreRed                     int32
	RedCurrentHp                 int32
	BlueCurrentHp                int32
	RedWarning1                  int32
	RecordsListLen               int32
	BlueWarning1                 int32
	BlueHits                     int32
	RedRuneCount                 int32
	RedDartHitCount              int32
	RedShootSmallCount           int32
	BlueShootBigCount            int32
	RedUavCallCount              byte
	GameId                       int64
	GameOverReasonId             byte
	ScoreBlue                    int32
	RedTotalHp                   int32
	BlueWarning2                 int32
	BlueShootSmallCount          int32
	BlueSnipeSuccCount           byte
	GameOverReason               []byte
	GameOrder                    byte
	TotalRound                   byte
	BlueTotalHp                  int32
	BlueGuardLeftLives           byte
	RobotsNum                    int32
	WebGameId                    int32
	RedHits                      int32
	BlueDartHitCount             int32
	RedShootBigCount             int32
	GuardFixedLives              byte
	RedRadarBuffDoubleUsedCount  byte
	BlueRadarBuffDoubleUsedCount byte
	StartTime                    int64
	DurationTime                 int32
	Winner                       byte
	RedWarning3                  int32
	BlueWarning3                 int32
	BlueRuneCount                int32
	RobotsInfos                  []S1BattleProto2022ClientRobotDataInfo
	Records                      []int32
	RedWarning2                  int32
	RedKillCount                 int32
	GameOverReasonListLen        int32
	BlueKillCount                int32
	RedGuardLeftLives            byte
	RedSnipeSuccCount            byte
	BlueUavCallCount             byte
}

const S1BattleProto2022ClientResultDataInfoSize = 1323

func (s *S1BattleProto2022ClientResultDataInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientResultDataInfoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientResultDataInfo) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
