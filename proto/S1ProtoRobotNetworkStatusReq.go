package proto

const S1ProtoRobotNetworkStatusReqSize = 1

func (s *S1ProtoRobotNetworkStatusReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoRobotNetworkStatusReqSize)
	bytes[0] = s.Nouse
	return bytes
}

func (s *S1ProtoRobotNetworkStatusReq) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoRobotNetworkStatusReqSize {
		return InvalidDataError
	}

	s.Nouse = bytes[0]
	return nil
}
