package some

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	assert := assert.New(t)

	g := Some{}

	assert.Equal(GlobalRand, g.Rand(nil))
	assert.Equal(GlobalCache, g.Cache())

	r := rand.New(rand.NewSource(12345))
	c := map[string]interface{}{}

	g.LocalRand = r
	g.LocalCache = c

	assert.Equal(r, g.Rand(nil))
	assert.Equal(c, g.Cache())

	assert.Equal(g.CacheKey([]string{"a"}, ""), g.CacheKey([]string{"a"}, ""))
	assert.NotEqual(g.CacheKey([]string{"a"}, ""), g.CacheKey([]string{"b"}, ""))
	assert.NotEqual(g.CacheKey([]string{"a"}, ""), g.CacheKey([]string{"a"}, 0))

	count := 0
	f := func() interface{} {
		count++
		return String(g.Rand(nil)).Gen()
	}
	g.WithCache(nil, "", f)
	g.WithCache(nil, "", f)
	assert.Equal(g.WithCache([]string{"a"}, "", f), g.WithCache([]string{"a"}, "", f))
	assert.NotEqual(g.WithCache([]string{"a"}, "", f), g.WithCache([]string{"b"}, "", f))
	g.WithCache([]string{"a"}, "", f)
	g.WithCache([]string{"a"}, "", f)
	g.WithCache([]string{"a"}, "", f)
	assert.Equal(4, count)
}
