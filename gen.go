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
	GlobalRand  *rand.Rand             = rand.New(rand.NewSource(time.Now().UnixNano()))
	GlobalCache map[string]interface{} = make(map[string]interface{})
)

type Gen struct {
	LocalRand  *rand.Rand
	LocalCache map[string]interface{}
}

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

func (g *Gen) Cache() map[string]interface{} {
	if g.LocalCache != nil {
		return g.LocalCache
	}
	return GlobalCache
}

func (g *Gen) CacheKey(keys []string, typ interface{}) string {
	return reflect.TypeOf(typ).String() + strconv.FormatInt(g.hash(keys), 10)
}

func (g *Gen) hash(keys []string) int64 {
	key := strings.Join(keys, "")
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}

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

func (g *Gen) Int(key ...string) *IntGen {
	return Int(g.Rand(key))
}

func (g *Gen) Uint(key ...string) *UintGen {
	return Uint(g.Rand(key))
}

func (g *Gen) Float(key ...string) *FloatGen {
	return Float(g.Rand(key))
}

func (g *Gen) String(key ...string) *StringGen {
	return String(g.Rand(key))
}

func (g *Gen) Bool(key ...string) *BoolGen {
	return Bool(g.Rand(key))
}

func (g *Gen) URL(key ...string) *URLGen {
	return URL(g.Rand(key))
}

func (g *Gen) Time(key ...string) *TimeGen {
	return Time(g.Rand(key))
}
