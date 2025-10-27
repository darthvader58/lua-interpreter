package main

import (
	"fmt"
	"lua-interpreter/repl"
	"os"
)

func main() {
	user := os.Getenv("USER")
	if user == "" {
		user = "user"
	}

	fmt.Printf("Hello %s! This is the Lua programming language!\n", user)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
