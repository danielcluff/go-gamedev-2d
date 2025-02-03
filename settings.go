package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH, WINDOW_HEIGHT = 1920, 1080
const PLAYER_SPEED = 500
const LASER_SPEED = 600
const METEOR_TIMER_DURATION = 0.4
const FONT_SIZE = 120

var BG_COLOR = r.NewColor(15, 10, 25, 255)
var METEOR_SPEED_RANGE = []int32{300, 400}
