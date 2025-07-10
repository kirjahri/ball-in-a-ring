package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct{}

const (
	screenWidth  = 800
	screenHeight = 450
)

func (g *Game) Init() {}

func (g *Game) Update() {}

func (g *Game) Draw() {
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
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
