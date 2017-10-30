package dummy

import (
	"math/rand"
)

func Bool(r *rand.Rand) *BoolGen {
	return &BoolGen{
		r,
	}
}

type BoolGen struct {
	r *rand.Rand
}

func (g *BoolGen) Gen() bool {
	return g.r.Int()&0x1 == 0
}

func (g *BoolGen) GenP() *bool {
	b := g.Gen()
	return &b
}
