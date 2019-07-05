package list

import (
	"sync"
)

type Cache struct {
	sync.RWMutex

	id uint64
	saved []interface{}
	name string
	mu sync.Mutex
}
