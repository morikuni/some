package dummy

import (
	"math/rand"
)

// Float returns float generator.
func Float(r *rand.Rand) *FloatGen {
	return &FloatGen{
		r,
		1,
		0,
	}
}

// FloatGen is a float generator.
type FloatGen struct {
	r   *rand.Rand
	max float64
	min float64
}

// Max sets the maximum value of a random float value.
func (g *FloatGen) Max(f float64) *FloatGen {
	g.max = f
	return g
}

// Min sets the minimum value of a random float value.
func (g *FloatGen) Min(f float64) *FloatGen {
	g.min = f
	return g
}

// Gen64 returns a float64 value.
func (g *FloatGen) Gen64() float64 {
	diff := g.max - g.min
	return g.min + diff*g.r.Float64()
}

// Gen64P returns a float64 pointer.
func (g *FloatGen) Gen64P() *float64 {
	f := g.Gen64()
	return &f
}

// Gen32 returns a float32 value.
func (g *FloatGen) Gen32() float32 {
	return float32(g.Gen64())
}

// Gen32P returns a float32 pointer.
func (g *FloatGen) Gen32P() *float32 {
	f := g.Gen32()
	return &f
}
