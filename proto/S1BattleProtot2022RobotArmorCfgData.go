package proto

type S1BattleProtot2022RobotArmorCfgData struct {
	ArmorParaDataCount int32
	ArmorParaDatas     []S1BattleProto2022ArmorParaConfigItem
}

const S1BattleProtot2022RobotArmorCfgDataSize = 4

func (s *S1BattleProtot2022RobotArmorCfgData) Serialize() []byte {
	bytes := make([]byte, S1BattleProtot2022RobotArmorCfgDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtot2022RobotArmorCfgData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
