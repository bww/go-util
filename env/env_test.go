package env

import (
  "os"
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
  
  os.Setenv("GOUTIL_BIN", "/x")
  assert.Equal(t, `/x/bin`, Bin())
  assert.Equal(t, `/x/bin/a`, Bin("a"))
  assert.Equal(t, `/x/bin/a/b`, Bin("a/b"))
  os.Setenv("GOUTIL_ETC", "/x")
  assert.Equal(t, `/x/etc`, Etc())
  assert.Equal(t, `/x/etc/a`, Etc("a"))
  assert.Equal(t, `/x/etc/a/b`, Etc("a/b"))
  os.Setenv("GOUTIL_WEB", "/x")
  assert.Equal(t, `/x/web`, Web())
  assert.Equal(t, `/x/web/a`, Web("a"))
  assert.Equal(t, `/x/web/a/b`, Web("a/b"))
  os.Unsetenv("GOUTIL_BIN")
  os.Unsetenv("GOUTIL_ETC")
  os.Unsetenv("GOUTIL_WEB")
  
  os.Setenv("TEST_BIN", "/x")
  assert.Equal(t, `/x/bin`, Bin())
  assert.Equal(t, `/x/bin/a`, Bin("a"))
  assert.Equal(t, `/x/bin/a/b`, Bin("a/b"))
  os.Setenv("TEST_ETC", "/x")
  assert.Equal(t, `/x/etc`, Etc())
  assert.Equal(t, `/x/etc/a`, Etc("a"))
  assert.Equal(t, `/x/etc/a/b`, Etc("a/b"))
  os.Setenv("TEST_WEB", "/x")
  assert.Equal(t, `/x/web`, Web())
  assert.Equal(t, `/x/web/a`, Web("a"))
  assert.Equal(t, `/x/web/a/b`, Web("a/b"))
  os.Unsetenv("TEST_BIN")
  os.Unsetenv("TEST_ETC")
  os.Unsetenv("TEST_WEB")
  
}
