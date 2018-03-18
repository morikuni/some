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

// Some generates and caches a dummy data.
type Some struct {
	// LocalRand is a instance specific random number generator.
	LocalRand *rand.Rand
	// LocalCache is a instance specific cache store.
	LocalCache map[string]interface{}
}

// Rand returns a random number generator for the specific key.
// GlobalRand or LocalCache are returned for the empty key.
func (s *Some) Rand(keys []string) *rand.Rand {
	if len(keys) > 0 {
		hash := s.hash(keys)
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

// CacheKey returns a unique key for the specific key and type.
func (s *Some) CacheKey(keys []string, typ interface{}) string {
	return reflect.TypeOf(typ).String() + strconv.FormatInt(s.hash(keys), 10)
}

func (s *Some) hash(keys []string) int64 {
	key := strings.Join(keys, "")
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}

// WithCache returns the result of the function f and cache it with the specific key and type.
// Note: Don't cache a pointer type, because the pointer value may be overwrited.
func (s *Some) WithCache(keys []string, typ interface{}, f func() interface{}) interface{} {
	if len(keys) == 0 {
		return f()
	}
	cache := s.Cache()
	cacheKey := s.CacheKey(keys, typ)
	if v, ok := cache[cacheKey]; ok {
		return v
	}
	v := f()
	cache[cacheKey] = v
	return v
}

// Int returns a int generator for the specific key.
func (s *Some) Int(key ...string) *SomeInt {
	return Int(s.Rand(key))
}

// Uint returns a uint generator for the specific key.
func (s *Some) Uint(key ...string) *SomeUint {
	return Uint(s.Rand(key))
}

// Float returns a float generator for the specific key.
func (s *Some) Float(key ...string) *SomeFloat {
	return Float(s.Rand(key))
}

// String returns a string generator for the specific key.
func (s *Some) String(key ...string) *SomeString {
	return String(s.Rand(key))
}

// Bool returns a bool generator for the specific key.
func (s *Some) Bool(key ...string) *SomeBool {
	return Bool(s.Rand(key))
}

// URL returns a url generator for the specific key.
func (s *Some) URL(key ...string) *SomeURL {
	return URL(s.Rand(key))
}

// Time returns a time generator for the specific key.
func (s *Some) Time(key ...string) *SomeTime {
	return Time(s.Rand(key))
}
