package main

// import (
// 	"math"
// 	"math/rand"
// 	"path/filepath"

// 	r "github.com/gen2brain/raylib-go/raylib"
// )

// type Sprite struct {
// 	Position  r.Vector2
// 	Direction r.Vector2
// 	Texture   r.Texture2D
// 	Speed     float32
// 	Scale     float32
// }

// func (e *Sprite) Move(dt float32) {
// 	e.Position.X = e.Direction.X * e.Speed * dt
// 	e.Position.Y = e.Direction.X * e.Speed * dt
// }
// func (e *Sprite) Update(dt float32) {
// 	e.Position.X += e.Direction.X * e.Speed * dt
// 	e.Position.Y += e.Direction.Y * e.Speed * dt
// }
// func (e Sprite) Draw() {
// 	r.DrawTextureEx(e.Texture, e.Position, 0, e.Scale, r.White)
// }
// func SpriteCreate(posx, posy float32, dirx, diry float32, texture string, speed float32) Sprite {
// 	sprite := Sprite{
// 		Position:  r.Vector2{X: posx, Y: posy},
// 		Direction: r.Vector2{X: dirx, Y: diry},
// 		Texture:   r.LoadTexture(filepath.Join("assets", "images", texture)),
// 		Scale:     float32(math.Max(math.Min(rand.Float64(), 1), .25)),
// 		Speed:     speed,
// 	}
// 	return sprite
// }
