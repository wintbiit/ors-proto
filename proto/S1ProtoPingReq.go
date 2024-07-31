package proto

const S1ProtoPingReqSize = 1

func (s *S1ProtoPingReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoPingReqSize)
	bytes[0] = s.Nouse
	return bytes
}

func (s *S1ProtoPingReq) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoPingReqSize {
		return InvalidDataError
	}

	s.Nouse = bytes[0]
	return nil
}
