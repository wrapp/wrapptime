package wrappclock

import "time"

var clock ClockProvider = &WrappClock{}

type WrappClock struct {}

func (*WrappClock) Now() time.Time {
	return time.Now()
}

func Set(t time.Time, running bool) {
	clock = &testClock{timeInstance: t, running:running, origTimeSet:time.Now()}
}

func Travel(d time.Duration) {
	if &testClock.timeInstance.IsZero() != true {
		&testClock{offset:d}
	}
	return
}

func Now() time.Time  {
	return clock.Now()
}
