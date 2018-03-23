//+build go1.8

package some

import "math/rand"

func (s *UintSpec) Generate(r *rand.Rand) uint {
	diff := s.max - s.min
	// TODO: use Uint64n if it's supported by math/rand.
	return s.min + (uint(r.Uint64()) % diff)
}
