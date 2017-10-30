package gen

import (
	"hash/fnv"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var GlobalRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func New() *Gen {
	return &Gen{
		Default: GlobalRand,
		Cache:   make(map[string]interface{}),
	}
}

type Gen struct {
	Default *rand.Rand
	Cache   map[string]interface{}
}

func (g *Gen) Rand(keys []string) *rand.Rand {
	if len(keys) > 0 {
		hash := g.Hash(keys)
		return rand.New(rand.NewSource(hash))
	}
	if g.Default == nil {
		return GlobalRand
	}
	return g.Default
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

func (g *Gen) WithCache(key []string, typ interface{}, f func() interface{}) interface{} {
	if len(key) == 0 {
		return f()
	}
	cacheKey := reflect.TypeOf(typ).String() + strconv.FormatInt(g.Hash(key), 10)
	if v, ok := g.Cache[cacheKey]; ok {
		return v
	}
	v := f()
	g.Cache[cacheKey] = v
	return v
}

func (g *Gen) Hash(keys []string) int64 {
	key := strings.Join(keys, "")
	hash := fnv.New64()
	hash.Write([]byte(key))
	return int64(hash.Sum64())
}
