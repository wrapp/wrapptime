package wrapptime

import "time"

var clock ClockProvider = &WrappClock{}

type WrappClock struct {}

func (*WrappClock) Now() time.Time {
	return time.Now()
}

func Set(t time.Time, running bool) {
	clock = &testClock{timeInstance: t, running:running, origTimeSet:time.Now()}
}

func Now() time.Time  {
	return clock.Now()
}
