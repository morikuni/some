package some

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	assert := assert.New(t)

	g := Some{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.Time(key).Gen(), g.Time(key).Gen())
		assert.Equal(g.Time(key).GenP(), g.Time(key).GenP())

		assert.NotEqual(g.Time().Gen(), g.Time().Gen())

		now := time.Now()
		v := g.Time().After(now.Add(-time.Hour * 24 * 10)).Before(now).GenP()
		if v.Before(now.Add(-time.Hour * 24 * 10)) {
			assert.Fail("", "before 10 days: %v", v)
		}
		if v.After(now) {
			assert.Fail("", "after now: %v", v)
		}
	}
}
