package proto

import "encoding/binary"

const S1ProtoSetReadyNotifySize = 9

func (s *S1ProtoSetReadyNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetReadyNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:], s.Uid)
	bytes[8] = s.State
	return bytes
}

func (s *S1ProtoSetReadyNotify) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:])
	s.State = bytes[8]
	return nil
}
