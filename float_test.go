package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	assert := assert.New(t)

	g := Some{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.Float(key).Gen64(), g.Float(key).Gen64())
		assert.Equal(g.Float(key).Gen64P(), g.Float(key).Gen64P())
		assert.Equal(g.Float(key).Gen32(), g.Float(key).Gen32())
		assert.Equal(g.Float(key).Gen32P(), g.Float(key).Gen32P())

		assert.NotEqual(g.Float().Gen64(), g.Float().Gen64())
		assert.NotEqual(g.Float().Gen32(), g.Float().Gen32())

		v := g.Float().Min(-100).Max(50).Gen64()
		if v < -100 {
			assert.Fail("", "less than -100: %v", v)
		}
		if 50 <= v {
			assert.Fail("", "greater than 50: %v", v)
		}
	}
}
