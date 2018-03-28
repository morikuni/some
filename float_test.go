package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	s := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.Float64(key, AnyFloat64), s.Float64(key, AnyFloat64))
		assert.Equal(t, s.Float32(key, AnyFloat32), s.Float32(key, AnyFloat32))

		v := s.Float64(key, Float64Spec{Min: -100, Max: 50})
		if v < -100 {
			assert.Fail(t, "", "less than -100: %v", v)
		}
		if 50 <= v {
			assert.Fail(t, "", "greater than 50: %v", v)
		}
	}
}
