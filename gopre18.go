//+build !go1.8

package some

import "math/rand"

func (s *UintSpec) Generate(r *rand.Rand) uint {
	diff := s.max - s.min
	return s.min + (uint(uint64(r.Uint32())<<32|uint64(r.Uint32())) % diff)
}
