package some

import (
	"math/rand"
	"time"
)

// Time returns a time generator.
func Time(r *rand.Rand) *SomeTime {
	return &SomeTime{
		r,
		time.Unix(0, 0),
		time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

// SomeTime is a time generator.
type SomeTime struct {
	r      *rand.Rand
	after  time.Time
	before time.Time
}

// After sets the minimum time of a random time.
func (s *SomeTime) After(t time.Time) *SomeTime {
	s.after = t
	return s
}

// Before sets the maximux time of a random time.
func (s *SomeTime) Before(t time.Time) *SomeTime {
	s.before = t
	return s
}

// Gen returns a time value.
func (s *SomeTime) Gen() time.Time {
	d := s.before.Sub(s.after)
	return s.after.Add(time.Duration(Int(s.r).Max64(int64(d)).Gen64()))
}

// GenP returns a time pointer.
func (s *SomeTime) GenP() *time.Time {
	t := s.Gen()
	return &t
}
