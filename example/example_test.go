package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVS_Get(t *testing.T) {
	gen := Generator{}
	some := Some{}

	kvs := setupKVS(gen)

	assert.Equal(t, gen.User("user1", some.User().WithID(111)), kvs.Get(111))
	assert.Equal(t, gen.User("user2", some.User()), kvs.Get(gen.User("user2", some.User()).ID))
	assert.Equal(t, gen.User("user3", some.User()), kvs.Get(gen.User("user3", some.User()).ID))
}

func setupKVS(gen Generator) *KVS {
	some := Some{}
	u1 := gen.User("user1", some.User().WithID(111))
	u2 := gen.User("user2", some.User())
	u3 := gen.User("user3", some.User())

	m := map[int64]*User{}
	m[u1.ID] = u1
	m[u2.ID] = u2
	m[u3.ID] = u3

	return &KVS{m}
}

func BenchmarkGenUser(b *testing.B) {
	some := &Generator{}
	for i := 0; i < b.N; i++ {
		some.User("user", UserSpec{})
	}
}
