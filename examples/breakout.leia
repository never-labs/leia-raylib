leia := require("github.com/never-labs/leia-raylib")

screenW := 800
screenH := 520

bg := leia.color(19, 23, 32, 255)
panel := leia.color(31, 38, 52, 255)
white := leia.color(242, 245, 248, 255)
muted := leia.color(148, 163, 184, 255)
cyan := leia.color(34, 211, 238, 255)
green := leia.color(52, 211, 153, 255)
yellow := leia.color(250, 204, 21, 255)
orange := leia.color(251, 146, 60, 255)
red := leia.color(248, 113, 113, 255)

paddleW := 104.0
paddleH := 16.0
paddleX := 0.0
paddleY := 468.0
paddleSpeed := 520.0

ballX := 0.0
ballY := 0.0
ballVX := 0.0
ballVY := 0.0
ballR := 8.0

score := 0
lives := 3
remaining := 0
state := "playing"
bricks := {}

func clamp(v, lo, hi) {
    if v < lo {
        return lo
    }
    if v > hi {
        return hi
    }
    return v
}

func circleRectHit(cx, cy, r, x, y, w, h) {
    nx := clamp(cx, x, x + w)
    ny := clamp(cy, y, y + h)
    dx := cx - nx
    dy := cy - ny
    return dx * dx + dy * dy <= r * r
}

func resetBall(direction) {
    paddleX = (screenW - paddleW) / 2
    ballX = paddleX + paddleW / 2
    ballY = paddleY - 28
    ballVX = 210 * direction
    ballVY = -260
}

func buildBricks() {
    bricks = {}
    colors := {red, orange, yellow, green, cyan}
    index := 1
    remaining = 0

    for row := 1; row <= 5; row = row + 1 {
        for col := 1; col <= 10; col = col + 1 {
            bricks[index] = {
                x: 46 + (col - 1) * 71,
                y: 76 + (row - 1) * 29,
                w: 62,
                h: 20,
                alive: true,
                color: colors[row],
            }
            index = index + 1
            remaining = remaining + 1
        }
    }
}

func resetGame() {
    score = 0
    lives = 3
    state = "playing"
    buildBricks()
    resetBall(1)
}

func updatePaddle(rl, dt) {
    if rl.isKeyDown(rl.KEY_LEFT) || rl.isKeyDown(rl.KEY_A) {
        paddleX = paddleX - paddleSpeed * dt
    }
    if rl.isKeyDown(rl.KEY_RIGHT) || rl.isKeyDown(rl.KEY_D) {
        paddleX = paddleX + paddleSpeed * dt
    }
    paddleX = clamp(paddleX, 22, screenW - paddleW - 22)
}

func updateBall(dt) {
    ballX = ballX + ballVX * dt
    ballY = ballY + ballVY * dt

    if ballX < ballR {
        ballX = ballR
        ballVX = -ballVX
    }
    if ballX > screenW - ballR {
        ballX = screenW - ballR
        ballVX = -ballVX
    }
    if ballY < 52 + ballR {
        ballY = 52 + ballR
        ballVY = -ballVY
    }
    if ballY > screenH + ballR {
        lives = lives - 1
        if lives <= 0 {
            state = "lost"
        } else {
            dir := 1
            if lives == 2 {
                dir = -1
            }
            resetBall(dir)
        }
    }
}

func hitPaddle() {
    if ballVY > 0 && circleRectHit(ballX, ballY, ballR, paddleX, paddleY, paddleW, paddleH) {
        ballY = paddleY - ballR
        center := paddleX + paddleW / 2
        offset := (ballX - center) / (paddleW / 2)
        ballVX = offset * 330
        ballVY = -285
    }
}

func hitBricks() {
    collided := false
    for i := 1; i <= #bricks; i = i + 1 {
        b := bricks[i]
        if !collided && b.alive && circleRectHit(ballX, ballY, ballR, b.x, b.y, b.w, b.h) {
            b.alive = false
            collided = true
            score = score + 10
            remaining = remaining - 1

            left := ballX < b.x + 5
            right := ballX > b.x + b.w - 5
            if left || right {
                ballVX = -ballVX
            } else {
                ballVY = -ballVY
            }
        }
    }
    if remaining <= 0 {
        state = "won"
    }
}

func updateGame(rl) {
    dt := rl.getFrameTime()
    if dt > 0.033 {
        dt = 0.033
    }

    if state != "playing" {
        if rl.isKeyPressed(rl.KEY_SPACE) {
            resetGame()
        }
        return nil
    }

    updatePaddle(rl, dt)
    updateBall(dt)
    hitPaddle()
    hitBricks()
    return nil
}

func drawBricks(rl) {
    for i := 1; i <= #bricks; i = i + 1 {
        b := bricks[i]
        if b.alive {
            rl.drawRectangle(b.x, b.y, b.w, b.h, b.color)
        }
    }
}

func drawHud(rl) {
    rl.drawRectangle(0, 0, screenW, 48, panel)
    rl.drawText("Leia Breakout", 24, 15, 22, white)
    rl.drawText("Score: " .. tostring(score), 570, 16, 18, yellow)
    rl.drawText("Lives: " .. tostring(lives), 690, 16, 18, cyan)
}

func drawOverlay(rl) {
    if state == "won" {
        rl.drawText("Stage clear", 318, 236, 34, green)
        rl.drawText("Press SPACE to play again", 274, 276, 20, white)
    }
    if state == "lost" {
        rl.drawText("Game over", 318, 236, 34, red)
        rl.drawText("Press SPACE to restart", 286, 276, 20, white)
    }
}

resetGame()

result, err := leia.run({
    width: screenW,
    height: screenH,
    title: "Leia Raylib Breakout",
    fps: 60,
    clear: bg,
    draw: func(rl) {
        updateGame(rl)
        drawHud(rl)
        drawBricks(rl)
        rl.drawRectangle(paddleX, paddleY, paddleW, paddleH, white)
        rl.drawCircle(ballX, ballY, ballR, cyan)
        rl.drawText("Move: A/D or arrows", 24, 492, 16, muted)
        drawOverlay(rl)
        return true, nil
    },
})

if err != nil {
    print(err)
}
