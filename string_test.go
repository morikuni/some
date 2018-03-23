package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	g := Some{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, g.String(key, AnyString), g.String(key, AnyString))

		v := g.String(key, StringSpec{Len: 7})
		assert.Len(t, v, 7)
	}
}
