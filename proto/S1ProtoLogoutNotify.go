package proto

import "encoding/binary"

const S1ProtoLogoutNotifySize = 9

func (s *S1ProtoLogoutNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoLogoutNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:], s.Uid)
	bytes[8] = s.Reason
	return bytes
}

func (s *S1ProtoLogoutNotify) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:])
	s.Reason = bytes[8]
	return nil
}
