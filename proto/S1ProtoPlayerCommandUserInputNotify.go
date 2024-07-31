package proto

const S1ProtoPlayerCommandUserInputNotifySize = 1

func (s *S1ProtoPlayerCommandUserInputNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoPlayerCommandUserInputNotifySize)
	bytes[0] = s.Value
	return bytes
}

func (s *S1ProtoPlayerCommandUserInputNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoPlayerCommandUserInputNotifySize {
		return InvalidDataError
	}

	s.Value = bytes[0]
	return nil
}
