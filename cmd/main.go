package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Circle struct {
	position rl.Vector2
	radius   float32
	color    rl.Color
}

type Ball struct {
	Circle
	acceleration rl.Vector2
	velocity     rl.Vector2
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
	gravity      = 0.98
	gravityMult  = 100
	speedMult    = 2
	restitution  = 0.98
)

func (b *Ball) Update(dt float32) {
	b.velocity = rl.Vector2Add(b.velocity, rl.Vector2Scale(b.acceleration, dt))
	b.position = rl.Vector2Add(b.position, rl.Vector2Scale(b.velocity, dt*speedMult))
}

func (b *Ball) ResolveRingCollision(ring *Ring) {
	distance := rl.Vector2Distance(b.position, ring.position)
	maxDist := ring.radius - b.radius

	if distance > ring.radius-b.radius {
		normal := rl.Vector2Normalize(rl.Vector2Subtract(b.position, ring.position))
		b.position = rl.Vector2Add(ring.position, rl.Vector2Scale(normal, maxDist))

		dotProduct := rl.Vector2DotProduct(b.velocity, normal)
		deltaVel := rl.Vector2Subtract(b.velocity, rl.Vector2Scale(normal, 2*dotProduct))
		b.velocity = rl.Vector2Scale(deltaVel, restitution)
	}
}

func (g *Game) Init() {
	g.balls = []Ball{{
		Circle: Circle{
			position: rl.NewVector2(screenWidth/2, screenHeight/2),
			radius:   10,
			color:    rl.Red,
		},
		acceleration: rl.NewVector2(0, gravity*gravityMult),
		velocity:     rl.NewVector2(-50, -50),
	}}

	g.ring = Ring{
		Circle: Circle{
			position: rl.NewVector2(screenWidth/2, screenHeight/2),
			radius:   200,
			color:    rl.Gray,
		},
	}
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()

	for i := range g.balls {
		ball := &g.balls[i]
		ball.Update(dt)
		ball.ResolveRingCollision(&g.ring)
	}
}

func (g *Game) Draw() {
	rl.DrawCircleLinesV(g.ring.position, g.ring.radius, g.ring.color)

	for i := range g.balls {
		ball := &g.balls[i]
		rl.DrawCircleV(ball.position, ball.radius, ball.color)
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

		rl.ClearBackground(rl.Black)
		game.Draw()

		rl.EndDrawing()
	}
}
