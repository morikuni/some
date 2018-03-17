package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVS_Get(t *testing.T) {
	assert := assert.New(t)

	gen := &Gen{}

	kvs := setupKVS(gen)

	assert.Equal(kvs.Get(gen.User("user1").ID), gen.User("user1"))
	assert.Equal(kvs.Get(gen.User("user2").ID), gen.User("user2"))
	assert.Equal(kvs.Get(gen.User("user3").ID), gen.User("user3"))
}

func setupKVS(gen *Gen) *KVS {
	u1 := gen.User("user1")
	u2 := gen.User("user2")
	u3 := gen.User("user3")

	m := map[int64]*User{}
	m[u1.ID] = u1
	m[u2.ID] = u2
	m[u3.ID] = u3

	return &KVS{m}
}

func BenchmarkGenUser(b *testing.B) {
	gen := &Gen{}
	for i := 0; i < b.N; i++ {
		gen.User("user")
	}
}
