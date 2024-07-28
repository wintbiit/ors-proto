package proto

type S1BattleProto2022BaseShellStatus struct {
	BaseStatus                byte
	BaseMotoOneBlockStatus    byte
	BaseMotoTwoStatus         byte
	BaseMotoTwoBlockStatus    byte
	BaseMotoThreeBlockStatus  byte
	Reserved0                 byte
	Reserved2                 byte
	BaseMotoOneStatus         byte
	BaseMotoOneOnlineStatus   byte
	BaseMotoTwoOnlineStatus   byte
	BaseMotoThreeStatus       byte
	BaseMotoThreeOnlineStatus byte
}

const S1BattleProto2022BaseShellStatusSize = 12

func (s *S1BattleProto2022BaseShellStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022BaseShellStatusSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022BaseShellStatus) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
