package main

type Entity interface {
	Draw()
	Update(dt float32)
}

func ERender(sprites []Entity) {
	for i := 0; i < len(sprites); i++ {
		sprites[i].Draw()
	}
}

func EUpdate(sprites []Entity, dt float32) {
	for i := 0; i < len(sprites); i++ {
		sprites[i].Update(dt)
	}
}
