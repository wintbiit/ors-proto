package proto

import "encoding/binary"

const S1ProtoSetTidReqSize = 36

func (s *S1ProtoSetTidReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetTidReqSize)
	binary.LittleEndian.PutUint32(bytes[0:4], s.Tid)
	copy(bytes[4:36], s.Token)
	return bytes
}

func (s *S1ProtoSetTidReq) Deserialize(bytes []byte) error {
	s.Tid = binary.LittleEndian.Uint32(bytes[0:4])
	s.Token = string(bytes[4:36])
	return nil
}
