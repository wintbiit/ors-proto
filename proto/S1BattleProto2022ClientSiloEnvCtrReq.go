package proto

type S1BattleProto2022ClientSiloEnvCtrReq struct {
	Teamid      byte
	ControlCode byte
	Target      byte
}

const S1BattleProto2022ClientSiloEnvCtrReqSize = 3

func (s *S1BattleProto2022ClientSiloEnvCtrReq) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientSiloEnvCtrReqSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientSiloEnvCtrReq) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
