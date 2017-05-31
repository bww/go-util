package text

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

/**
 * Test normalize string
 */
func TestIndent(t *testing.T) {
  assert.Equal(t, `> Hello`, Indent(`Hello`, `> `))
  assert.Equal(t, `> Hello
> There`, Indent(`Hello
There`, `> `))
  assert.Equal(t, `> Hello
> There
> `, Indent(`Hello
There
`, `> `))
}
