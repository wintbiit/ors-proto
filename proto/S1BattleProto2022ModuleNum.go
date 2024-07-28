package proto

type S1BattleProto2022ModuleNum struct {
	CameraNum      byte
	SmallArmorImp  byte
	BigArmorImp    byte
	LightBrdImp    byte
	CapNum         byte
	SmallStrikeNum byte
	BigStrikeNum   byte
	ExtTypeBimp    byte
	LightBarNum    byte
	RfidImp        byte
	ExtTypeAnum    byte
	ExtTypeCimp    byte
	Res            byte
	SmallShotImp   byte
	ExtTypeAimp    byte
	SmallShootNum  byte
	BigShootNum    byte
	UwbNum         byte
	PowerBrdImp    byte
	ExtTypeBnum    byte
	ExtTypeCnum    byte
	CameraImp      byte
	UwbImp         byte
	CapImp         byte
	PowerNum       byte
	RfidNum        byte
	BigShotImp     byte
}

const S1BattleProto2022ModuleNumSize = 27

func (s *S1BattleProto2022ModuleNum) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ModuleNumSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ModuleNum) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
