package proto

const S1BattleProto2022ClientAllClientStatusSyncSize = 15

func (s *S1BattleProto2022ClientAllClientStatusSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientAllClientStatusSyncSize)
	bytes[0] = s.Listlength
	copy(bytes[1:], s.Status)
	return bytes
}

func (s *S1BattleProto2022ClientAllClientStatusSync) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022ClientAllClientStatusSyncSize {
		return InvalidDataError
	}

	s.Listlength = bytes[0]
	s.Status = bytes[1 : s.Listlength+1]
	return nil
}
