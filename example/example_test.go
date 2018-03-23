package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVS_Get(t *testing.T) {
	some := &Some{}

	kvs := setupKVS(some)

	assert.Equal(t, kvs.Get(some.User("user1", UserSpec{}).ID), some.User("user1", UserSpec{}))
	assert.Equal(t, kvs.Get(some.User("user2", UserSpec{}).ID), some.User("user2", UserSpec{}))
	assert.Equal(t, kvs.Get(some.User("user3", UserSpec{}).ID), some.User("user3", UserSpec{}))
}

func setupKVS(some *Some) *KVS {
	u1 := some.User("user1", UserSpec{})
	u2 := some.User("user2", UserSpec{})
	u3 := some.User("user3", UserSpec{})

	m := map[int64]*User{}
	m[u1.ID] = u1
	m[u2.ID] = u2
	m[u3.ID] = u3

	return &KVS{m}
}

func BenchmarkGenUser(b *testing.B) {
	some := &Some{}
	for i := 0; i < b.N; i++ {
		some.User("user", UserSpec{})
	}
}
