package proto

import "encoding/binary"

const S1ProtoLeaveRoomAckSize = 4

func (s *S1ProtoLeaveRoomAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLeaveRoomAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	return bytes
}

func (s *S1ProtoLeaveRoomAck) Deserialize(bytes []byte) error {
	s.ResultId = ELeaveRoomAckResultType(binary.LittleEndian.Uint32(bytes[0:]))
	return nil
}
