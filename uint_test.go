package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	s := Some{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.Uint(key, Uint), s.Uint(key, Uint))

		v := s.Uint(key, Uint.Min(10).Max(200))
		if v < 10 {
			assert.Fail(t, "", "less than -100: %v", v)
		}
		if 200 <= v {
			assert.Fail(t, "", "greater than 50: %v", v)
		}
	}
}
