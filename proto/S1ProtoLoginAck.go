package proto

import "encoding/binary"

const S1ProtoLoginAckSize = 44

func (s *S1ProtoLoginAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoLoginAckSize)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(s.ResultId))
	binary.LittleEndian.PutUint64(bytes[4:12], s.Uid)
	copy(bytes[12:44], s.Token)
	return bytes
}

func (s *S1ProtoLoginAck) Deserialize(bytes []byte) error {
	s.ResultId = ELoginAckResultType(binary.LittleEndian.Uint32(bytes[0:4]))
	s.Uid = binary.LittleEndian.Uint64(bytes[4:12])
	s.Token = string(bytes[12:44])
	return nil
}
