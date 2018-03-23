//+build !go1.8

package some

import "math/rand"

// Generate generates a random uint from r.
func (s *UintSpec) Generate(r *rand.Rand) uint {
	diff := s.Max - s.Min
	return s.Min + (uint(uint64(r.Uint32())<<32|uint64(r.Uint32())) % diff)
}
