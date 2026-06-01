leia := require("github.com/never-labs/leia-raylib")

white := leia.color(245, 245, 245, 255)
red := leia.color(230, 41, 55, 255)

result, err := leia.run({
    width: 800,
    height: 450,
    title: "Leia Raylib",
    fps: 60,
    clear: white,
    draw: func(rl) {
        rl.drawText("Hello from GScript + raylib", 220, 200, 24, red)
        return true, nil
    },
})

if err != nil {
    print(err)
}
