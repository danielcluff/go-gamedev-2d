package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"path/filepath"

// 	r "github.com/gen2brain/raylib-go/raylib"
// )

// type Meteor struct {
// 	Texture  r.Texture2D
// 	Rotation float32
// 	Timer    *Timer
// 	Hitbox   r.Rectangle
// 	Sprite
// }

// func (p Meteor) Draw() {
// 	r.DrawTextureEx(p.Texture, p.Position, p.Rotation, p.Scale, r.White)
// }
// func (p *Meteor) Update(dt float32) {
// 	p.Direction = r.Vector2Normalize(p.Direction)

// 	p.Position.X += p.Direction.X * p.Speed * dt
// 	p.Position.Y += p.Direction.Y * p.Speed * dt
// 	p.Hitbox.X = p.Position.X
// 	p.Hitbox.Y = p.Position.Y

// 	p.Rotation += p.Speed * dt / 2 * p.Direction.X
// }
// func (p *Meteor) Reset() {
// 	p.Position.X = float32(rand.Intn(2000) - 40)
// 	p.Position.Y = float32(rand.Intn(100) - 150)
// 	fmt.Sprintln("reset function")
// }
// func MeteorCreate() *Meteor {
// 	meteor := &Meteor{
// 		Texture:  r.LoadTexture(filepath.Join("assets", "images", "meteor.png")),
// 		Rotation: (rand.Float32() * 2) - 1,
// 		Sprite: Sprite{
// 			Position:  r.Vector2{X: float32(rand.Intn(2000) - 40), Y: float32(rand.Intn(100) - 150)},
// 			Direction: r.Vector2{X: (rand.Float32() * 2) - 1, Y: 1},
// 			Speed:     float32(math.Max(300, float64(rand.Intn(400)))),
// 			Scale:     float32(math.Max(math.Min(rand.Float64(), 1.25), .75)),
// 		},
// 	}
// 	meteor.Hitbox = r.Rectangle{X: meteor.Position.X, Y: meteor.Position.Y, Width: 101, Height: 84}
// 	meteor.Timer = TimerCreate(0.4, true, true, meteor.Reset)
// 	return meteor
// }
