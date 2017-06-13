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
  r := New("Hello").Warn(time.Millisecond)
  r.Start("Sub-operation").Finish()
  r.Start("Another operation").Finish()
  r.Start("Another operation").Finish()
  r.Start("Another operation").Finish()
  r.Start("Open operation")
  s := r.Start("Enjoy this one as well")
  u := s.Start("Sub-op")
  <- time.After(time.Millisecond * 1)
  u.Finish()
  d := u.Start("Sub-sub-op!")
  <- time.After(time.Millisecond * 1)
  d.Finish()
  u = s.Start("Sub-op again")
  <- time.After(time.Millisecond * 1)
  u.Finish()
  d = u.Start("Sub-sub-op again!")
  <- time.After(time.Millisecond * 1)
  d.Finish()
  u.Start("Sub-sub-op again!").Finish()
  u.Start("Sub-sub-op again!").Finish()
  s.Finish()
  r.Write(os.Stdout)
  // assert.Equal(t, true, ResemblesUUID("ACE24573-5BD5-4C5F-B143-5E9E17F18BDB"))
}
