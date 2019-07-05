package list

import (
	"sync"
	"sync/atomic"
)

type Env struct {
	sync.RWMutex

	id  uint64
	mu  sync.Mutex
	cfg atomic.Value
}
