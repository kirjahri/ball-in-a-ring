package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Circle struct {
	position rl.Vector2
	radius   float32
	color    rl.Color
}

type Ball struct {
	Circle
}

type Ring struct {
	Circle
}

type Game struct {
	balls []Ball
	ring  Ring
}

const (
	screenWidth  = 800
	screenHeight = 450
)

func (g *Game) Init() {
	g.balls = []Ball{{
		Circle: Circle{
			position: rl.NewVector2(screenWidth/2, screenHeight/2),
			radius:   20,
			color:    rl.Red,
		},
	}}
	g.ring = Ring{
		Circle: Circle{
			position: rl.NewVector2(screenWidth/2, screenHeight/2),
			radius:   200,
			color:    rl.Gray,
		},
	}
}

func (g *Game) Update() {}

func (g *Game) Draw() {
	for i := range g.balls {
		rl.DrawCircleV(g.balls[i].position, g.balls[i].radius, g.balls[i].color)
	}

	rl.DrawCircleLinesV(g.ring.position, g.ring.radius, g.ring.color)
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

		rl.ClearBackground(rl.Black)
		game.Draw()

		rl.EndDrawing()
	}
}
