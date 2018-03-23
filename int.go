package some

import (
	"math"
	"math/rand"
)

var Int IntSpec = IntSpec{
	Int64Spec{
		math.MaxInt32,
		0,
	},
}

// IntSpec is specification of a int value.
type IntSpec struct {
	Int64Spec
}

// Max sets the maximum value of a random int value.
func (s IntSpec) Max(n int) IntSpec {
	return IntSpec{s.Int64Spec.Max(int64(n))}
}

// Min sets the minimum value of a random int value.
func (s IntSpec) Min(n int) IntSpec {
	return IntSpec{s.Int64Spec.Min(int64(n))}
}

func (s IntSpec) Generate(r *rand.Rand) int {
	i := s.Int64Spec.Generate(r)
	return int(i)
}

func (s *Some) Int(key string, spec IntSpec) int {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(int)
}

var Int64 Int64Spec = Int64Spec{
	math.MaxInt64,
	0,
}

// IntSpec is specification of a int value.
type Int64Spec struct {
	max int64
	min int64
}

// Max sets the maximum value of a random int value.
func (s Int64Spec) Max(n int64) Int64Spec {
	s.max = n
	return s
}

// Min sets the minimum value of a random int value.
func (s Int64Spec) Min(n int64) Int64Spec {
	s.min = n
	return s
}

func (s Int64Spec) Generate(r *rand.Rand) int64 {
	diff := s.max - s.min
	return s.min + r.Int63n(diff)
}

func (s *Some) Int64(key string, spec Int64Spec) int64 {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(int64)
}
