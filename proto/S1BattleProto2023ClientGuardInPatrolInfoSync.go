package proto

type S1BattleProto2023ClientGuardInPatrolInfoSync struct {
	RobotId                  byte
	IsCurrentInPatrolArea    byte
	IsLeavePatrolAreaTimeout byte
	LeavePatrolAreaLeftTime  float32
	IsCountDownActive        byte
}

const S1BattleProto2023ClientGuardInPatrolInfoSyncSize = 8

func (s *S1BattleProto2023ClientGuardInPatrolInfoSync) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientGuardInPatrolInfoSyncSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2023ClientGuardInPatrolInfoSync) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
