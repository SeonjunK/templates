// Package main provides a CLI tool for Claude Code hooks.
//
// Usage:
//
//	hook <command>
//
// Commands:
//
//	guard-read   Check if a file read should be blocked
//	guard-write  Check if a file write should be blocked
//	guard-bash   Check if a bash command should be blocked
//	format       Format a Go file after write/edit
//	verify       Run all quality checks before session stop
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: hook <command>")
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "guard-read":
		err = guardRead()
	case "guard-write":
		err = guardWrite()
	case "guard-bash":
		err = guardBash()
	case "format":
		err = format()
	case "verify":
		err = verify()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "hook %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
}
