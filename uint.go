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
	g.max = uint64(n)
	return g
}

func (g *UintGen) Min(n uint) *UintGen {
	g.min = uint64(n)
	return g
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

func (g *UintGen) Gen64() uint64 {
	diff := g.max - g.min
	// TODO: use Uint64n if it's supported by math/rand.
	return g.min + (g.r.Uint64() % diff)
}

func (g *UintGen) GenString() string {
	return strconv.FormatUint(g.Gen64(), 10)
}
