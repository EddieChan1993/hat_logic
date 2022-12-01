package cst

type PlayerSignalTyp int8

const (
	SignalExitWithSaveDb PlayerSignalTyp = 1 //退出进程保存数据
	SignalOnlyExit       PlayerSignalTyp = 2 //仅仅退出进程
)
