package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	position rl.Vector2
	radius   float32
	color    rl.Color
}

type Game struct {
	balls []Ball
}

const (
	screenWidth  = 800
	screenHeight = 450
)

func (g *Game) Init() {
	g.balls = []Ball{{
		position: rl.NewVector2(screenWidth/2, screenHeight/2),
		radius:   20,
		color:    rl.Red,
	}}
}

func (g *Game) Update() {}

func (g *Game) Draw() {
	for i := range g.balls {
		rl.DrawCircleV(g.balls[i].position, g.balls[i].radius, g.balls[i].color)
	}
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(160)

	var game Game
	game.Init()

	for !rl.WindowShouldClose() {
		game.Update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		game.Draw()

		rl.EndDrawing()
	}
}
