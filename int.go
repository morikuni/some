package some

import (
	"math"
	"math/rand"
	"strconv"
)

// Int returns a int generator.
func Int(r *rand.Rand) *SomeInt {
	return &SomeInt{
		r,
		math.MaxInt64,
		0,
	}
}

// SomeInt is a int generator.
type SomeInt struct {
	r   *rand.Rand
	max int64
	min int64
}

// Max sets the maximum value of a random int value.
func (g *SomeInt) Max(n int) *SomeInt {
	return g.Max64(int64(n))
}

// Min sets the minimum value of a random int value.
func (g *SomeInt) Min(n int) *SomeInt {
	return g.Min64(int64(n))
}

// Max64 sets the maximum value of a random int value with int64.
func (g *SomeInt) Max64(n int64) *SomeInt {
	g.max = n
	return g
}

// Min64 sets the minimum value of a random int value with int64.
func (g *SomeInt) Min64(n int64) *SomeInt {
	g.min = n
	return g
}

// Gen  returns a int value.
func (g *SomeInt) Gen() int {
	return int(g.Gen64())
}

// GenP returns a int pointer.
func (g *SomeInt) GenP() *int {
	i := g.Gen()
	return &i
}

// Gen64 returns a int64 value.
func (g *SomeInt) Gen64() int64 {
	diff := g.max - g.min
	return g.min + g.r.Int63n(diff)
}

// Gen64P returns a int64 pointer.
func (g *SomeInt) Gen64P() *int64 {
	i := g.Gen64()
	return &i
}

// GenString returns a int64 value as a string value.
func (g *SomeInt) GenString() string {
	return strconv.FormatInt(g.Gen64(), 10)
}

// GenStringP returns a int64 value as a string pointer.
func (g *SomeInt) GenStringP() *string {
	s := g.GenString()
	return &s
}
