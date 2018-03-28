package some

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	s := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.Time(key, Time), s.Time(key, Time))

		now := time.Now()
		v := s.Time(key, TimeSpec{After: now.Add(-time.Hour * 24 * 10), Before: now})
		if v.Before(now.Add(-time.Hour * 24 * 10)) {
			assert.Fail(t, "", "before 10 days: %v", v)
		}
		if v.After(now) {
			assert.Fail(t, "", "after now: %v", v)
		}
	}
}
