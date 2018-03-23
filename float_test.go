package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	s := Some{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.Float(key, Float), s.Float(key, Float))

		v := s.Float(key, Float.Min(-100).Max(50))
		if v < -100 {
			assert.Fail(t, "", "less than -100: %v", v)
		}
		if 50 <= v {
			assert.Fail(t, "", "greater than 50: %v", v)
		}
	}
}
