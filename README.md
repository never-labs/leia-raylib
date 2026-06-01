# leia-raylib

Small GScript helper package for the built-in `rl` raylib module.

Use with a GScript binary built with raylib support:

```bash
go build -tags rl ./cmd/gscript
```

Example:

```gscript
leia := require("github.com/never-labs/leia-raylib")

leia.run({
    width: 800,
    height: 450,
    title: "Leia Raylib",
    fps: 60,
    clear: leia.color(245, 245, 245, 255),
    draw: func(rl) {
        rl.drawText("Hello", 20, 20, 24, rl.RED)
        return true, nil
    },
})
```
