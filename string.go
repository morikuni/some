package dummy

import (
	"math/rand"
)

func String(r *rand.Rand) *StringGen {
	return &StringGen{
		r,
		10,
	}
}

type StringGen struct {
	r   *rand.Rand
	len int
}

func (g *StringGen) Len(n int) *StringGen {
	g.len = n
	return g
}

func (g *StringGen) Gen() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, g.len)
	for i := range result {
		result[i] = chars[Int(g.r).Max(len(chars)).Gen()]
	}
	return string(result)
}

func (g *StringGen) GenP() *string {
	s := g.Gen()
	return &s
}
