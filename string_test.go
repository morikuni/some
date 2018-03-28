package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	g := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, g.String(key, String), g.String(key, String))

		v := g.String(key, StringSpec{Len: 7})
		assert.Len(t, v, 7)
	}
}
