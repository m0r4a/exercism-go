// package gigasecond allows you to determine the date
// and time one gigasecond after a certain date.
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond receives a time.Time value and returns a new time.Time
// representing the date and time after adding one gigasecond (1,000,000,000 seconds) to the input.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * (1e9))
}
