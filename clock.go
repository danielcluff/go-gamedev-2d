package main

import (
	"fmt"
	"math"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Clock struct {
	counter int
}

func (c Clock) Draw(font r.Font) {
	r.DrawTextEx(font, fmt.Sprintf("%v", c.counter), r.Vector2{X: float32(WINDOW_WIDTH/2 - 24), Y: 50}, 64, 2, r.White)
}
func (c *Clock) UpdateClock() {
	c.counter = int(math.Floor(r.GetTime()))
}
func ClockCreate() Clock {
	clock := Clock{
		counter: 0,
	}
	return clock
}
