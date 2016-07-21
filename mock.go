package wrappclock

import "time"

// Mock struct for testClock
type testClock struct {
	// The actual, real-world time at which the mock clock was set
	origTimeSet time.Time

	// The time at which the mock clock was set
	timeInstance time.Time

	// Whether the clock should run, keeping pace with the real world clock, or stand still
	running bool

	// Optional offset by which time is set forwards or backwards
	offset time.Duration
}

func (tc *testClock) Now() time.Time {

	// Do we want this time to keep updating
	if tc.running == true {
		// Get the duration of time elapsed since we first set the mock clock
		timeDelta := time.Now().Sub(tc.origTimeSet)

		//Add the time elapsed to the mock clock
		return tc.timeInstance.Add(timeDelta, tc.offset)
	}
	return tc.timeInstance.Add(tc.offset)
}

