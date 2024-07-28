package proto

type S1ProtoLeaveRoomAck struct {
	ResultId e_LeaveRoomAck_ResultType
}

const S1ProtoLeaveRoomAckSize = 4

func (s *S1ProtoLeaveRoomAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLeaveRoomAckSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoLeaveRoomAck) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
