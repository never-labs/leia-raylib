# leia-raylib

Raylib support for GScript as an external package. GScript itself does not ship
raylib bindings; this module provides the native binding and a small script
helper layer.

Run the demo directly with Go:

```bash
go run github.com/never-labs/leia-raylib/cmd/leia-raylib@v0.2.0 examples/bounce.gs
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
