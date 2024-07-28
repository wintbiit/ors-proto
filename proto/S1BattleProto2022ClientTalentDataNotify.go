package proto

type S1BattleProto2022ClientTalentDataNotify struct {
	IsBalance           byte
	IsSemiAutomaticCtrl byte
	IsTempManualCtrl    byte
	Robotid             byte
	Data                []性能体系_数据
}

const S1BattleProto2022ClientTalentDataNotifySize = 172

func (s *S1BattleProto2022ClientTalentDataNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientTalentDataNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientTalentDataNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
