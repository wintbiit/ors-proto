package proto

type S1BattleProto2022Energy2019ArmorHitUpload struct {
	Row     byte
	ArmorId byte
	HitType uint16
}

const S1BattleProto2022Energy2019ArmorHitUploadSize = 4

func (s *S1BattleProto2022Energy2019ArmorHitUpload) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022Energy2019ArmorHitUploadSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022Energy2019ArmorHitUpload) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
