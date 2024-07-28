package proto

type S1ClientTransferRobotProto struct {
	Data    []byte
	DataLen int32
}

const S1ClientTransferRobotProtoSize = 4

func (s *S1ClientTransferRobotProto) Serialize() []byte {
	bytes := make([]byte, S1ClientTransferRobotProtoSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ClientTransferRobotProto) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
