package proto

import "encoding/binary"

const S1ProtoEnterRoomAckSize = 12

func (s *S1ProtoEnterRoomAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoEnterRoomAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	binary.LittleEndian.PutUint64(bytes[4:], s.Uid)
	return bytes
}

func (s *S1ProtoEnterRoomAck) Deserialize(bytes []byte) error {
	s.ResultId = EEnterRoomAckResultType(binary.LittleEndian.Uint32(bytes[0:]))
	s.Uid = binary.LittleEndian.Uint64(bytes[4:])
	return nil
}
