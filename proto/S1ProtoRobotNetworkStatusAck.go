package proto

type S1ProtoRobotNetworkStatusAck struct {
	WifiDownlinkPlr uint16
	VtDownlinkPlr   uint16
}

const S1ProtoRobotNetworkStatusAckSize = 4

func (s *S1ProtoRobotNetworkStatusAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoRobotNetworkStatusAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoRobotNetworkStatusAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
