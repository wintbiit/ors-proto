package proto

type S1BattleProto2022StuCustomControlData struct {
	Cmd     uint16
	ListLen byte
	Data    []byte
}

const S1BattleProto2022StuCustomControlDataSize = 33

func (s *S1BattleProto2022StuCustomControlData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuCustomControlDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuCustomControlData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
