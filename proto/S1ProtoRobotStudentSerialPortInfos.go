package proto

import "encoding/binary"

const S1ProtoRobotStudentSerialPortInfosSize = 5

func (s *S1ProtoRobotStudentSerialPortInfos) Serialize() []byte {
	bytes := make([]byte, S1ProtoRobotStudentSerialPortInfosSize)
	bytes[0] = s.OtherType
	binary.LittleEndian.PutUint16(bytes[1:], s.OtherPlr)
	binary.LittleEndian.PutUint16(bytes[3:], s.OtherDelay)

	return bytes
}

func (s *S1ProtoRobotStudentSerialPortInfos) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoRobotStudentSerialPortInfosSize {
		return InvalidDataError
	}

	s.OtherType = bytes[0]
	s.OtherPlr = binary.LittleEndian.Uint16(bytes[1:])
	s.OtherDelay = binary.LittleEndian.Uint16(bytes[3:])
	return nil
}
