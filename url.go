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
func (s *SomeURL) Schemes(schemes ...string) *SomeURL {
	s.schemes = schemes
	return s
}

// Gen returns a url value.
func (s *SomeURL) Gen() url.URL {
	scheme := s.schemes[Int(s.r).Max(len(s.schemes)).Gen()]
	return url.URL{
		Scheme: scheme,
		Host:   String(s.r).Len(10).Gen() + ".com",
		Path:   String(s.r).Len(10).Gen() + "/" + String(s.r).Len(10).Gen(),
	}
}

// GenP returns a url pointer.
func (s *SomeURL) GenP() *url.URL {
	u := s.Gen()
	return &u
}
