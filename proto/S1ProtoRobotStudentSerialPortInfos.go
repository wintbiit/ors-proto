package proto

type S1ProtoRobotStudentSerialPortInfos struct {
	OtherType  byte
	OtherPlr   uint16
	OtherDelay uint16
}

const S1ProtoRobotStudentSerialPortInfosSize = 5

func (s *S1ProtoRobotStudentSerialPortInfos) Serialize() []byte {
	bytes := make([]byte, S1ProtoRobotStudentSerialPortInfosSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoRobotStudentSerialPortInfos) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
