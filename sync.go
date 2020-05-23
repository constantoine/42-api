package api

import (
	"sync"
	"time"
)

var synchro struct {
	last time.Time
	sync.Mutex
}

func lock() {
	synchro.Lock()
	if time.Now().Before(synchro.last.Add(500 * time.Millisecond)) {
		time.Sleep(time.Until(synchro.last.Add(500 * time.Millisecond)))
	}
}

func unlock() {
	synchro.last = time.Now()
	synchro.Unlock()
}
