package dummy

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)

	g := Gen{}

	n := 100
	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		assert.Equal(g.URL(key).Gen(), g.URL(key).Gen())
		assert.Equal(g.URL(key).GenP(), g.URL(key).GenP())

		assert.NotEqual(g.URL().Gen(), g.URL().Gen())

		v := g.URL().Schemes("hoge", "huga").Gen()
		u, err := url.Parse(v.String())
		assert.NoError(err)
		assert.Contains([]string{"hoge", "huga"}, u.Scheme)
	}
}
