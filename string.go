package some

import (
	"math/rand"
)

var String StringSpec = StringSpec{
	10,
}

// StringSpec is a specification of a string value.
type StringSpec struct {
	len int
}

// Len sets the length of a random string value.
func (s StringSpec) Len(n int) StringSpec {
	s.len = n
	return s
}

func (s StringSpec) Generate(r *rand.Rand) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, s.len)
	for i := range result {
		result[i] = chars[Int.Max(len(chars)).Generate(r)]
	}
	return string(result)
}

func (s *Some) String(key string, spec StringSpec) string {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(string)
}
