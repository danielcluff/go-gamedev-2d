package main

import (
	"fmt"
	"math"
	"math/rand"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Texture         r.Texture2D
	Pos             r.Vector2
	Speed           float32
	Direction       r.Vector2
	Size            r.Vector2
	Rotation        float32
	Scale           float32
	Discard         bool
	CollisionRadius float32
}

func SpriteCreate(texture r.Texture2D, pos r.Vector2, speed float32, direction r.Vector2, rotation float32, scale float32) Sprite {
	sprite := Sprite{
		Texture:         texture,
		Pos:             pos,
		Speed:           speed,
		Direction:       direction,
		Size:            r.Vector2{X: float32(texture.Width), Y: float32(texture.Height)},
		Rotation:        rotation,
		Scale:           scale,
		Discard:         false,
		CollisionRadius: float32(texture.Width) / 2,
	}
	return sprite
}
func (s *Sprite) CheckDiscard() {
	s.Discard = !(-300 < s.Pos.Y) || !(s.Pos.Y < WINDOW_HEIGHT+300)
}
func (s *Sprite) Move(dt float32) {
	s.Pos.X += s.Direction.X * s.Speed * dt
	s.Pos.Y += s.Direction.Y * s.Speed * dt
}
func (s *Sprite) Update(dt float32) {
	s.Move(dt)
	s.CheckDiscard()
	// fmt.Printf("sprite %v\n", s.Pos)
}
func (s *Sprite) GetCenter() r.Vector2 {
	return r.Vector2{
		X: s.Pos.X + s.Size.X/2,
		Y: s.Pos.Y + s.Size.Y/2,
	}
}
func (s Sprite) Draw() {
	r.DrawTextureV(s.Texture, s.Pos, r.White)
}

type Player struct {
	Sprite
	ShootLaser func(pos r.Vector2)
}

func PlayerCreate(texture r.Texture2D, position r.Vector2, shootLaser func(pos r.Vector2)) Player {
	player := Player{
		Sprite:     SpriteCreate(texture, position, PLAYER_SPEED, r.Vector2{X: 0, Y: 0}, 0, 1),
		ShootLaser: shootLaser,
	}
	return player
}
func (p *Player) Input() {
	p.Direction.X = float32(BoolToInt(r.IsKeyDown(r.KeyRight))) - float32(BoolToInt(r.IsKeyDown(r.KeyLeft)))
	p.Direction.Y = float32(BoolToInt(r.IsKeyDown(r.KeyDown))) - float32(BoolToInt(r.IsKeyDown(r.KeyUp)))
	p.Direction = r.Vector2Normalize(p.Direction)

	if r.IsKeyPressed(r.KeySpace) {
		p.ShootLaser(r.Vector2{X: p.Pos.X + p.Size.X/2 - 5, Y: p.Pos.Y - 60})
	}
}
func (p *Player) Update(dt float32) {
	p.Input()
	p.Move(dt)
	p.Constraint()
}
func (p *Player) Constraint() {
	p.Pos.X = float32(math.Max(0, math.Min(float64(p.Pos.X), float64(WINDOW_WIDTH-p.Size.X))))
	p.Pos.Y = float32(math.Max(0, math.Min(float64(p.Pos.Y), float64(WINDOW_HEIGHT-p.Size.Y))))
}

type Laser struct {
	Sprite
}

func LaserCreate(texture r.Texture2D, position r.Vector2) Laser {
	laser := Laser{
		Sprite: SpriteCreate(texture, position, LASER_SPEED, r.Vector2{X: 0, Y: -1}, 0, 1),
	}
	return laser
}

type Meteor struct {
	Hitbox r.Rectangle
	Sprite
}

func MeteorCreate(texture r.Texture2D) Meteor {
	meteor := Meteor{
		Hitbox: r.Rectangle{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)},
		Sprite: Sprite{
			Texture:         texture,
			Pos:             r.Vector2{X: float32(rand.Intn(2000) - 40), Y: float32(rand.Intn(600) - 150)},
			Direction:       r.Vector2{X: float32(math.Max(math.Min(0.5, (rand.Float64()*2)-1), -0.5)), Y: 1},
			Size:            r.Vector2{X: float32(texture.Width), Y: float32(texture.Height)},
			Speed:           float32(rand.Intn(101) + 300),
			Scale:           float32(math.Max(math.Min(rand.Float64(), 1.25), .75)),
			Rotation:        float32(math.Max(math.Min(0.5, (rand.Float64()*2)-1), -0.5)),
			CollisionRadius: float32(texture.Width) / 2,
		},
	}
	return meteor
}

func (p *Meteor) Update(dt float32) {
	fmt.Printf("meteor before\np %v\n", p.Sprite)
	p.Sprite.Update(dt)
	fmt.Printf("meteor after\np %v\n", p.Sprite)
	p.Rotation += p.Speed * dt / 2 * p.Direction.X
}
func (p *Meteor) GetCenter() r.Vector2 {
	return p.Pos
}
func (p Meteor) Draw() {
	targetRect := r.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: p.Size.X, Height: p.Size.Y}
	r.DrawTexturePro(p.Texture, p.Hitbox, targetRect, r.Vector2{X: p.Size.X / 2, Y: p.Size.Y / 2}, p.Rotation, r.White)
}

type ExplosionAnimation struct {
	Textures []r.Texture2D
	Pos      r.Vector2
	Size     r.Vector2
	Discard  bool
	Index    int
}

func ExplosionAnimationCreate(textures []r.Texture2D, pos r.Vector2) ExplosionAnimation {
	animation := ExplosionAnimation{
		Textures: textures,
		Size:     r.Vector2{X: float32(textures[0].Width), Y: float32(textures[0].Height)},
		Pos:      r.Vector2{X: pos.X - float32(textures[0].Width)/2, Y: pos.Y - float32(textures[0].Height)/2},
		Discard:  false,
		Index:    0,
	}
	return animation
}
func (a *ExplosionAnimation) Update(dt float32) {
	if a.Index < len(a.Textures)-1 {
		a.Index += int(20 * dt)
	} else {
		a.Discard = true
	}
}
func (a ExplosionAnimation) Draw() {
	r.DrawTextureV(a.Textures[a.Index], a.Pos, r.White)
}
