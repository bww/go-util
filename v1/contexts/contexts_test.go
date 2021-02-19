package contexts

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContinue(t *testing.T) {
	var wg sync.WaitGroup
	var iter, done int32

	cxt, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		for {
			atomic.AddInt32(&iter, 1)
			if !Continue(cxt) {
				break // this construct is weird for testing purposes
			}
		}
		atomic.AddInt32(&done, 1)
		wg.Done()
	}()

	go cancel()
	wg.Wait()

	assert.Equal(t, true, atomic.LoadInt32(&iter) > 0)
	assert.Equal(t, int32(1), atomic.LoadInt32(&done))
}
