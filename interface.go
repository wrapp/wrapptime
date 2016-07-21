package wrapptime
import "time"

type ClockProvider interface {
	Now() time.Time
}
