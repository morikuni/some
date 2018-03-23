package some

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	s := Some{}

	n := 1000
	trueCount := 0.0
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.Bool(key, Bool), s.Bool(key, Bool))

		if s.Bool(key, Bool) {
			trueCount += 1
		}
	}
	assert.InDelta(t, 0.5, trueCount/float64(n), 0.1)
}
