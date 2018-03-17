package some

import (
	"hash/fnv"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

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
func (g *Generator) Rand(keys []string) *rand.Rand {
	if len(keys) > 0 {
		hash := g.hash(keys)
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

// CacheKey returns a unique key for the specific key and type.
func (g *Generator) CacheKey(keys []string, typ interface{}) string {
	return reflect.TypeOf(typ).String() + strconv.FormatInt(g.hash(keys), 10)
}

func (g *Generator) hash(keys []string) int64 {
	key := strings.Join(keys, "")
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}

// WithCache returns the result of the function f and cache it with the specific key and type.
// Note: Don't cache a pointer type, because the pointer value may be overwrited.
func (g *Generator) WithCache(keys []string, typ interface{}, f func() interface{}) interface{} {
	if len(keys) == 0 {
		return f()
	}
	cache := g.Cache()
	cacheKey := g.CacheKey(keys, typ)
	if v, ok := cache[cacheKey]; ok {
		return v
	}
	v := f()
	cache[cacheKey] = v
	return v
}

// Int returns a int generator for the specific key.
func (g *Generator) Int(key ...string) *SomeInt {
	return Int(g.Rand(key))
}

// Uint returns a uint generator for the specific key.
func (g *Generator) Uint(key ...string) *SomeUint {
	return Uint(g.Rand(key))
}

// Float returns a float generator for the specific key.
func (g *Generator) Float(key ...string) *SomeFloat {
	return Float(g.Rand(key))
}

// String returns a string generator for the specific key.
func (g *Generator) String(key ...string) *SomeString {
	return String(g.Rand(key))
}

// Bool returns a bool generator for the specific key.
func (g *Generator) Bool(key ...string) *SomeBool {
	return Bool(g.Rand(key))
}

// URL returns a url generator for the specific key.
func (g *Generator) URL(key ...string) *SomeURL {
	return URL(g.Rand(key))
}

// Time returns a time generator for the specific key.
func (g *Generator) Time(key ...string) *SomeTime {
	return Time(g.Rand(key))
}
