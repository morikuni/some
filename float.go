package some

import (
	"math/rand"
)

var Float FloatSpec = FloatSpec{
	1,
	0,
}

type FloatSpec struct {
	max float64
	min float64
}

// Max sets the maximum value of a random float value.
func (s FloatSpec) Max(f float64) FloatSpec {
	s.max = f
	return s
}

// Min sets the minimum value of a random float value.
func (s FloatSpec) Min(f float64) FloatSpec {
	s.min = f
	return s
}

func (s *FloatSpec) Generate(r *rand.Rand) float64 {
	diff := s.max - s.min
	return s.min + diff*r.Float64()
}

func (s *Some) Float(key string, spec FloatSpec) float64 {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(float64)
}
