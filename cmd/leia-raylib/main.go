package main

import (
	"fmt"
	"os"
	"path/filepath"
	goruntime "runtime"

	"github.com/never-labs/gscript"
	leiaraylib "github.com/never-labs/leia-raylib"
)

func init() {
	goruntime.LockOSThread()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: leia-raylib <file.gs> [args...]")
		os.Exit(2)
	}
	script := os.Args[1]
	opts := append([]gscript.Option{}, gscript.ModuleOptionsForScript(script)...)
	opts = append(opts, gscript.WithArgs(script, os.Args[2:]...))
	if helper, ok := helperModulePath(); ok {
		opts = append(opts, gscript.WithModuleReplace("github.com/never-labs/leia-raylib", helper))
	}
	vm := gscript.New(opts...)
	if err := vm.RegisterModule("github.com/never-labs/leia-raylib/native", leiaraylib.Module()); err != nil {
		fmt.Fprintf(os.Stderr, "leia-raylib: %v\n", err)
		os.Exit(1)
	}
	if err := vm.ExecFile(script); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", script, err)
		os.Exit(1)
	}
}

func helperModulePath() (string, bool) {
	_, file, _, ok := goruntime.Caller(0)
	if !ok {
		return "", false
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	return filepath.Join(root, "leia-raylib.gs"), true
}
