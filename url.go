package dummy

import (
	"math/rand"
	"net/url"
)

// URL returns a url genertor.
func URL(r *rand.Rand) *URLGen {
	return &URLGen{
		r,
		[]string{"http", "https"},
	}
}

// URLGen is a url generator.
type URLGen struct {
	r       *rand.Rand
	schemes []string
}

// Schemes sets the schemes, one of them will be picked for
// a random url.
func (g *URLGen) Schemes(schemes ...string) *URLGen {
	g.schemes = schemes
	return g
}

// Gen returns a url value.
func (g *URLGen) Gen() url.URL {
	scheme := g.schemes[Int(g.r).Max(len(g.schemes)).Gen()]
	return url.URL{
		Scheme: scheme,
		Host:   String(g.r).Len(10).Gen() + ".com",
		Path:   String(g.r).Len(10).Gen() + "/" + String(g.r).Len(10).Gen(),
	}
}

// Gen returns a url pointer.
func (g *URLGen) GenP() *url.URL {
	u := g.Gen()
	return &u
}
