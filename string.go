package dummy

import (
	"math/rand"
)

// String returns a string generator.
func String(r *rand.Rand) *StringGen {
	return &StringGen{
		r,
		10,
	}
}

// StringGen is a string generator.
type StringGen struct {
	r   *rand.Rand
	len int
}

// Len sets the lentgh of a random string value.
func (g *StringGen) Len(n int) *StringGen {
	g.len = n
	return g
}

// Gen returns a string value.
func (g *StringGen) Gen() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, g.len)
	for i := range result {
		result[i] = chars[Int(g.r).Max(len(chars)).Gen()]
	}
	return string(result)
}

// GenP returns a string pointer.
func (g *StringGen) GenP() *string {
	s := g.Gen()
	return &s
}
