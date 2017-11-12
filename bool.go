package dummy

import (
	"math/rand"
)

// Bool returns a bool generator.
func Bool(r *rand.Rand) *BoolGen {
	return &BoolGen{
		r,
	}
}

// BoolGen is a bool generator.
type BoolGen struct {
	r *rand.Rand
}

// Gen returns a random bool value.
func (g *BoolGen) Gen() bool {
	return g.r.Int()&0x1 == 0
}

// GenP returns a random bool pointer.
func (g *BoolGen) GenP() *bool {
	b := g.Gen()
	return &b
}
