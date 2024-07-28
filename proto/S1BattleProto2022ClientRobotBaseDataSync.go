package proto

type S1BattleProto2022ClientRobotBaseDataSync struct {
	FixSmallHeatEnergyCoolRate  float32
	FixBigFreq                  float32
	CurExpAddRate               float32
	FixExpAddRate               float32
	FixHp                       int32
	FixPower                    float32
	ExpOffer                    float32
	FixExpUpgradeNeed           float32
	FixRebornCd                 int32
	FixBigSpeed                 float32
	FixSmallHeatEnergy2         float32
	UserId                      uint16
	FixPowerBufferRecoverable   float32
	FixPowerBufferUnrecoverable float32
	FixSmallSpeed               float32
	FixSmallHeatEnergy          float32
	CurRebornCd                 int32
	Type                        byte
	Level                       byte
	CpuId                       uint32
	FixSmallSpeed2              float32
	FixBigHeatEnergy            float32
	FixBigHeatEnergyCoolRate    float32
	Id                          byte
	FixSmallFreq                float32
	FixSmallFreq2               float32
	FixSmallHeatEnergyCoolRate2 float32
}

const S1BattleProto2022ClientRobotBaseDataSyncSize = 97

func (s *S1BattleProto2022ClientRobotBaseDataSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientRobotBaseDataSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientRobotBaseDataSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
