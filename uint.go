package some

import (
	"math"
	"math/rand"
	"strconv"
)

// Uint returns a uint generator.
func Uint(r *rand.Rand) *SomeUint {
	return &SomeUint{
		r,
		math.MaxUint64,
		0,
	}
}

// SomeUint is a uint generator.
type SomeUint struct {
	r   *rand.Rand
	max uint64
	min uint64
}

// Max sets the maximum value of a random uint value.
func (s *SomeUint) Max(n uint) *SomeUint {
	return s.Max64(uint64(n))
}

// Min sets the minimum value of a random uint value.
func (s *SomeUint) Min(n uint) *SomeUint {
	return s.Min64(uint64(n))
}

// Max64 sets the maximum value of a random uint value with uint64.
func (s *SomeUint) Max64(n uint64) *SomeUint {
	s.max = n
	return s
}

// Min64 sets the minimum value of a random uint value with uint64.
func (s *SomeUint) Min64(n uint64) *SomeUint {
	s.min = n
	return s
}

// Gen returns a uint value.
func (s *SomeUint) Gen() uint {
	return uint(s.Gen64())
}

// GenP returns a uint pointer.
func (s *SomeUint) GenP() *uint {
	u := s.Gen()
	return &u
}

// Gen64P returns a uint64 pointer.
func (s *SomeUint) Gen64P() *uint64 {
	u := s.Gen64()
	return &u
}

// GenString returns a uint64 value as a string value.
func (s *SomeUint) GenString() string {
	return strconv.FormatUint(s.Gen64(), 10)
}

// GenStringP returns a uint64 value as a string value.
func (s *SomeUint) GenStringP() *string {
	str := s.GenString()
	return &str
}
