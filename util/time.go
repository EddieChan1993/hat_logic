package util

import "time"

const Milli int64 = 1000    //秒和毫秒的倍数关系,必须要给个确切类型值，因为通常和int64相乘，所以设定为int64，否则hu
const DaySecs int64 = 86400 //一天的秒数

//TimeNowUnix 当前秒
func TimeNowUnix() int64 {
	return time.Now().Unix()
}

//TimeNowUnixMilli 当前毫秒
func TimeNowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

// TimeNowUnixNano 当前纳秒
func TimeNowUnixNano() int64 {
	return time.Now().UnixNano()
}

//Week 第几周
func Week() int32 {
	year, week := time.Now().ISOWeek()
	return int32(week*10000 + year)
}

//ZeroToday 当天0点（秒）
func ZeroToday() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

//ZeroOtherDay 其他天（秒）
func ZeroOtherDay(day int) int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()+day, 0, 0, 0, 0, t.Location()).Unix()
}

//DayNow 当前位于开始时间第几天
//start 开始时间毫秒
func DayNow(start int64) int32 {
	now := TimeNowUnixMilli()
	if start <= 0 || start > now {
		return 0
	}
	t := time.UnixMilli(start)
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixMilli()
	return int32((now-startDay)/(DaySecs*Milli) + 1)
}

//NextDayCd 距离跨天剩余时间
func NextDayCd() int32 {
	zeroOtherDay := ZeroOtherDay(1) * Milli
	cd := zeroOtherDay - TimeNowUnixMilli()
	return int32(cd)
}
