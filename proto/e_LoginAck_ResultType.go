package proto

/*
public enum e_LoginAck_ResultType
{
  elrt_null,
  elrt_成功,
  elrt_无效版本,
  elrt_账号或密码错误,
  elrt_账号长度不符,
  elrt_已经登录,
  elrt_不在前台,
  elrt_离开前台失败,
  elrt_之前并未在线,
  elrt_之前尚未登录,
  elrt_操作太频繁,
  elrt_未知错误,
  elrt_max,
}

*/

type e_LoginAck_ResultType int32

const (
	e_LoginAck_ResultType_elrt_null    e_LoginAck_ResultType = 0
	e_LoginAck_ResultType_elrt_成功      e_LoginAck_ResultType = 1
	e_LoginAck_ResultType_elrt_无效版本    e_LoginAck_ResultType = 2
	e_LoginAck_ResultType_elrt_账号或密码错误 e_LoginAck_ResultType = 3
	e_LoginAck_ResultType_elrt_账号长度不符  e_LoginAck_ResultType = 4
	e_LoginAck_ResultType_elrt_已经登录    e_LoginAck_ResultType = 5
	e_LoginAck_ResultType_elrt_不在前台    e_LoginAck_ResultType = 6
	e_LoginAck_ResultType_elrt_离开前台失败  e_LoginAck_ResultType = 7
	e_LoginAck_ResultType_elrt_之前并未在线  e_LoginAck_ResultType = 8
	e_LoginAck_ResultType_elrt_之前尚未登录  e_LoginAck_ResultType = 9
	e_LoginAck_ResultType_elrt_操作太频繁   e_LoginAck_ResultType = 10
	e_LoginAck_ResultType_elrt_未知错误    e_LoginAck_ResultType = 11
	e_LoginAck_ResultType_elrt_max     e_LoginAck_ResultType = 12
)
