package iterator

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	var err error

	n := 10
	it := New(context.Background(), n)
	for i := 0; i < n; i++ {
		err = it.Write(Result{Elem: i})
		assert.Nil(t, err, fmt.Sprint(err))
	}

	err = it.Close()
	assert.Nil(t, err, fmt.Sprint(err))
	err = it.Close() // close twice, this should be fine
	assert.Nil(t, err, fmt.Sprint(err))

	for i := 0; ; i++ {
		val, err := it.Next()
		if err == ErrClosed {
			break
		} else {
			assert.Nil(t, err, fmt.Sprint(err))
		}
		assert.Equal(t, i, val.(int))
	}

}

func TestCancellation(t *testing.T) {
	cxt, cancel := context.WithCancel(context.Background())
	var err error

	n := 10
	it := New(cxt, n)
	for i := 0; i < n; i++ {
		err = it.Write(Result{Elem: i})
		assert.Nil(t, err, fmt.Sprint(err))
	}

	cancel()

	// this is deterministic because the channel buffer is full, so
	// cancellation is the only case that is ready when Write() is
	// called
	err = it.Write(Result{Elem: 100})
	assert.Equal(t, ErrCancelled, err)

	// this is not deterministic, so we have to process up to N before
	// we're guaranteed to get the cancellation error
	var found bool
	for i := 0; i < n; i++ {
		_, err = it.Next()
		if err != nil {
			assert.Equal(t, ErrCancelled, err)
			found = true
			break
		}
	}
	assert.Equal(t, true, found)

}
