package some

import (
	"math"
	"math/rand"
	"strconv"
)

// Uint returns a uint generator.
func Uint(r *rand.Rand) *SomeUint {
	return &SomeUint{
		r,
		math.MaxUint64,
		0,
	}
}

// SomeUint is a uint generator.
type SomeUint struct {
	r   *rand.Rand
	max uint64
	min uint64
}

// Max sets the maximum value of a random uint value.
func (g *SomeUint) Max(n uint) *SomeUint {
	return g.Max64(uint64(n))
}

// Min sets the minimum value of a random uint value.
func (g *SomeUint) Min(n uint) *SomeUint {
	return g.Min64(uint64(n))
}

// Max64 sets the maximum value of a random uint value with uint64.
func (g *SomeUint) Max64(n uint64) *SomeUint {
	g.max = n
	return g
}

// Min64 sets the minimum value of a random uint value with uint64.
func (g *SomeUint) Min64(n uint64) *SomeUint {
	g.min = n
	return g
}

// Gen returns a uint value.
func (g *SomeUint) Gen() uint {
	return uint(g.Gen64())
}

// GenP returns a uint pointer.
func (g *SomeUint) GenP() *uint {
	u := g.Gen()
	return &u
}

// Gen64P returns a uint64 pointer.
func (g *SomeUint) Gen64P() *uint64 {
	u := g.Gen64()
	return &u
}

// GenString returns a uint64 value as a string value.
func (g *SomeUint) GenString() string {
	return strconv.FormatUint(g.Gen64(), 10)
}

// GenStringP returns a uint64 value as a string value.
func (g *SomeUint) GenStringP() *string {
	s := g.GenString()
	return &s
}
