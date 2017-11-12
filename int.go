package dummy

import (
	"math"
	"math/rand"
	"strconv"
)

// Int returns a int generator.
func Int(r *rand.Rand) *IntGen {
	return &IntGen{
		r,
		math.MaxInt64,
		0,
	}
}

// IntGen is a int generator.
type IntGen struct {
	r   *rand.Rand
	max int64
	min int64
}

// Max sets the maximum value of a random int value.
func (g *IntGen) Max(n int) *IntGen {
	return g.Max64(int64(n))
}

// Min sets the minimum value of a random int value.
func (g *IntGen) Min(n int) *IntGen {
	return g.Min64(int64(n))
}

// Max sets the maximum value of a random int value with int64.
func (g *IntGen) Max64(n int64) *IntGen {
	g.max = n
	return g
}

// Min sets the minimum value of a random int value with int64.
func (g *IntGen) Min64(n int64) *IntGen {
	g.min = n
	return g
}

// Gen  returns a int value.
func (g *IntGen) Gen() int {
	return int(g.Gen64())
}

// GenP returns a int pointer.
func (g *IntGen) GenP() *int {
	i := g.Gen()
	return &i
}

// Gen64 returns a int64 value.
func (g *IntGen) Gen64() int64 {
	diff := g.max - g.min
	return g.min + g.r.Int63n(diff)
}

// Gen64P returns a int64 pointer.
func (g *IntGen) Gen64P() *int64 {
	i := g.Gen64()
	return &i
}

// GenString returns a int64 value as a string value.
func (g *IntGen) GenString() string {
	return strconv.FormatInt(g.Gen64(), 10)
}

// GenString returns a int64 value as a string pointer.
func (g *IntGen) GenStringP() *string {
	s := g.GenString()
	return &s
}
