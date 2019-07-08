package main

import (
	"fmt"
	"os"
	"os/exec"
)

func printError(s string) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printError("You must provide file/url")
	}

	exec.Command("mpv", os.Args[1]).Start()
}
