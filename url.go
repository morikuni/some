package some

import (
	"math/rand"
	"net/url"
)

// URL returns a url genertor.
func URL(r *rand.Rand) *SomeURL {
	return &SomeURL{
		r,
		[]string{"http", "https"},
	}
}

// SomeURL is a url generator.
type SomeURL struct {
	r       *rand.Rand
	schemes []string
}

// Schemes sets the schemes, one of them will be picked for
// a random url.
func (g *SomeURL) Schemes(schemes ...string) *SomeURL {
	g.schemes = schemes
	return g
}

// Gen returns a url value.
func (g *SomeURL) Gen() url.URL {
	scheme := g.schemes[Int(g.r).Max(len(g.schemes)).Gen()]
	return url.URL{
		Scheme: scheme,
		Host:   String(g.r).Len(10).Gen() + ".com",
		Path:   String(g.r).Len(10).Gen() + "/" + String(g.r).Len(10).Gen(),
	}
}

// GenP returns a url pointer.
func (g *SomeURL) GenP() *url.URL {
	u := g.Gen()
	return &u
}
