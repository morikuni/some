package dummy

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

// Gen is a factory and cache for a dummy data generator.
type Gen struct {
	// LocalRand is a instace specific random number generator.
	LocalRand *rand.Rand
	// LocalCache is a instace specific cache store.
	LocalCache map[string]interface{}
}

// Rand returns a random number generator for the specific key.
// GlobalRand or LocalCache are returned for the empty key.
func (g *Gen) Rand(keys []string) *rand.Rand {
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
func (g *Gen) Cache() map[string]interface{} {
	if g.LocalCache != nil {
		return g.LocalCache
	}
	return GlobalCache
}

// CacheKey returns a unique key for the specific key and type.
func (g *Gen) CacheKey(keys []string, typ interface{}) string {
	return reflect.TypeOf(typ).String() + strconv.FormatInt(g.hash(keys), 10)
}

func (g *Gen) hash(keys []string) int64 {
	key := strings.Join(keys, "")
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}

// WithCache returns the result of the function f and cache it with the specific key and type.
// Note: Don't cache a pointer type, because the pointer value may be overwrited.
func (g *Gen) WithCache(keys []string, typ interface{}, f func() interface{}) interface{} {
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
func (g *Gen) Int(key ...string) *IntGen {
	return Int(g.Rand(key))
}

// Uint returns a uint generator for the specific key.
func (g *Gen) Uint(key ...string) *UintGen {
	return Uint(g.Rand(key))
}

// Float returns a float generator for the specific key.
func (g *Gen) Float(key ...string) *FloatGen {
	return Float(g.Rand(key))
}

// String returns a string generator for the specific key.
func (g *Gen) String(key ...string) *StringGen {
	return String(g.Rand(key))
}

// Bool returns a bool generator for the specific key.
func (g *Gen) Bool(key ...string) *BoolGen {
	return Bool(g.Rand(key))
}

// URL returns a url generator for the specific key.
func (g *Gen) URL(key ...string) *URLGen {
	return URL(g.Rand(key))
}

// Time returns a time generator for the specific key.
func (g *Gen) Time(key ...string) *TimeGen {
	return Time(g.Rand(key))
}
