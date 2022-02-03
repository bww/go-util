package iterator

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrClosed    = errors.New("Closed")
	ErrCancelled = errors.New("Cancelled")
)

type Result struct {
	Elem  interface{}
	Error error
}

type Iterator struct {
	mx     sync.RWMutex
	cxt    context.Context
	data   chan Result
	closed bool
}

func New(cxt context.Context, buf int) *Iterator {
	return &Iterator{
		mx:   sync.RWMutex{},
		cxt:  cxt,
		data: make(chan Result, buf),
	}
}

func (t *Iterator) Write(r Result) error {
	t.mx.RLock()
	defer t.mx.RUnlock()
	if t.closed {
		return ErrClosed
	}
	select {
	case <-t.cxt.Done():
		return ErrCancelled
	case t.data <- r:
		return nil
	}
}

func (t *Iterator) Close() error {
	t.mx.Lock()
	defer t.mx.Unlock()
	if !t.closed {
		close(t.data)
		t.closed = true
	}
	return nil
}

func (t *Iterator) Next() (interface{}, error) {
	// don't need to check closed	for read
	select {
	case <-t.cxt.Done():
		return nil, ErrCancelled
	case v, ok := <-t.data:
		if ok {
			return v.Elem, v.Error
		} else {
			return nil, ErrClosed
		}
	}
}
