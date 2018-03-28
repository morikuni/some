package some

import (
	"math/rand"
)

// AnyBool is a default bool spec.
var AnyBool BoolSpec = BoolSpec{}

// BoolSpec is a spec of a bool.
type BoolSpec struct{}

// Generate generates a random bool from r.
func (s BoolSpec) Generate(r *rand.Rand) bool {
	return r.Int()&0x1 == 0
}

// Bool generates a bool according to a key and spec.
func (g *Generator) Bool(key string, spec BoolSpec) bool {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(bool)
}
