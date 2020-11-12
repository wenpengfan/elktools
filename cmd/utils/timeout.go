package utils

import (
	"time"
)

var (
	TimeoutTimer *timeout
)

type timeout struct {
	duration time.Duration
	timer    *time.Timer
}

func NewTimeout(d time.Duration, callback func()) {
	t := timeout{duration: d}
	t.timer = time.AfterFunc(t.duration, callback)
	TimeoutTimer = &t
}

func (t *timeout) Reset() bool {
	return t.timer.Reset(t.duration)
}
