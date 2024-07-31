package proto

const S1ProtoSetCurrMatchNotifySize = 160

func (s *S1ProtoSetCurrMatchNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetCurrMatchNotifySize)
	copy(bytes[0:], s.CurrMatch[:32])
	copy(bytes[32:], s.CurrToken[:128])
	return bytes
}

func (s *S1ProtoSetCurrMatchNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoSetCurrMatchNotifySize {
		return InvalidDataError
	}

	s.CurrMatch = string(bytes[0:32])
	s.CurrToken = string(bytes[32:128])
	return nil
}
