package dnslog

import (
	"fmt"
	"sync/atomic"
)

type Count64 int64

func (c *Count64) Inc() int64 {
	return atomic.AddInt64((*int64)(c), 1)
}

func (c *Count64) Get() int64 {
	return atomic.LoadInt64((*int64)(c))
}

type Dnslog struct {
	Allreq             Count64
	RcodeSuccess       Count64
	RcodeServerFailure Count64
	RcodeNameError     Count64
	Err                Count64
}

func (d *Dnslog) String() string {
	return fmt.Sprintln(d.Allreq, d.RcodeSuccess, d.RcodeServerFailure, d.RcodeNameError, d.Err)
}
