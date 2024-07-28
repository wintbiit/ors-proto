package proto

/*
public enum e_LeaveRoomAck_ResultType
{
  elrt_null,
  elrt_Success,
  elrt_UnKnown,
  elrt_max,
}

*/

type e_LeaveRoomAck_ResultType int32

const (
	e_LeaveRoomAck_ResultType_elrt_null    e_LeaveRoomAck_ResultType = 0
	e_LeaveRoomAck_ResultType_elrt_Success e_LeaveRoomAck_ResultType = 1
	e_LeaveRoomAck_ResultType_elrt_UnKnown e_LeaveRoomAck_ResultType = 2
	e_LeaveRoomAck_ResultType_elrt_max     e_LeaveRoomAck_ResultType = 3
)
