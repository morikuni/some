package some

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSpec struct{}

func (s testSpec) Generate(r *rand.Rand) int {
	return r.Int()
}

func (s testSpec) CacheKey() string {
	return "test"
}

func TestSome(t *testing.T) {
	s := Generator{}

	assert.Equal(t, GlobalRand, s.Rand(""))
	assert.Equal(t, GlobalCache, s.Cache())

	r := rand.New(rand.NewSource(12345))
	c := map[string]interface{}{}

	s.LocalRand = r
	s.LocalCache = c

	assert.Equal(t, r, s.Rand(""))
	assert.Equal(t, c, s.Cache())

	count := 0
	spec := testSpec{}
	f := func(r *rand.Rand) interface{} {
		count++
		return spec.Generate(r)
	}
	s.Generate("", spec, f)
	s.Generate("", spec, f)
	assert.Equal(t, s.Generate("a", spec, f), s.Generate("a", spec, f))
	assert.NotEqual(t, s.Generate("a", spec, f), s.Generate("b", spec, f))
	s.Generate("a", spec, f)
	s.Generate("a", spec, f)
	s.Generate("a", spec, f)
	assert.Equal(t, 4, count)
}
