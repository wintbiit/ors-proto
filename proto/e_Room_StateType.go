package proto

/*
public enum e_Room_StateType
{
  ers_Wait,
  ers_CountDown1,
  ers_CountDown2,
  ers_CountDown3,
  ers_Playing,
  ers_max,
}

*/

type e_Room_StateType int32

const (
	e_Room_StateType_ers_Wait       e_Room_StateType = 0
	e_Room_StateType_ers_CountDown1 e_Room_StateType = 1
	e_Room_StateType_ers_CountDown2 e_Room_StateType = 2
	e_Room_StateType_ers_CountDown3 e_Room_StateType = 3
	e_Room_StateType_ers_Playing    e_Room_StateType = 4
	e_Room_StateType_ers_max        e_Room_StateType = 5
)
