package dummy

import (
	"math"
	"math/rand"
	"strconv"
)

func Int(r *rand.Rand) *IntGen {
	return &IntGen{
		r,
		math.MaxInt64,
		0,
	}
}

type IntGen struct {
	r   *rand.Rand
	max int64
	min int64
}

func (g *IntGen) Max(n int) *IntGen {
	g.max = int64(n)
	return g
}

func (g *IntGen) Min(n int) *IntGen {
	g.min = int64(n)
	return g
}

func (g *IntGen) Max64(n int64) *IntGen {
	g.max = n
	return g
}

func (g *IntGen) Min64(n int64) *IntGen {
	g.min = n
	return g
}

func (g *IntGen) Gen() int {
	return int(g.Gen64())
}

func (g *IntGen) GenP() *int {
	i := g.Gen()
	return &i
}

func (g *IntGen) Gen64() int64 {
	diff := g.max - g.min
	return g.min + g.r.Int63n(diff)
}

func (g *IntGen) Gen64P() *int64 {
	i := g.Gen64()
	return &i
}

func (g *IntGen) GenString() string {
	return strconv.FormatInt(g.Gen64(), 10)
}

func (g *IntGen) GenStringP() *string {
	s := g.GenString()
	return &s
}
