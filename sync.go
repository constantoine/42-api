package api

import (
	"sync"
	"time"
)

type Sync struct {
	last time.Time
	sync.Mutex
}

func (s *Sync) lock() {
	s.Lock()
	if time.Now().Before(s.last.Add(500 * time.Millisecond)) {
		time.Sleep(time.Until(s.last.Add(500 * time.Millisecond)))
	}
}

func (s *Sync) unlock() {
	s.last = time.Now()
	s.Unlock()
}
