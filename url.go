package some

import (
	"math/rand"
	"net/url"
)

var URL URLSpec = URLSpec{
	[]string{"http", "https"},
}

type URLSpec struct {
	schemes []string
}

// Schemes sets the schemes, one of them will be picked for
// a random url.
func (s URLSpec) Schemes(schemes ...string) URLSpec {
	s.schemes = schemes
	return s
}

func (s URLSpec) Generate(r *rand.Rand) url.URL {
	scheme := s.schemes[Int.Max(len(s.schemes)).Generate(r)]
	return url.URL{
		Scheme: scheme,
		Host:   String.Len(10).Generate(r) + ".com",
		Path:   String.Len(10).Generate(r) + "/" + String.Len(10).Generate(r),
	}
}

func (s *Some) URL(key string, spec URLSpec) url.URL {
	return s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(url.URL)
}
