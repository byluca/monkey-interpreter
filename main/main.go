package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s\n, This is the Monkey Programming language", user.Username)
	fmt.Printf("Write something\n")
	repl.Start(os.Stdin, os.Stdout)
}
