package wrappclock

import "time"

type ClockProvider interface {
	Now() time.Time
}
