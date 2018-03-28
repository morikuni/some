package some

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	s := Generator{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(t, s.URL(key, URL), s.URL(key, URL))

		v := s.URL(key, URLSpec{Schemes: []string{"hoge", "huga"}})
		u, err := url.Parse(v.String())
		assert.NoError(t, err)
		assert.Contains(t, []string{"hoge", "huga"}, u.Scheme)
	}
}
