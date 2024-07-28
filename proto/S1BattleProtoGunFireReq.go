package proto

type S1BattleProtoGunFireReq struct {
	AttackerUid uint64
	GunType     byte
	GunSpeed    float32
}

const S1BattleProtoGunFireReqSize = 13

func (s *S1BattleProtoGunFireReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProtoGunFireReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProtoGunFireReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
