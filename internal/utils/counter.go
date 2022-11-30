package utils

import "sync/atomic"

type Count64 int64

func (c *Count64) Inc() int64 {
	return atomic.AddInt64((*int64)(c), 1)
}

func (c *Count64) Get() int64 {
	return atomic.LoadInt64((*int64)(c))
}
