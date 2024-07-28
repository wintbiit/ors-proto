package proto

type ESceneStateType int32

const (
	e_Scene_StateType_ess_null       ESceneStateType = 0
	e_Scene_StateType_ess_Wait       ESceneStateType = 1
	e_Scene_StateType_ess_Battle     ESceneStateType = 2
	e_Scene_StateType_ess_Settlement ESceneStateType = 3
	e_Scene_StateType_ess_End        ESceneStateType = 4
	e_Scene_StateType_ess_max        ESceneStateType = 5
)
