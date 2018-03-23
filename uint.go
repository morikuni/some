package some

import (
	"math"
	"math/rand"
)

var Uint UintSpec = UintSpec{
	math.MaxUint64,
	0,
}

type UintSpec struct {
	max uint
	min uint
}

// Max sets the maximum value of a random uint value with uint64.
func (s UintSpec) Max(n uint) UintSpec {
	s.max = n
	return s
}

// Min sets the minimum value of a random uint value with uint64.
func (s UintSpec) Min(n uint) UintSpec {
	s.min = n
	return s
}

func (s *Some) Uint(key string, spec UintSpec) uint {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(uint)
}
