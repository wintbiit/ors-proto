package proto

/*
public enum ELoginAckResultType
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

type ELoginAckResultType int32

const (
	e_LoginAck_ResultType_elrt_null    ELoginAckResultType = 0
	e_LoginAck_ResultType_elrt_成功      ELoginAckResultType = 1
	e_LoginAck_ResultType_elrt_无效版本    ELoginAckResultType = 2
	e_LoginAck_ResultType_elrt_账号或密码错误 ELoginAckResultType = 3
	e_LoginAck_ResultType_elrt_账号长度不符  ELoginAckResultType = 4
	e_LoginAck_ResultType_elrt_已经登录    ELoginAckResultType = 5
	e_LoginAck_ResultType_elrt_不在前台    ELoginAckResultType = 6
	e_LoginAck_ResultType_elrt_离开前台失败  ELoginAckResultType = 7
	e_LoginAck_ResultType_elrt_之前并未在线  ELoginAckResultType = 8
	e_LoginAck_ResultType_elrt_之前尚未登录  ELoginAckResultType = 9
	e_LoginAck_ResultType_elrt_操作太频繁   ELoginAckResultType = 10
	e_LoginAck_ResultType_elrt_未知错误    ELoginAckResultType = 11
	e_LoginAck_ResultType_elrt_max     ELoginAckResultType = 12
)
