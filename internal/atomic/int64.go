package atomic

import "sync"

type Int64 struct {
	val int64
	mx  sync.Mutex
}

func (a *Int64) Add(b int64) {
	a.mx.Lock()
	a.val = a.val + b
	a.mx.Unlock()
}

func (a *Int64) Load() int64 {
	a.mx.Lock()
	defer a.mx.Unlock()
	return a.val
}
