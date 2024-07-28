package proto

/*
public enum e_EnterRoomAck_ResultType
{
  eert_null,
  eert_Success,
  eert_UnKnown,
  eert_max,
}

*/

type e_EnterRoomAck_ResultType int32

const (
	e_EnterRoomAck_ResultType_eert_null    e_EnterRoomAck_ResultType = 0
	e_EnterRoomAck_ResultType_eert_Success e_EnterRoomAck_ResultType = 1
	e_EnterRoomAck_ResultType_eert_UnKnown e_EnterRoomAck_ResultType = 2
	e_EnterRoomAck_ResultType_eert_max     e_EnterRoomAck_ResultType = 3
)
