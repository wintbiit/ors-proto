package proto

import "encoding/binary"

const S1ProtoSetTeamAckSize = 6

func (s *S1ProtoSetTeamAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamAckSize)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(s.ResultId))
	bytes[4] = s.TeamId
	bytes[5] = s.SeatIndex
	return bytes
}

func (s *S1ProtoSetTeamAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:4]))
	s.TeamId = bytes[4]
	s.SeatIndex = bytes[5]
	return nil
}
