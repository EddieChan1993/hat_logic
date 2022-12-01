package register

import (
	"time"
)

type DefaultScheduler struct {
	Interval time.Duration //间隔执行时间
}

func (s *DefaultScheduler) Next(prev time.Time) time.Time {
	return prev.Add(s.Interval)
}
