package some

import (
	"math/rand"
)

var (
	// AnyFloat64 is a default float spec.
	AnyFloat64 Float64Spec = Float64Spec{
		1,
		0,
	}
	// AnyFloat32 is a default float spec.
	AnyFloat32 Float32Spec = Float32Spec{
		1,
		0,
	}
)

// Float64Spec is a spec of a float64.
type Float64Spec struct {
	Max float64
	Min float64
}

// Generate generates a random bool from r.
func (s *Float64Spec) Generate(r *rand.Rand) float64 {
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

// Generate generates a random bool from r.
func (s *Float32Spec) Generate(r *rand.Rand) float32 {
	diff := s.Max - s.Min
	return s.Min + diff*r.Float32()
}

// Float32 generates a float according to a key and spec.
func (g *Generator) Float32(key string, spec Float32Spec) float32 {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(float32)
}
