package trace

import (
  "os"
  "time"
  "testing"
  // "github.com/stretchr/testify/assert"
)

/**
 * Test trace
 */
func TestTrace(t *testing.T) {
  r := New("Hello")
  r.Start("Sub-operation").Finish()
  r.Start("Another operation").Finish()
  r.Start("Open operation")
  s := r.Start("Enjoy this one as well")
  <- time.After(time.Second * 2)
  s.Finish()
  r.Write(os.Stdout)
  // assert.Equal(t, true, ResemblesUUID("ACE24573-5BD5-4C5F-B143-5E9E17F18BDB"))
}
