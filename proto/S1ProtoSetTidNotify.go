package proto

import "encoding/binary"

const S1ProtoSetTidNotifySize = 12

func (s *S1ProtoSetTidNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTidNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:8], s.Uid)
	binary.LittleEndian.PutUint32(bytes[8:12], s.Tid)
	return bytes
}

func (s *S1ProtoSetTidNotify) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:8])
	s.Tid = binary.LittleEndian.Uint32(bytes[8:12])
	return nil
}
