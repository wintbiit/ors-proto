package proto

type S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify struct {
	TeamId      byte
	DoorState   byte
	DoorOpenCnt byte
}

const S1BattleProto2022ClientSiloEnvDoorOpenCloseNotifySize = 3

func (s *S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientSiloEnvDoorOpenCloseNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientSiloEnvDoorOpenCloseNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
