package proto

import "encoding/binary"

const S1ProtoRobotNetworkStatusAckSize = 4

func (s *S1ProtoRobotNetworkStatusAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoRobotNetworkStatusAckSize)
	binary.LittleEndian.PutUint16(bytes[0:], s.WifiDownlinkPlr)
	binary.LittleEndian.PutUint16(bytes[2:], s.VtDownlinkPlr)
	return bytes
}

func (s *S1ProtoRobotNetworkStatusAck) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoRobotNetworkStatusAckSize {
		return InvalidDataError
	}

	s.WifiDownlinkPlr = binary.LittleEndian.Uint16(bytes[0:])
	s.VtDownlinkPlr = binary.LittleEndian.Uint16(bytes[2:])
	return nil
}
