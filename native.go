package leiaraylib

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/never-labs/leia"
)

// Module returns the native raylib binding exposed to Leia.
func Module() leia.Module {
	m := leia.Module{
		"_native":               true,
		"initWindow":            initWindow,
		"closeWindow":           rl.CloseWindow,
		"windowShouldClose":     rl.WindowShouldClose,
		"setTargetFPS":          func(fps int32) { rl.SetTargetFPS(fps) },
		"getFrameTime":          func() float64 { return float64(rl.GetFrameTime()) },
		"getFPS":                func() int64 { return int64(rl.GetFPS()) },
		"beginDrawing":          rl.BeginDrawing,
		"endDrawing":            rl.EndDrawing,
		"clearBackground":       clearBackground,
		"drawText":              drawText,
		"drawCircle":            drawCircle,
		"drawRectangle":         drawRectangle,
		"isKeyDown":             func(key int32) bool { return rl.IsKeyDown(key) },
		"isKeyPressed":          func(key int32) bool { return rl.IsKeyPressed(key) },
		"checkCollisionCircles": checkCollisionCircles,
	}
	for name, value := range colorConstants() {
		m[name] = colorMap(value)
	}
	for name, value := range keyConstants() {
		m[name] = int64(value)
	}
	return m
}

func initWindow(width, height int32, title string) {
	rl.InitWindow(width, height, title)
}

func clearBackground(c map[string]interface{}) {
	rl.ClearBackground(toColor(c))
}

func drawText(text string, x, y, fontSize int32, c map[string]interface{}) {
	rl.DrawText(text, x, y, fontSize, toColor(c))
}

func drawCircle(x, y int32, radius float64, c map[string]interface{}) {
	rl.DrawCircle(x, y, float32(radius), toColor(c))
}

func drawRectangle(x, y, width, height int32, c map[string]interface{}) {
	rl.DrawRectangle(x, y, width, height, toColor(c))
}

func checkCollisionCircles(x1, y1, r1, x2, y2, r2 float64) bool {
	return rl.CheckCollisionCircles(
		rl.Vector2{X: float32(x1), Y: float32(y1)},
		float32(r1),
		rl.Vector2{X: float32(x2), Y: float32(y2)},
		float32(r2),
	)
}

func toColor(v map[string]interface{}) color.RGBA {
	return color.RGBA{
		R: byteNumber(v["r"]),
		G: byteNumber(v["g"]),
		B: byteNumber(v["b"]),
		A: byteNumber(v["a"]),
	}
}

func colorMap(c color.RGBA) map[string]interface{} {
	return map[string]interface{}{"r": int64(c.R), "g": int64(c.G), "b": int64(c.B), "a": int64(c.A)}
}

func byteNumber(v interface{}) uint8 {
	switch n := v.(type) {
	case int:
		return uint8(n)
	case int64:
		return uint8(n)
	case float64:
		return uint8(n)
	case float32:
		return uint8(n)
	default:
		return 0
	}
}

func colorConstants() map[string]color.RGBA {
	return map[string]color.RGBA{
		"LIGHTGRAY":  {R: 200, G: 200, B: 200, A: 255},
		"GRAY":       {R: 130, G: 130, B: 130, A: 255},
		"DARKGRAY":   {R: 80, G: 80, B: 80, A: 255},
		"YELLOW":     {R: 253, G: 249, B: 0, A: 255},
		"GOLD":       {R: 255, G: 203, B: 0, A: 255},
		"ORANGE":     {R: 255, G: 161, B: 0, A: 255},
		"PINK":       {R: 255, G: 109, B: 194, A: 255},
		"RED":        {R: 230, G: 41, B: 55, A: 255},
		"MAROON":     {R: 190, G: 33, B: 55, A: 255},
		"GREEN":      {R: 0, G: 228, B: 48, A: 255},
		"LIME":       {R: 0, G: 158, B: 47, A: 255},
		"DARKGREEN":  {R: 0, G: 117, B: 44, A: 255},
		"SKYBLUE":    {R: 102, G: 191, B: 255, A: 255},
		"BLUE":       {R: 0, G: 121, B: 241, A: 255},
		"DARKBLUE":   {R: 0, G: 82, B: 172, A: 255},
		"PURPLE":     {R: 200, G: 122, B: 255, A: 255},
		"VIOLET":     {R: 135, G: 60, B: 190, A: 255},
		"DARKPURPLE": {R: 112, G: 31, B: 126, A: 255},
		"BEIGE":      {R: 211, G: 176, B: 131, A: 255},
		"BROWN":      {R: 127, G: 106, B: 79, A: 255},
		"DARKBROWN":  {R: 76, G: 63, B: 47, A: 255},
		"WHITE":      {R: 255, G: 255, B: 255, A: 255},
		"BLACK":      {R: 0, G: 0, B: 0, A: 255},
		"BLANK":      {R: 0, G: 0, B: 0, A: 0},
		"MAGENTA":    {R: 255, G: 0, B: 255, A: 255},
		"RAYWHITE":   {R: 245, G: 245, B: 245, A: 255},
	}
}

func keyConstants() map[string]int32 {
	return map[string]int32{
		"KEY_NULL":   0,
		"KEY_SPACE":  32,
		"KEY_ESCAPE": 256,
		"KEY_ENTER":  257,
		"KEY_RIGHT":  262,
		"KEY_LEFT":   263,
		"KEY_DOWN":   264,
		"KEY_UP":     265,
		"KEY_W":      87,
		"KEY_A":      65,
		"KEY_S":      83,
		"KEY_D":      68,
	}
}
