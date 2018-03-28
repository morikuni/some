package some

import (
	"math/rand"
)

var (
	// Float64 is a default float spec.
	Float64 Float64Spec = Float64Spec{
		1,
		0,
	}
	// Float32 is a default float spec.
	Float32 Float32Spec = Float32Spec{
		1,
		0,
	}
)

// Float64Spec is a spec of a float64.
type Float64Spec struct {
	Max float64
	Min float64
}

// WithMax returns new spec with given max.
func (s Float64Spec) WithMax(max float64) Float64Spec {
	s.Max = max
	return s
}

// WithMin returns new spec with given min.
func (s Float64Spec) WithMin(min float64) Float64Spec {
	s.Min = min
	return s
}

// Generate generates a random bool from r.
func (s Float64Spec) Generate(r *rand.Rand) float64 {
	diff := s.Max - s.Min
	return s.Min + diff*r.Float64()
}

// Float64 generates a float according to a key and spec.
func (g *Generator) Float64(key string, spec Float64Spec) float64 {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(float64)
}

// Float32Spec is a spec of a float64.
type Float32Spec struct {
	Max float32
	Min float32
}

// WithMax returns new spec with given max.
func (s Float32Spec) WithMax(max float32) Float32Spec {
	s.Max = max
	return s
}

// WithMin returns new spec with given min.
func (s Float32Spec) WithMin(min float32) Float32Spec {
	s.Min = min
	return s
}

// Generate generates a random bool from r.
func (s *Float32Spec) Generate(r *rand.Rand) float32 {
	diff := s.Max - s.Min
	return s.Min + diff*r.Float32()
}

// Float32 generates a float according to a key and spec.
func (g *Generator) Float32(key string, spec Float32Spec) float32 {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(float32)
}
