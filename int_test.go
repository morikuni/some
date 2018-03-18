package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	assert := assert.New(t)

	g := Some{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.Int(key).Gen64(), g.Int(key).Gen64())
		assert.Equal(g.Int(key).Gen64P(), g.Int(key).Gen64P())
		assert.Equal(g.Int(key).Gen(), g.Int(key).Gen())
		assert.Equal(g.Int(key).GenP(), g.Int(key).GenP())

		assert.NotEqual(g.Int().Gen(), g.Int().Gen())
		assert.NotEqual(g.Int().Gen64(), g.Int().Gen64())

		v := g.Int().Min(-100).Max(50).Gen64()
		if v < -100 {
			assert.Fail("", "less than -100: %v", v)
		}
		if 50 <= v {
			assert.Fail("", "greater than 50: %v", v)
		}
	}
}
