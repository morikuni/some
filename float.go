package some

import (
	"math/rand"
)

// Float returns float generator.
func Float(r *rand.Rand) *SomeFloat {
	return &SomeFloat{
		r,
		1,
		0,
	}
}

// SomeFloat is a float generator.
type SomeFloat struct {
	r   *rand.Rand
	max float64
	min float64
}

// Max sets the maximum value of a random float value.
func (s *SomeFloat) Max(f float64) *SomeFloat {
	s.max = f
	return s
}

// Min sets the minimum value of a random float value.
func (s *SomeFloat) Min(f float64) *SomeFloat {
	s.min = f
	return s
}

// Gen64 returns a float64 value.
func (s *SomeFloat) Gen64() float64 {
	diff := s.max - s.min
	return s.min + diff*s.r.Float64()
}

// Gen64P returns a float64 pointer.
func (s *SomeFloat) Gen64P() *float64 {
	f := s.Gen64()
	return &f
}

// Gen32 returns a float32 value.
func (s *SomeFloat) Gen32() float32 {
	return float32(s.Gen64())
}

// Gen32P returns a float32 pointer.
func (s *SomeFloat) Gen32P() *float32 {
	f := s.Gen32()
	return &f
}
