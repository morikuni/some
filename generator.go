package some

import (
	"hash/fnv"
	"math/rand"
	"reflect"
	"time"
)

//TODO: make cache to interface, then implement concurrent safe cache.

var (
	// GlobalRand is a default random number generator.
	GlobalRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	// GlobalCache is a default cache store.
	GlobalCache map[string]interface{} = make(map[string]interface{})
)

// Generator generates and caches a dummy data.
type Generator struct {
	// LocalRand is a instance specific random number generator.
	LocalRand *rand.Rand
	// LocalCache is a instance specific cache store.
	LocalCache map[string]interface{}
}

// Rand returns a random number generator for the specific key.
// GlobalRand or LocalCache are returned for the empty key.
func (g *Generator) Rand(key string) *rand.Rand {
	if key != "" {
		hash := g.hash(key)
		return rand.New(rand.NewSource(hash))
	}
	if g.LocalRand != nil {
		return g.LocalRand
	}
	return GlobalRand
}

// Cache returns a cache store.
// GlobalRand or LocalCache are returned.
func (g *Generator) Cache() map[string]interface{} {
	if g.LocalCache != nil {
		return g.LocalCache
	}
	return GlobalCache
}

func (g *Generator) cacheKey(key string, cacheKey CacheKey) string {
	return reflect.TypeOf(cacheKey).String() + "@" + key + "@" + cacheKey.CacheKey()
}

func (g *Generator) hash(key string) int64 {
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}

func (g *Generator) Generate(key string, spec interface{}, f func(r *rand.Rand) interface{}) interface{} {
	if ck, ok := spec.(CacheKey); ok && key != "" {
		cache := g.Cache()
		cacheKey := g.cacheKey(key, ck)
		if v, ok := cache[cacheKey]; ok {
			return v
		}
		r := g.Rand(key)
		v := f(r)
		cache[cacheKey] = v
		return v
	} else {
		r := g.Rand(key)
		return f(r)
	}
}

type CacheKey interface {
	CacheKey() string
}
