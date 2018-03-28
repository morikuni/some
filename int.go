package some

import (
	"math"
	"math/rand"
)

var (
	// AnyInt is a default int spec.
	AnyInt IntSpec = IntSpec{
		math.MaxInt32,
		0,
	}
	// AnyInt64 is a default int64 spec.
	AnyInt64 Int64Spec = Int64Spec{
		math.MaxInt64,
		0,
	}
)

// IntSpec is spec of a int.
type IntSpec struct {
	Max int
	Min int
}

// Generate generates a random int from r.
func (s IntSpec) Generate(r *rand.Rand) int {
	diff := s.Max - s.Min
	return s.Min + r.Intn(diff)
}

// Int generates a int according to a key and spec.
func (g *Generator) Int(key string, spec IntSpec) int {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(int)
}

// Int64Spec is spec of a int64.
type Int64Spec struct {
	Max int64
	Min int64
}

// Generate generates a random int64 from r.
func (s Int64Spec) Generate(r *rand.Rand) int64 {
	diff := s.Max - s.Min
	return s.Min + r.Int63n(diff)
}

// Int64 generates a int64 according to a key and spec.
func (g *Generator) Int64(key string, spec Int64Spec) int64 {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(int64)
}
