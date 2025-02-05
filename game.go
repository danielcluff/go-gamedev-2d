package main

import (
	"fmt"
	"math"
	"math/rand"
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Assets
	Starfield   []Sprite
	Lasers      []Laser
	Meteors     []Meteor
	Explosions  []ExplosionAnimation
	MeteorTimer Timer
	Player      Player
}
type Assets struct {
	player    r.Texture2D
	star      r.Texture2D
	laser     r.Texture2D
	meteor    r.Texture2D
	explosion []r.Texture2D
	font      r.Font
	audio     Audio
}
type Audio struct {
	music     r.Music
	laser     r.Sound
	explosion r.Sound
}

func (g *Game) Init() {
	r.InitWindow(1920, 1080, "2d space shooter")
	r.InitAudioDevice()
	g.ImportAssets()
	g.Player = PlayerCreate(g.Assets.player, r.Vector2{X: 1920/2 - float32(g.Assets.player.Width)/2, Y: (1080 / 3) * 2}, g.ShootLaser)
	g.Starfield = g.StarsGenerate(40)
	g.MeteorTimer = *TimerCreate(5, true, true, g.CreateMeteor)
	// r.PlayMusicStream(g.audio.music)

}
func (g *Game) Update() {
	dt := r.GetFrameTime()
	g.Player.Update(dt)
	g.MeteorTimer.Update()
	g.DiscardSprites()
	for _, laser := range g.Lasers {
		laser.Update(dt)
	}
	for _, meteor := range g.Meteors {
		meteor.Update(dt)
	}
	// r.UpdateMusicStream(g.audio.music)
}
func (g *Game) Draw() {
	r.BeginDrawing()
	r.ClearBackground(BG_COLOR)
	g.StarsDraw()
	g.DrawScore()
	g.MeteorsDraw()
	g.LasersDraw()
	g.Player.Draw()
	r.EndDrawing()
}
func (g *Game) Run() {
	g.Init()
	defer r.CloseWindow()
	defer r.CloseAudioDevice()
	// defer r.UnloadMusicStream(g.audio.music)

	for !r.WindowShouldClose() {

		g.Update()
		g.Draw()
	}
}
func (g *Game) ImportAssets() {
	g.Assets.player = r.LoadTexture(filepath.Join("assets", "images", "spaceship.png"))
	g.Assets.star = r.LoadTexture(filepath.Join("assets", "images", "star.png"))
	g.Assets.laser = r.LoadTexture(filepath.Join("assets", "images", "laser.png"))
	g.Assets.meteor = r.LoadTexture(filepath.Join("assets", "images", "meteor.png"))
	g.Assets.font = r.LoadFontEx(filepath.Join("assets", "font", "Stormfaze.otf"), FONT_SIZE, nil, 0)
	g.Assets.audio = Audio{
		music:     r.LoadMusicStream(filepath.Join("assets", "audio", "music.wav")),
		laser:     r.LoadSound(filepath.Join("assets", "audio", "laser.wav")),
		explosion: r.LoadSound(filepath.Join("assets", "audio", "explosion.wav")),
	}
}
func (g *Game) DiscardSprites() {
	var lasers []Laser
	for _, laser := range g.Lasers {
		if !laser.Discard {
			lasers = append(lasers, laser)
		}
	}
	g.Lasers = lasers

	var meteors []Meteor
	for _, meteor := range g.Meteors {
		if !meteor.Discard {
			meteors = append(meteors, meteor)
		}
	}
	g.Meteors = meteors

	var explosions []ExplosionAnimation
	for _, explosion := range g.Explosions {
		if !explosion.Discard {
			explosions = append(explosions, explosion)
		}
	}
	g.Explosions = explosions
}
func (g *Game) CheckCollisions() {
	// lasers and meteors
	for _, laser := range g.Lasers {
		for _, meteor := range g.Meteors {
			if r.CheckCollisionCircles(meteor.GetCenter(), meteor.CollisionRadius, laser.GetCenter(), laser.CollisionRadius) {
				laser.Discard = true
				meteor.Discard = true
				pos := r.Vector2{X: laser.Pos.X - laser.Size.X/2, Y: laser.Pos.Y}
				explosion := ExplosionAnimationCreate(g.Assets.explosion, pos)
				g.Explosions = append(g.Explosions, explosion)
				r.PlaySound(g.audio.explosion)
			}
		}
	}

	// player and meteors
	for _, meteor := range g.Meteors {
		if r.CheckCollisionCircles(g.Player.GetCenter(), g.Player.CollisionRadius, meteor.GetCenter(), meteor.CollisionRadius) {
			r.CloseWindow()
		}
	}

}
func (g *Game) DrawScore() {
	score := int(r.GetTime())
	textSize := r.MeasureTextEx(g.Assets.font, fmt.Sprintf("%d", score), FONT_SIZE, 0)
	r.DrawTextEx(g.Assets.font, fmt.Sprintf("%d", score), r.Vector2{X: WINDOW_WIDTH/2 - textSize.X/2, Y: 100}, FONT_SIZE, 0, r.White)
}
func (g *Game) StarsGenerate(num int) []Sprite {
	var stars []Sprite
	for i := 0; i < num; i++ {
		star := SpriteCreate(
			g.Assets.star,
			r.Vector2{X: float32(rand.Intn(WINDOW_WIDTH)), Y: float32(rand.Intn(WINDOW_HEIGHT))},
			0,
			r.Vector2{X: 0, Y: 0},
			(rand.Float32()*2-1)*100,
			float32(math.Min(math.Max(rand.Float64()*1.5, 0.5), 1.2)),
		)
		stars = append(stars, star)
	}
	return stars
}
func (g *Game) StarsDraw() {
	for _, star := range g.Starfield {
		r.DrawTextureEx(star.Texture, star.Pos, star.Rotation, star.Scale, r.White)
	}
}
func (g *Game) LasersDraw() {
	for _, laser := range g.Lasers {
		laser.Draw()
	}
}
func (g *Game) MeteorsDraw() {
	for _, meteor := range g.Meteors {
		meteor.Draw()
	}
}
func (g *Game) ShootLaser(pos r.Vector2) {
	laser := LaserCreate(g.Assets.laser, pos)
	g.Lasers = append(g.Lasers, laser)
	r.PlaySound(g.audio.laser)
}
func (g *Game) CreateMeteor() {
	meteor := MeteorCreate(g.Assets.meteor)
	// fmt.Printf("meteor props\ntexture %v\ndirection %v\nhitbox %v\npos %v\nspeed %v\nscale %v\nsprite %v\n", meteor.Texture, meteor.Direction, meteor.Hitbox, meteor.Pos, meteor.Speed, meteor.Scale, meteor.Sprite)
	g.Meteors = append(g.Meteors, meteor)
	// fmt.Printf("create meteor \n%v\n", g.Meteors)
}
