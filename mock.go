package wrappclock

import "time"

// Mock struct for testClock
type testClock struct {
	timeInstance time.Time
}

func (tc *testClock) Now() time.Time {
	return tc.timeInstance
}
