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
func (s *SomeString) Len(n int) *SomeString {
	s.len = n
	return s
}

// Gen returns a string value.
func (s *SomeString) Gen() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, s.len)
	for i := range result {
		result[i] = chars[Int(s.r).Max(len(chars)).Gen()]
	}
	return string(result)
}

// GenP returns a string pointer.
func (s *SomeString) GenP() *string {
	str:= s.Gen()
	return &str
}
