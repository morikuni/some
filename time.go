package some

import (
	"math/rand"
	"time"
)

var Time TimeSpec = TimeSpec{
	time.Unix(0, 0),
	time.Date(3000, time.January, 1, 0, 0, 0, 0, time.UTC),
}

type TimeSpec struct {
	after  time.Time
	before time.Time
}

// After sets the minimum time of a random time.
func (s TimeSpec) After(t time.Time) TimeSpec {
	s.after = t
	return s
}

// Before sets the maximux time of a random time.
func (s TimeSpec) Before(t time.Time) TimeSpec {
	s.before = t
	return s
}

func (s TimeSpec) Generate(r *rand.Rand) time.Time {
	d := s.before.Sub(s.after)
	return s.after.Add(time.Duration(Int64.Max(int64(d)).Generate(r)))
}

func (s *Some) Time(key string, spec TimeSpec) time.Time {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(time.Time)
}
