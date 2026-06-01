leia := require("github.com/never-labs/leia-raylib")

sky := leia.color(36, 42, 54, 255)
lime := leia.color(0, 228, 48, 255)
white := leia.color(245, 245, 245, 255)
gold := leia.color(255, 203, 0, 255)

x := 120.0
y := 120.0
vx := 180.0
vy := 140.0
radius := 24.0
score := 0

result, err := leia.run({
    width: 800,
    height: 450,
    title: "Leia Raylib Bounce",
    fps: 60,
    clear: sky,
    maxFrames: 900,
    draw: func(rl) {
        dt := rl.getFrameTime()

        if rl.isKeyDown(rl.KEY_RIGHT) || rl.isKeyDown(rl.KEY_D) {
            vx = vx + 260 * dt
        }
        if rl.isKeyDown(rl.KEY_LEFT) || rl.isKeyDown(rl.KEY_A) {
            vx = vx - 260 * dt
        }
        if rl.isKeyDown(rl.KEY_DOWN) || rl.isKeyDown(rl.KEY_S) {
            vy = vy + 260 * dt
        }
        if rl.isKeyDown(rl.KEY_UP) || rl.isKeyDown(rl.KEY_W) {
            vy = vy - 260 * dt
        }

        x = x + vx * dt
        y = y + vy * dt

        if x < radius {
            x = radius
            vx = -vx
            score = score + 1
        }
        if x > 800 - radius {
            x = 800 - radius
            vx = -vx
            score = score + 1
        }
        if y < radius {
            y = radius
            vy = -vy
            score = score + 1
        }
        if y > 450 - radius {
            y = 450 - radius
            vy = -vy
            score = score + 1
        }

        rl.drawText("Leia Raylib external package", 24, 22, 24, white)
        rl.drawText("Arrow keys or WASD change velocity", 24, 54, 18, white)
        rl.drawText("Bounces: " .. tostring(score), 24, 82, 18, gold)
        rl.drawCircle(x, y, radius, lime)
        return true, nil
    },
})

if err != nil {
    print(err)
}
