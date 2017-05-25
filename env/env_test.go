package env

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func init() {
  prefix = "TEST_"
  environ = "test"
  home = "/"
}

/**
 * Test env
 */
func TestEnv(t *testing.T) {
  assert.Equal(t, `/`, Resource())
  assert.Equal(t, `/a`, Resource("a"))
  assert.Equal(t, `/a/b`, Resource("a/b"))
  assert.Equal(t, `/bin`, Bin())
  assert.Equal(t, `/bin/a`, Bin("a"))
  assert.Equal(t, `/bin/a/b`, Bin("a/b"))
  assert.Equal(t, `/etc`, Etc())
  assert.Equal(t, `/etc/a`, Etc("a"))
  assert.Equal(t, `/etc/a/b`, Etc("a/b"))
  assert.Equal(t, `/web`, Web())
  assert.Equal(t, `/web/a`, Web("a"))
  assert.Equal(t, `/web/a/b`, Web("a/b"))
}
