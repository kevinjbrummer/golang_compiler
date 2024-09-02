package main

import (
	"goblin/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}