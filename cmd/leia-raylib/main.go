package main

import (
	"fmt"
	"os"
	"path/filepath"
	goruntime "runtime"

	"github.com/never-labs/leia"
	leiaraylib "github.com/never-labs/leia-raylib"
)

func init() {
	goruntime.LockOSThread()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: leia-raylib <file.leia> [args...]")
		os.Exit(2)
	}
	script := os.Args[1]
	opts := append([]leia.Option{}, leia.ModuleOptionsForScript(script)...)
	opts = append(opts, leia.WithArgs(script, os.Args[2:]...))
	if helper, ok := helperModulePath(); ok {
		opts = append(opts, leia.WithModuleReplace("github.com/never-labs/leia-raylib", helper))
	}
	vm := leia.New(opts...)
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
	return filepath.Join(root, "leia-raylib.leia"), true
}
