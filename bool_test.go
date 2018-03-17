package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	assert := assert.New(t)

	g := Generator{}

	n := 1000
	trueCount := 0.0
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.Bool(key).Gen(), g.Bool(key).Gen())
		assert.Equal(g.Bool(key).GenP(), g.Bool(key).GenP())

		if g.Bool().Gen() {
			trueCount += 1
		}
	}
	assert.InDelta(0.5, trueCount/float64(n), 0.1)
}
