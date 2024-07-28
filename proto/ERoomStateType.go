package proto

type ERoomStateType int32

const (
	e_Room_StateType_ers_Wait       ERoomStateType = 0
	e_Room_StateType_ers_CountDown1 ERoomStateType = 1
	e_Room_StateType_ers_CountDown2 ERoomStateType = 2
	e_Room_StateType_ers_CountDown3 ERoomStateType = 3
	e_Room_StateType_ers_Playing    ERoomStateType = 4
	e_Room_StateType_ers_max        ERoomStateType = 5
)
