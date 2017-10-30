//+build go1.8

package dummy

func (g *UintGen) Gen64() uint64 {
	diff := g.max - g.min
	// TODO: use Uint64n if it's supported by math/rand.
	return g.min + (g.r.Uint64() % diff)
}
