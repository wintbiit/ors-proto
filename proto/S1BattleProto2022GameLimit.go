package proto

type S1BattleProto2022GameLimit struct {
	PowerBuffer1          uint16
	PowerMaxhurt          uint16
	ShooterFreqLimitCount int32
	FreqLimits            []S1BattleProto2022ShooterFreqLimit
	ShooterLimits         []S1BattleProto2022ShooterLimit
	HeatLimitsCount       int32
	HeatLimits            []S1BattleProto2022HeatLimit
	PowerThreshold        uint16
	PowerBuffer2          uint16
	PowerHurtTabCount     int32
	PowerHurtTabs         []byte
	ShooterLimitsCount    int32
}

const S1BattleProto2022GameLimitSize = 20

func (s *S1BattleProto2022GameLimit) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022GameLimitSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022GameLimit) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
