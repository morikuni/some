package some

import (
	"math/rand"
)

// Bool returns a bool generator.
var Bool BoolSpec = BoolSpec{}

type BoolSpec struct{}

func (s BoolSpec) Generate(r *rand.Rand) bool {
	return r.Int()&0x1 == 0
}

func (s *Some) Bool(key string, spec BoolSpec) bool {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(bool)
}
