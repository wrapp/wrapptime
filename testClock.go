package testClock

import "time"

type TestClock struct {
	timeInstance time.Time
	time
}

func SetTime(passedTime time.Time) {
	return &TestClock{timeInstance: passedTime}
}

func (tc *TestClock) Now() time.Time {
	if tc == nil {
		return time.Now()
	} else {
		return tc.timeInstance
	}

}

func Now() time.Time {
	return time.Now()
}
