package some

import (
	"math"
	"math/rand"
)

// AnyUinwt is a default uint spec.
var AnyUint UintSpec = UintSpec{
	math.MaxUint32,
	0,
}

// UintSpec is a spec of a uint.
type UintSpec struct {
	Max uint
	Min uint
}

// Uint generates a uint according to a key and spec.
func (g *Generator) Uint(key string, spec UintSpec) uint {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(uint)
}
