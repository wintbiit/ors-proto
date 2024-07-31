package proto

import "encoding/binary"

const S1ProtoSetTeamNotifySize = 10

func (s *S1ProtoSetTeamNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTeamNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:8], s.Uid)
	bytes[8] = s.TeamId
	bytes[9] = s.SeatIndex
	return bytes
}

func (s *S1ProtoSetTeamNotify) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:8])
	s.TeamId = bytes[8]
	s.SeatIndex = bytes[9]
	return nil
}
