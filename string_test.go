package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)

	g := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.String(key).Gen(), g.String(key).Gen())
		assert.Equal(g.String(key).GenP(), g.String(key).GenP())

		assert.NotEqual(g.String().Gen(), g.String().Gen())

		v := g.String().Len(7).Gen()
		assert.Len(v, 7)
	}
}
