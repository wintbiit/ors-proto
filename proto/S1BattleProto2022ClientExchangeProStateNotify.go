package proto

type S1BattleProto2022ClientExchangeProStateNotify struct {
	MineralRfidState byte
	IsOnline         byte
	ErrorCode        byte
	Z                float32
	Pitch            float32
	CoinsExchange    int32
	RobotId          byte
	ExchangeLevel    byte
	CoinsAddRate     float32
	EngineerLevel    byte
	Y                float32
	State            byte
	Ir2              byte
	X                float32
	Roll             float32
	Yaw              float32
	GoldCount        byte
	SilverCount      byte
	IrState          byte
	Ir1              byte
}

const S1BattleProto2022ClientExchangeProStateNotifySize = 44

func (s *S1BattleProto2022ClientExchangeProStateNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientExchangeProStateNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022ClientExchangeProStateNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
