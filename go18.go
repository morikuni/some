//+build go1.8

package some

import "math/rand"

// Generate generates a random uint from r.
func (s *UintSpec) Generate(r *rand.Rand) uint {
	diff := s.Max - s.Min
	// TODO: use Uint64n if it's supported by math/rand.
	return s.Min + (uint(r.Uint64()) % diff)
}
