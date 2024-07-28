package proto

/*
public enum e_Scene_StateType
{
  ess_null,
  ess_Wait,
  ess_Battle,
  ess_Settlement,
  ess_End,
  ess_max,
}

*/

type e_Scene_StateType int32

const (
	e_Scene_StateType_ess_null       e_Scene_StateType = 0
	e_Scene_StateType_ess_Wait       e_Scene_StateType = 1
	e_Scene_StateType_ess_Battle     e_Scene_StateType = 2
	e_Scene_StateType_ess_Settlement e_Scene_StateType = 3
	e_Scene_StateType_ess_End        e_Scene_StateType = 4
	e_Scene_StateType_ess_max        e_Scene_StateType = 5
)
