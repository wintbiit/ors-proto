package proto

const S1ProtoRobotNetworkStatusReqSize = 1

func (s *S1ProtoRobotNetworkStatusReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoRobotNetworkStatusReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoRobotNetworkStatusReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
