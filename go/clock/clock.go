package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	for m < 0 {
		m = 60 + m
		h -= 1
	}
	for h < 0 {
		h = 24 + h
	}
	for m >= 60 {
		m -= 60
		h += 1
	}
	for h >= 24 {
		h -= 24
	}
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	c.m += m
	for c.m > 59 {
		c.m -= 60
		c.h += 1
		if c.h == 24 {
			c.h = 0
		}
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m -= m
	for c.m < 0 {
		c.m += 60
		c.h -= 1
		if c.h == -1 {
			c.h = 23
		}
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
