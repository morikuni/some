package dummy

import (
	"math/rand"
	"net/url"
)

func URL(r *rand.Rand) *URLGen {
	return &URLGen{
		r,
		[]string{"http", "https"},
	}
}

type URLGen struct {
	r       *rand.Rand
	schemes []string
}

func (g *URLGen) Schemes(schemes ...string) *URLGen {
	g.schemes = schemes
	return g
}

func (g *URLGen) Gen() *url.URL {
	scheme := g.schemes[Int(g.r).Max(len(g.schemes)).Gen()]
	return &url.URL{
		Scheme: scheme,
		Host:   String(g.r).Len(10).Gen() + ".com",
		Path:   String(g.r).Len(10).Gen() + "/" + String(g.r).Len(10).Gen(),
	}
}
