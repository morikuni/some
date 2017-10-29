package gen

import (
	"math/rand"
)

func Float(r *rand.Rand) *FloatGen {
	return &FloatGen{
		r,
		1,
		0,
	}
}

type FloatGen struct {
	r   *rand.Rand
	max float64
	min float64
}

func (g *FloatGen) Max(f float64) *FloatGen {
	g.max = f
	return g
}

func (g *FloatGen) Min(f float64) *FloatGen {
	g.min = f
	return g
}

func (g *FloatGen) Gen() float64 {
	diff := g.max - g.min
	return g.min + diff*g.r.Float64()
}

func (g *FloatGen) Gen32() float32 {
	return float32(g.Gen())
}
