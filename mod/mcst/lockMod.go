package mcst

/**
模块解锁
*/
import config "hat_logic/pbgo/pbcfg"

type LockTyp = DefaultInt

const (
	LockArtifactId LockTyp = 2 //神器
	LockBuildId    LockTyp = 3 //建筑
)

//IsLockMod 是否解锁模块
func IsLockMod(lockId LockTyp, leaderLv, maxLevel DefaultInt) bool {
	FunctionOpenData := config.FunctionOpenData.Get(lockId)
	if leaderLv >= FunctionOpenData.Level &&
		maxLevel >= FunctionOpenData.Stage {
		//领主等级和开启关卡满足条件才解锁
		return true
	}
	return false
}
