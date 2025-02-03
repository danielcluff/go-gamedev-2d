package main

import (
	"fmt"
	"math/rand"
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

// type Game struct{}

// func (g Game) Init() {
// 	r.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "2d space shooter")
// }
// func (g *Game) Run() {
// 	g.Init()
// 	defer r.CloseWindow()

//		for !r.WindowShouldClose() {
//			r.BeginDrawing()
//			r.EndDrawing()
//		}
//	}
func main() {
	// game := &Game{}
	// game.Run()
	r.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "2d space shooter")
	defer r.CloseWindow()

	r.SetTargetFPS(60)
	dafont := r.LoadFont(filepath.Join("assets", "font", "Stormfaze.otf"))

	// timer to generate asteroids

	var starField []Entity
	for i := 0; i < 40; i++ {
		star := SpriteCreate(float32(rand.Intn(WINDOW_WIDTH)), float32(rand.Intn(WINDOW_HEIGHT)), 0, 0, "star.png", 0)
		starField = append(starField, &star)
	}
	var meteors []Entity
	for i := 0; i < 10; i++ {
		meteor := MeteorCreate()
		fmt.Println(meteor)
		meteors = append(meteors, meteor)
	}
	clock := ClockCreate()

	player := PlayerCreate(1920/2-56, (1080/3)*2, PLAYER_SPEED)

	for !r.WindowShouldClose() {
		r.SetExitKey(r.KeyEscape)
		// Generate asteroids
		// movement and despawn

		// fire command
		// laser movement and hitbox collision

		// Update
		clock.UpdateClock()
		dt := r.GetFrameTime()
		EUpdate(meteors, dt)
		player.Update(dt)

		// Drawing
		r.BeginDrawing()
		r.ClearBackground(BG_COLOR)
		ERender(starField)
		ERender(meteors)

		player.Draw()
		clock.Draw(dafont)
		r.DrawFPS(5, 5)
		r.EndDrawing()
	}
}
