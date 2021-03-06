package some

import (
	"math/rand"
)

// String is a default string spec.
var String StringSpec = StringSpec{
	10,
}

// BoolSpec is a spec of a string.
type StringSpec struct {
	Len int
}

// WithLen returns new spec with given len.
func (s StringSpec) WithLen(len int) StringSpec {
	s.Len = len
	return s
}

// Generate generates a random string from r.
func (s StringSpec) Generate(r *rand.Rand) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, s.Len)
	for i := range result {
		result[i] = chars[IntSpec{Max: len(chars)}.Generate(r)]
	}
	return string(result)
}

// String generates a string according to a key and spec.
func (g *Generator) String(key string, spec StringSpec) string {
	return g.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(string)
}
