package proto

const S1ProtoSetWifiInfoNotifySize = 64

func (s *S1ProtoSetWifiInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoSetWifiInfoNotifySize)
	copy(bytes[0:32], s.WifiName)
	copy(bytes[32:64], s.WifiPassword)
	return bytes
}

func (s *S1ProtoSetWifiInfoNotify) Deserialize(bytes []byte) error {
	s.WifiName = string(bytes[0:32])
	s.WifiPassword = string(bytes[32:64])
	return nil
}
