package dummy

import (
	"math/rand"
	"time"
)

// Time returns a time generator.
func Time(r *rand.Rand) *TimeGen {
	return &TimeGen{
		r,
		time.Unix(0, 0),
		time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

// TimeGen is a time generator.
type TimeGen struct {
	r      *rand.Rand
	after  time.Time
	before time.Time
}

// After sets the minimum time of a random time.
func (g *TimeGen) After(t time.Time) *TimeGen {
	g.after = t
	return g
}

// Before sets the maximux time of a random time.
func (g *TimeGen) Before(t time.Time) *TimeGen {
	g.before = t
	return g
}

// Gen returns a time value.
func (g *TimeGen) Gen() time.Time {
	d := g.before.Sub(g.after)
	return g.after.Add(time.Duration(Int(g.r).Max64(int64(d)).Gen64()))
}

// GenP returns a time pointer.
func (g *TimeGen) GenP() *time.Time {
	t := g.Gen()
	return &t
}
