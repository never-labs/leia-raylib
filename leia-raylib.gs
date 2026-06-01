func requireRaylib() {
    if rl._stub {
        return nil, "raylib support is not compiled in; rebuild gscript with -tags rl"
    }
    return rl, nil
}

func color(r, g, b, a) {
    return {r: r, g: g, b: b, a: a}
}

func camera2D(opts) {
    if opts == nil {
        opts = {}
    }
    zoom := opts.zoom
    if zoom == nil || zoom == 0 {
        zoom = 1
    }
    offsetX := opts.offsetX
    offsetY := opts.offsetY
    targetX := opts.targetX
    targetY := opts.targetY
    rotation := opts.rotation
    if offsetX == nil { offsetX = 0 }
    if offsetY == nil { offsetY = 0 }
    if targetX == nil { targetX = 0 }
    if targetY == nil { targetY = 0 }
    if rotation == nil { rotation = 0 }
    return {
        offsetX: offsetX,
        offsetY: offsetY,
        targetX: targetX,
        targetY: targetY,
        rotation: rotation,
        zoom: zoom,
    }
}

func open(width, height, title, fps) {
    backend, err := requireRaylib()
    if err != nil {
        return nil, err
    }
    backend.initWindow(width, height, title)
    if fps != nil && fps > 0 {
        backend.setTargetFPS(fps)
    }
    return backend, nil
}

func close() {
    backend, err := requireRaylib()
    if err != nil {
        return nil, err
    }
    backend.closeWindow()
    return true, nil
}

func frame(clearColor, draw) {
    backend, err := requireRaylib()
    if err != nil {
        return nil, err
    }
    backend.beginDrawing()
    if clearColor != nil {
        backend.clearBackground(clearColor)
    }
    result, drawErr := draw(backend)
    backend.endDrawing()
    if drawErr != nil {
        return nil, drawErr
    }
    return result, nil
}

func run(opts) {
    backend, err := open(opts.width, opts.height, opts.title, opts.fps)
    if err != nil {
        return nil, err
    }
    frames := 0
    maxFrames := opts.maxFrames
    for !backend.windowShouldClose() {
        if maxFrames != nil && frames >= maxFrames {
            break
        }
        frames = frames + 1
        _, frameErr := frame(opts.clear, opts.draw)
        if frameErr != nil {
            backend.closeWindow()
            return nil, frameErr
        }
    }
    backend.closeWindow()
    return {frames: frames}, nil
}

return {
    requireRaylib: requireRaylib,
    color: color,
    camera2D: camera2D,
    open: open,
    close: close,
    frame: frame,
    run: run,
}
