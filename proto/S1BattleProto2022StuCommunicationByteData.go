package proto

type S1BattleProto2022StuCommunicationByteData struct {
	ByteDataList    []byte
	S0HeaderBodyLen uint16
	Dataid          uint16
	Senderid        uint16
	Receiverid      uint16
	Cmd             uint16
	ByteDataListLen byte
}

const S1BattleProto2022StuCommunicationByteDataSize = 139

func (s *S1BattleProto2022StuCommunicationByteData) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022StuCommunicationByteDataSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022StuCommunicationByteData) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
