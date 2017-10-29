package gen

import (
	"hash/fnv"
	"math/rand"
	"strings"
	"time"
)

var GlobalRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Gen struct {
	Default *rand.Rand
}

func (g *Gen) Rand(keys []string) *rand.Rand {
	if len(keys) > 0 {
		key := strings.Join(keys, "")
		hash := fnv.New64()
		hash.Write([]byte(key))
		return rand.New(rand.NewSource(int64(hash.Sum64())))
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
