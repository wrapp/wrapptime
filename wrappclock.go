package wrappclock

import "time"

var clock ClockProvider = &WrappClock{}

type WrappClock struct {}

func (*WrappClock) Now() time.Time {
	return time.Now()
}

func SetClock(t time.Time) {
	clock = &testClock{t}
}

func Now() time.Time  {
	return clock.Now()
}
