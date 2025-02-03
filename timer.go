package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	duration  float32
	startTime float64
	active    bool
	repeat    bool
	callback  func()
}

func (t *Timer) Activate() {
	t.active = true
	t.startTime = r.GetTime()
}
func (t *Timer) Deactivate() {
	t.active = false
	t.startTime = 0
	if t.repeat {
		t.Activate()
	}
}
func (t *Timer) Update() {
	if t.active {
		if r.GetTime()-t.startTime >= float64(t.duration) {
			if t.callback != nil && t.startTime > 0 {
				t.callback()
			}
			t.Deactivate()
		}
	}
}
func TimerCreate(duration float32, repeat bool, autostart bool, callback func()) *Timer {
	timer := &Timer{
		duration: duration,
		repeat:   repeat,
		callback: callback,
	}

	if autostart {
		timer.Activate()
	}

	return timer
}
