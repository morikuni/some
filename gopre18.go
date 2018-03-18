//+build !go1.8

package some

// Gen64 returns a uint64 value.
func (s *SomeUint) Gen64() uint64 {
	diff := s.max - s.min
	return s.min + (uint64(s.r.Uint32())<<32 | uint64(s.r.Uint32())%diff)
}
