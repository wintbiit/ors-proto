package proto

import (
	"encoding/binary"
	"math"
)

const S1ProtoLoginReqSize = 84

func (s *S1ProtoLoginReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLoginReqSize)
	copy(bytes[0:32], s.Account)
	copy(bytes[32:64], s.Password)
	binary.LittleEndian.PutUint32(bytes[64:68], math.Float32bits(s.Version))
	binary.LittleEndian.PutUint32(bytes[68:72], s.Tid)
	binary.LittleEndian.PutUint32(bytes[72:76], s.Teamid)
	binary.LittleEndian.PutUint64(bytes[76:84], uint64(s.Hash))

	return bytes
}

func (s *S1ProtoLoginReq) Deserialize(bytes []byte) error {
	s.Account = string(bytes[0:32])
	s.Password = string(bytes[32:64])
	s.Version = math.Float32frombits(binary.LittleEndian.Uint32(bytes[64:68]))
	s.Tid = binary.LittleEndian.Uint32(bytes[68:72])
	s.Teamid = binary.LittleEndian.Uint32(bytes[72:76])
	s.Hash = int64(binary.LittleEndian.Uint64(bytes[76:84]))

	return nil
}
