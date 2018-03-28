package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	s := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.Int(key, AnyInt), s.Int(key, AnyInt))
		assert.Equal(t, s.Int64(key, AnyInt64), s.Int64(key, AnyInt64))

		v := s.Int(key, IntSpec{Min: -100, Max: 50})
		if v < -100 {
			assert.Fail(t, "", "less than -100: %v", v)
		}
		if 50 <= v {
			assert.Fail(t, "", "greater than 50: %v", v)
		}
	}
}
