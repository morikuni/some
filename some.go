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

// Some generates and caches a dummy data.
type Some struct {
	// LocalRand is a instance specific random number generator.
	LocalRand *rand.Rand
	// LocalCache is a instance specific cache store.
	LocalCache map[string]interface{}
}

// Rand returns a random number generator for the specific key.
// GlobalRand or LocalCache are returned for the empty key.
func (s *Some) Rand(key string) *rand.Rand {
	if key != "" {
		hash := s.hash(key)
		return rand.New(rand.NewSource(hash))
	}
	if s.LocalRand != nil {
		return s.LocalRand
	}
	return GlobalRand
}

// Cache returns a cache store.
// GlobalRand or LocalCache are returned.
func (s *Some) Cache() map[string]interface{} {
	if s.LocalCache != nil {
		return s.LocalCache
	}
	return GlobalCache
}

func (s *Some) cacheKey(key string, cacheKey CacheKey) string {
	return reflect.TypeOf(cacheKey).String() + "@" + key + "@" + cacheKey.CacheKey()
}

func (s *Some) hash(key string) int64 {
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}

func (s *Some) Generate(key string, spec interface{}, f func(r *rand.Rand) interface{}) interface{} {
	r := s.Rand(key)
	if ck, ok := spec.(CacheKey); ok && key != "" {
		cache := s.Cache()
		cacheKey := s.cacheKey(key, ck)
		if v, ok := cache[cacheKey]; ok {
			return v
		}
		v := f(r)
		cache[cacheKey] = v
		return v
	} else {
		return f(r)
	}
}

type CacheKey interface {
	CacheKey() string
}
