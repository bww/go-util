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
	cxt    context.Context
	data   chan Result
	closer sync.Once
}

func New(cxt context.Context, buf int) Iterator {
	return Iterator{
		cxt:  cxt,
		data: make(chan Result, buf),
	}
}

func (t Iterator) Write(r Result) error {
	select {
	case <-t.cxt.Done():
		return ErrCancelled
	case t.data <- r:
		return nil
	}
}

func (t Iterator) Close() error {
	t.closer.Do(func() {
		close(t.data)
	})
	return nil
}

func (t Iterator) Next() (interface{}, error) {
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
