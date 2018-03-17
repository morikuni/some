package some

import (
	"math/rand"
)

// String returns a string generator.
func String(r *rand.Rand) *SomeString {
	return &SomeString{
		r,
		10,
	}
}

// SomeString is a string generator.
type SomeString struct {
	r   *rand.Rand
	len int
}

// Len sets the lentgh of a random string value.
func (g *SomeString) Len(n int) *SomeString {
	g.len = n
	return g
}

// Gen returns a string value.
func (g *SomeString) Gen() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, g.len)
	for i := range result {
		result[i] = chars[Int(g.r).Max(len(chars)).Gen()]
	}
	return string(result)
}

// GenP returns a string pointer.
func (g *SomeString) GenP() *string {
	s := g.Gen()
	return &s
}
