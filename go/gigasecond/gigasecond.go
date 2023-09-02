package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000_000_000)
	return t
}
