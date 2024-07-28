package proto

type S1BattleProto2022ClientSupplySync struct {
	State               int32
	MakerRobotid        int32
	RfidRobotIdListLen  int32
	Connect             int32
	UsedBullets         int32
	UsableBullets       int32
	NextaddBullets      int32
	ErrorCodeArrListLen int32
	RfidRobotIdList     []int32
	MakeBullets         int32
	ErrorCodeArrList    []int32
	SupplyId            int32
	ReadyBox            int32
	FreedBox            int32
	Timespan            float32
}

const S1BattleProto2022ClientSupplySyncSize = 212

func (s *S1BattleProto2022ClientSupplySync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientSupplySyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientSupplySync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
