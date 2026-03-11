// Package main provides the entry point for the hook CLI.
package main

import (
	"fmt"
	"os"

	"github.com/example/go-hook/internal/format"
	"github.com/example/go-hook/internal/guard"
	"github.com/example/go-hook/internal/verify"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: hook <command>")
		os.Exit(1)
	}

	var err error

	switch os.Args[1] {
	case "guard-read":
		err = guard.Read()
	case "guard-write":
		err = guard.Write()
	case "guard-bash":
		err = guard.Bash()
	case "format":
		err = format.Run()
	case "verify":
		err = verify.Run()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "hook %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
}
