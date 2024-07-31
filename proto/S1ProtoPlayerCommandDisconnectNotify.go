package proto

import "encoding/binary"

const S1ProtoPlayerCommandDisconnectNotifySize = 9

func (s *S1ProtoPlayerCommandDisconnectNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoPlayerCommandDisconnectNotifySize)
	binary.LittleEndian.PutUint64(bytes[0:], s.Uid)
	bytes[8] = s.Value
	return bytes
}

func (s *S1ProtoPlayerCommandDisconnectNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoPlayerCommandDisconnectNotifySize {
		return InvalidDataError
	}

	s.Uid = binary.LittleEndian.Uint64(bytes[0:])
	s.Value = bytes[8]
	return nil
}
