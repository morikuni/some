package dummy

import (
	"math"
	"math/rand"
	"strconv"
)

func Uint(r *rand.Rand) *UintGen {
	return &UintGen{
		r,
		math.MaxUint64,
		0,
	}
}

type UintGen struct {
	r   *rand.Rand
	max uint64
	min uint64
}

func (g *UintGen) Max(n uint) *UintGen {
	return g.Max64(uint64(n))
}

func (g *UintGen) Min(n uint) *UintGen {
	return g.Min64(uint64(n))
}

func (g *UintGen) Max64(n uint64) *UintGen {
	g.max = n
	return g
}

func (g *UintGen) Min64(n uint64) *UintGen {
	g.min = n
	return g
}

func (g *UintGen) Gen() uint {
	return uint(g.Gen64())
}

func (g *UintGen) GenP() *uint {
	u := g.Gen()
	return &u
}

func (g *UintGen) Gen64P() *uint64 {
	u := g.Gen64()
	return &u
}

func (g *UintGen) GenString() string {
	return strconv.FormatUint(g.Gen64(), 10)
}

func (g *UintGen) GenStringP() *string {
	s := g.GenString()
	return &s
}
