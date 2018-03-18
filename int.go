package some

import (
	"math"
	"math/rand"
	"strconv"
)

// Int returns a int generator.
func Int(r *rand.Rand) *SomeInt {
	return &SomeInt{
		r,
		math.MaxInt64,
		0,
	}
}

// SomeInt is a int generator.
type SomeInt struct {
	r   *rand.Rand
	max int64
	min int64
}

// Max sets the maximum value of a random int value.
func (s *SomeInt) Max(n int) *SomeInt {
	return s.Max64(int64(n))
}

// Min sets the minimum value of a random int value.
func (s *SomeInt) Min(n int) *SomeInt {
	return s.Min64(int64(n))
}

// Max64 sets the maximum value of a random int value with int64.
func (s *SomeInt) Max64(n int64) *SomeInt {
	s.max = n
	return s
}

// Min64 sets the minimum value of a random int value with int64.
func (s *SomeInt) Min64(n int64) *SomeInt {
	s.min = n
	return s
}

// Gen  returns a int value.
func (s *SomeInt) Gen() int {
	return int(s.Gen64())
}

// GenP returns a int pointer.
func (s *SomeInt) GenP() *int {
	i := s.Gen()
	return &i
}

// Gen64 returns a int64 value.
func (s *SomeInt) Gen64() int64 {
	diff := s.max - s.min
	return s.min + s.r.Int63n(diff)
}

// Gen64P returns a int64 pointer.
func (s *SomeInt) Gen64P() *int64 {
	i := s.Gen64()
	return &i
}

// GenString returns a int64 value as a string value.
func (s *SomeInt) GenString() string {
	return strconv.FormatInt(s.Gen64(), 10)
}

// GenStringP returns a int64 value as a string pointer.
func (s *SomeInt) GenStringP() *string {
	str := s.GenString()
	return &str
}
