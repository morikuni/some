package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	assert := assert.New(t)

	g := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.Uint(key).Gen64(), g.Uint(key).Gen64())
		assert.Equal(g.Uint(key).Gen64P(), g.Uint(key).Gen64P())
		assert.Equal(g.Uint(key).Gen(), g.Uint(key).Gen())
		assert.Equal(g.Uint(key).GenP(), g.Uint(key).GenP())

		assert.NotEqual(g.Uint().Gen(), g.Uint().Gen())
		assert.NotEqual(g.Uint().Gen64(), g.Uint().Gen64())

		v := g.Uint().Min(10).Max(200).Gen64()
		if v < 10 {
			assert.Fail("", "less than -100: %v", v)
		}
		if 200 <= v {
			assert.Fail("", "greater than 50: %v", v)
		}
	}
}
