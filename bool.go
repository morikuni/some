package some

import (
	"math/rand"
)

// Bool returns a bool generator.
func Bool(r *rand.Rand) *SomeBool {
	return &SomeBool{
		r,
	}
}

// SomeBool is a bool generator.
type SomeBool struct {
	r *rand.Rand
}

// Gen returns a random bool value.
func (g *SomeBool) Gen() bool {
	return g.r.Int()&0x1 == 0
}

// GenP returns a random bool pointer.
func (g *SomeBool) GenP() *bool {
	b := g.Gen()
	return &b
}
