package concurrent

import (
	"sync"
	"time"
)

type Timer struct {
	*time.Timer
	mu sync.Mutex
}

func NewTimer(d time.Duration) *Timer {
	t := time.NewTimer(d)
	return &Timer{
		Timer: t,
	}
}

func (t *Timer) Reset(d time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Timer.Reset(d)
}
