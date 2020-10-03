package throttle

import (
	"context"
	"sync"
	"time"
)

type Throttle struct {
	sync.Mutex
	ops   int
	until time.Time
}

func New() *Throttle {
	return &Throttle{}
}

func (t *Throttle) Update(ops int, until time.Time) *Throttle {
	t.Lock()
	defer t.Unlock()
	if ops < 0 {
		t.ops = 0
	} else {
		t.ops = ops
	}
	t.until = until
	return t
}

func (t *Throttle) Ops() int {
	t.Lock()
	defer t.Unlock()
	return t.ops
}

func (t *Throttle) Dec(ops int) int {
	t.Lock()
	defer t.Unlock()
	if r := t.ops - ops; r >= 0 {
		t.ops = r
	} else {
		t.ops = 0
	}
	return t.ops
}

func (t *Throttle) delay() time.Duration {
	if t.ops > 0 {
		return 0
	} else {
		return time.Now().Sub(t.until)
	}
}

func (t *Throttle) Delay() time.Duration {
	t.Lock()
	defer t.Unlock()
	return t.delay()
}

func (t *Throttle) Wait(cxt context.Context) (int, int) {
	return t.WaitN(cxt, 1)
}

func (t *Throttle) WaitN(cxt context.Context, ops int) (int, int) {
	d := t.Delay()
	if d > 0 {
		select {
		case <-time.After(d):
			break
		case <-cxt.Done():
			return -1, -1
		}
	}
	t.Lock()
	defer t.Unlock()
	n := ops
	if n > t.ops {
		n = t.ops
	}
	t.ops -= n
	return n, t.ops
}
