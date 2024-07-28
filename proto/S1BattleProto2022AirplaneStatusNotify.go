package proto

type S1BattleProto2022AirplaneStatusNotify struct {
	Fixshoottime  float32
	Cdrefreshcost int32
	Curbullet     int16
	Maxbullet     int16
	IsFire        byte
	Curshoottime  float32
	Cdleft        float32
	Countleft     uint16
	Isincd        byte
	Robotid       byte
	Energy        int32
}

const S1BattleProto2022AirplaneStatusNotifySize = 29

func (s *S1BattleProto2022AirplaneStatusNotify) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022AirplaneStatusNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022AirplaneStatusNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
