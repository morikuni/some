//+build !go1.8

package dummy

func (g *UintGen) Gen64() uint64 {
	diff := g.max - g.min
	return g.min + (uint64(g.r.Uint32())<<32 | uint64(g.r.Uint32())%diff)
}
