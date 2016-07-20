package wrappclock

import "time"

var clock ClockProvider = &WrappClock{}

type WrappClock struct {}

func (*WrappClock) Now() time.Time {
	return time.Now()
}

func SetClock(clock ClockProvider) {
	clock = clock
}
