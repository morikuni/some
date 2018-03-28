package some

import (
	"math/rand"
	"time"
)

// AnyBool is a default time.Time spec.
var AnyTime TimeSpec = TimeSpec{
	time.Unix(0, 0),
	time.Date(3000, time.January, 1, 0, 0, 0, 0, time.UTC),
}

// TimeSpec is a spec of a time.Time.
type TimeSpec struct {
	After  time.Time
	Before time.Time
}

// Generate generates a random time.Time from r.
func (s TimeSpec) Generate(r *rand.Rand) time.Time {
	d := s.Before.Sub(s.After)
	return s.After.Add(time.Duration(Int64Spec{Max: int64(d)}.Generate(r)))
}

// Time generates a time.Time according to a key and spec.
func (g *Generator) Time(key string, spec TimeSpec) time.Time {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(time.Time)
}
