package gigasecond

import "time"

const gigaSecond = 1000000000

// AddGigasecond add giga second(10^9) to t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gigaSecond)
}
