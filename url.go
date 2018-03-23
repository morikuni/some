package some

import (
	"math/rand"
	"net/url"
)

// AnyURL is a default url.URL spec.
var AnyURL URLSpec = URLSpec{
	[]string{"http", "https"},
}

// URLSpec is a spec of a url.URL.
type URLSpec struct {
	Schemes []string
}

// Generate generates a random url.URL from r.
func (s URLSpec) Generate(r *rand.Rand) url.URL {
	scheme := s.Schemes[IntSpec{Max: len(s.Schemes)}.Generate(r)]
	return url.URL{
		Scheme: scheme,
		Host:   StringSpec{Len: 10}.Generate(r) + ".com",
		Path:   StringSpec{Len: 10}.Generate(r) + "/" + StringSpec{Len: 10}.Generate(r),
	}
}

// URL generates a url.URL according to a key and spec.
func (s *Some) URL(key string, spec URLSpec) url.URL {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(url.URL)
}
