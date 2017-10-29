package gen

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

func (s *StringGen) Len(n int) *StringGen {
	s.len = n
	return s
}

func (s *StringGen) Gen() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, s.len)
	for i := range result {
		result[i] = chars[Int(s.r).Max(len(chars)).Gen()]
	}
	return string(result)
}
