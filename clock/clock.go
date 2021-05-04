package clock

import "fmt"

// Clock API:
//
// type Clock
// New(hour, minute int) Clock
// (Clock) String() string
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock

// Clock type
type Clock struct {
	m int // minute
}

// New is a clock "constructor"
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}

	return Clock{m}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.m/60, clock.m%60)
}

// Add will add minutes to clock
func (clock Clock) Add(minutes int) Clock {
	return New(0, clock.m+minutes)
}

// Subtract will subtract minutes to clock
func (clock Clock) Subtract(minutes int) Clock {
	return New(0, clock.m-minutes)
}
