//+build go1.8

package some

// Gen64 returns a uint64 value.
func (s *SomeUint) Gen64() uint64 {
	diff := s.max - s.min
	// TODO: use Uint64n if it's supported by math/rand.
	return s.min + (s.r.Uint64() % diff)
}
